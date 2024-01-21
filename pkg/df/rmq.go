package df

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqTransport struct {
	Ctx context.Context

	IsReconnecting    *sync.WaitGroup
	WorkersInProgress *sync.WaitGroup

	ConnectionString string

	Connection *amqp.Connection

	GlobalExchange   string
	GlobalRpcChannel *amqp.Channel

	DefaultMandatory    bool
	DefaultImmediate    bool
	DefaultCallTimeout  time.Duration
	DefaultReplyToQueue string

	ReplyQueueIdentifier string

	CallResponseChan map[string](chan *amqp.Delivery)
	ErrorChan        chan error
}

func (t *RmqTransport) Close() error {
	err := t.GlobalRpcChannel.Close()
	if err != nil {
		return err
	}
	err = t.Connection.Close()
	return err
}

func (t *RmqTransport) SubscribeOnCall(ctx context.Context, request Request, fn HandleCallRequest, mods ...SubscribeOnCallOptsModifier) terrors.Error {
	if request.GetName() == "" {
		return terrors.NewPrivateError("request name is empty")
	}

	t.IsReconnecting.Wait()

	consumerTag := uuid.New().String()
	channel, err := t.Connection.Channel()
	if err != nil {
		return terrors.NewPrivateError(fmt.Sprintf("failed to open a channel: %s", err.Error()))
	}
	opts := &SubscribeOnCallOpts{
		Parallel:  0,
		Exclusive: false,
		NoWait:    false,
	}

	for _, mod := range mods {
		err := mod(opts)
		if err != nil {
			return terrors.NewPrivateError(fmt.Sprintf("failed to apply modifier: %s", err.Error()))
		}
	}

	channelQueueOptions := amqp.Table{}

	if opts.SingleActiveConsumer {
		channelQueueOptions["x-single-active-consumer"] = true
	}

	queue, err := channel.QueueDeclare(
		request.GetName(),   // name
		false,               // durable
		true,                // auto delete
		opts.Exclusive,      // exclusive
		opts.NoWait,         // no wait
		channelQueueOptions, // args
	)
	if err != nil {
		return terrors.NewPrivateError(fmt.Sprintf("failed to create queue: %s", err.Error()))
	}

	msgs, err := channel.ConsumeWithContext(
		ctx,
		queue.Name,          // queue
		consumerTag,         // consumer
		false,               // auto ack
		false,               // exclusive
		false,               // no local
		false,               // no wait
		channelQueueOptions, // args
	)
	if err != nil {
		return terrors.NewPrivateError(fmt.Sprintf("failed to register a consumer: %s", err.Error()))
	}

	go func() {
		for msg := range msgs {
			select {
			case <-ctx.Done():
				return
			default:
				t.WorkersInProgress.Add(1)
				// # Defer cleanup
				defer func() {
					t.WorkersInProgress.Done()
					ackErr := msg.Ack(false)
					if ackErr != nil {
						t.ErrorChan <- err
					}
					if recRes := recover(); recRes != nil {
						if err, ok := recRes.(error); ok {
							t.ErrorChan <- err
						}
						// TODO: Add logger
						fmt.Println("Recover:", recRes)
					}
				}()
				// # Handle request
				response, err := fn(t.Ctx, msg.Body)
				if err != nil {
					t.ErrorChan <- err
				} else if response != nil {
					err := channel.PublishWithContext(
						ctx,
						t.GlobalExchange, // Default exchange in RabbitMQ called ""
						msg.ReplyTo,
						false,
						false,
						amqp.Publishing{
							ContentType:   "application/json",
							Body:          response,
							CorrelationId: msg.CorrelationId,
						},
					)
					if err != nil {
						t.ErrorChan <- err
					}
				} else {
					err := terrors.NewPrivateError(fmt.Sprintf("empty response and no error on %s (%s)", request.GetName(), msg.CorrelationId))
					t.ErrorChan <- err
				}
			}
		}
	}()

	return nil
}

func (t *RmqTransport) PublishCall(ctx context.Context, request Request, modifiersList ...PublishCallOptsModifier) ([]byte, error) {
	if request.GetName() == "" {
		return nil, terrors.NewPrivateError("request name is empty")
	}

	if request.GetId() == "" {
		return nil, terrors.NewPrivateError("request id is empty")
	}

	t.IsReconnecting.Wait()

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	correlationId := request.GetId()

	// # Sub on response
	responseChan := make(chan *amqp.Delivery)
	t.CallResponseChan[correlationId] = responseChan
	defer func() {
		close(responseChan)
		delete(t.CallResponseChan, correlationId)
	}()

	opts := &PublishCallOpts{
		Mandatory: t.DefaultMandatory,
		Immediate: t.DefaultImmediate,
		Timeout:   t.DefaultCallTimeout,
		ReplyTo:   t.DefaultReplyToQueue,
	}
	for _, mod := range modifiersList {
		mod(opts)
	}

	err = t.GlobalRpcChannel.PublishWithContext(
		ctx,
		t.GlobalExchange,  // exchange
		request.GetName(), // routing key
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: correlationId,
			ReplyTo:       opts.ReplyTo,
		},
	)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(opts.Timeout):
		return nil, terrors.NewTimeoutError(fmt.Sprintf("timeout on %s (%s)", request.GetName(), correlationId), nil)
	case response := <-responseChan:
		return response.Body, nil
	}
}

func (t *RmqTransport) Init(ctx context.Context) error {
	connection, err := amqp.Dial(t.ConnectionString)
	if err != nil {
		return err
	}

	t.Connection = connection

	globalWriteChannel, err := connection.Channel()
	if err != nil {
		return err
	}

	t.GlobalRpcChannel = globalWriteChannel

	// TODO: # Add this
	// returnCh := make(chan amqp.Return)

	// returnCh = globalWriteChannel.NotifyReturn(returnCh)

	// go func() {
	// 	for ret := range returnCh {
	// 		fmt.Println("Return:", ret)
	// 	}
	// }()

	msgs, err := globalWriteChannel.ConsumeWithContext(
		ctx,
		t.DefaultReplyToQueue,  // queue
		t.ReplyQueueIdentifier, // consumer
		true,                   // auto ack
		false,                  // exclusive
		false,                  // no local
		false,                  // no wait
		nil,                    //args
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for msg := range msgs {
			select {
			case <-ctx.Done():
				return
			default:
				if responseChan, ok := t.CallResponseChan[msg.CorrelationId]; ok {
					responseChan <- &msg
				}
			}
		}
	}()

	// # On done
	go func() {
		<-ctx.Done()
		globalWriteChannel.Close()
		connection.Close()
	}()

	return err
}

type RmqTransportModifier func(transport RmqTransport) (RmqTransport, error)

func NewRmqTransport(modifiers ...RmqTransportModifier) (*RmqTransport, error) {
	transport := RmqTransport{
		Ctx:                  context.Background(),
		GlobalExchange:       "",
		IsReconnecting:       &sync.WaitGroup{},
		WorkersInProgress:    &sync.WaitGroup{},
		DefaultMandatory:     true,
		DefaultImmediate:     false,
		DefaultCallTimeout:   30 * time.Second,
		DefaultReplyToQueue:  "amq.rabbitmq.reply-to",
		CallResponseChan:     map[string](chan *amqp.Delivery){},
		ReplyQueueIdentifier: uuid.New().String(),
		ErrorChan:            make(chan error),
	}

	for _, changer := range modifiers {
		var err error
		transport, err = changer(transport)
		if err != nil {
			return nil, err
		}
	}

	return &transport, nil
}

package df_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	inttests "github.com/Dionid/go-boiler/internal/int-tests"
	"github.com/Dionid/go-boiler/pkg/df"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIntRmqTransport(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	testConfig, err := inttests.NewTestConfig()

	if err != nil {
		t.Fatal(err)
	}

	rmqTransport, err := df.NewRmqTransport(
		func(transport df.RmqTransport) (df.RmqTransport, error) {
			transport.ConnectionString = testConfig.RmqConnectionString

			return transport, nil
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	err = rmqTransport.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for err := range rmqTransport.ErrorChan {
			cancel(err)
		}
	}()

	hit := 0

	callHandler := func(ctx context.Context, requestRaw []byte) ([]byte, terrors.Error) {
		request := &proto.SignInCallRequest{}
		err := json.Unmarshal(requestRaw, request)
		if err != nil {
			return nil, terrors.NewPrivateError(err.Error())
		}

		response := &proto.SignInCallResponse{
			Id: request.Id,
		}

		hit += 1

		result, err := json.Marshal(response)
		if err != nil {
			return nil, terrors.NewPrivateError(err.Error())
		}

		return result, nil
	}

	err = rmqTransport.SubscribeOnCall(
		ctx,
		&proto.SignInCallRequest{
			Name: "create-office",
		},
		callHandler,
	)
	if err != nil {
		t.Fatal(err)
	}

	requestId := uuid.New().String()

	response, err := rmqTransport.PublishCall(ctx, &proto.SignInCallRequest{
		Name: "create-office",
		Id:   requestId,
	}, func(opts *df.PublishCallOpts) error {
		opts.Timeout = 10 * time.Second
		opts.Mandatory = false

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	callResponse := &proto.SignInCallResponse{}
	json.Unmarshal(response, callResponse)

	assert.Equal(t, 1, hit)
	assert.Equal(t, callResponse.Id, requestId)
	assert.Nil(t, ctx.Err())
}

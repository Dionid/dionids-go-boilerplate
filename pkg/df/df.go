package df

import (
	"context"
	"time"

	"github.com/Dionid/go-boiler/pkg/terrors"
)

type Transport interface {
}

type Request interface {
	GetId() string
	GetName() string
}

type PublishCallOpts struct {
	Mandatory bool
	Immediate bool
	Timeout   time.Duration
	ReplyTo   string
}

type PublishCallOptsModifier func(opts *PublishCallOpts) error

type SubscribeOnCallOpts struct {
	Parallel             int
	SingleActiveConsumer bool

	Exclusive bool
	NoWait    bool
}

type SubscribeOnCallOptsModifier func(opts *SubscribeOnCallOpts) error

type HandleCallRequest func(ctx context.Context, request []byte) ([]byte, terrors.Error)

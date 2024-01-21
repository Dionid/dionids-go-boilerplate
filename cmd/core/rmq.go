package main

import (
	"context"

	"github.com/Dionid/go-boiler/pkg/df"
	"github.com/Dionid/go-boiler/pkg/terrors"
)

func initRmq(ctx context.Context, config *Config) (*df.RmqTransport, terrors.Error) {
	rmqTransport, err := df.NewRmqTransport(
		func(transport df.RmqTransport) (df.RmqTransport, error) {
			transport.ConnectionString = config.RmqConnectionString
			return transport, nil
		},
	)
	if err != nil {
		return nil, terrors.NewPrivateError("failed to create rmq transport")
	}

	err = rmqTransport.Init(ctx)
	if err != nil {
		return nil, terrors.NewPrivateError("failed to init rmq transport")
	}

	return rmqTransport, nil
}

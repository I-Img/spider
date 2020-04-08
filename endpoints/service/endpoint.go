package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

type Set struct {
	FetchEndpoint endpoint.Endpoint
}

func NewSet(svc PictureService, logger log.Logger, duration metrics.Histogram) Set {
	var fetchEndpoint endpoint.Endpoint

	{
		fetchEndpoint = MakeFetchEndpoint(svc)
	}

	return Set{
		FetchEndpoint: fetchEndpoint,
	}
}

func MakeFetchEndpoint(s PictureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)

		return s.Fetch(ctx, req)
	}
}

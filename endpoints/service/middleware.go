package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

type Middleware func(PictureService) PictureService

// LoggingMiddleware takes a logger as a dependency
// and returns a service Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PictureService) PictureService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   PictureService
}

func (l loggingMiddleware) Fetch(context context.Context, request Request) (response Response, err error) {
	l.logger.Log("method", "Fetch", "UUID", request.UUID, "Pos", request.Pos, "err", err)
	defer func() {
		l.logger.Log("method", "Fetch", "UUID", request.UUID, "Pos", request.Pos, "err", err)
	}()
	return l.next.Fetch(context, request)
}

func InstrumentingMiddleware(ints metrics.Counter) Middleware {
	return func(next PictureService) PictureService {
		return instrumentingMiddleware{
			ints: ints,
			next: next,
		}
	}
}

type instrumentingMiddleware struct {
	ints metrics.Counter
	next PictureService
}

func (i instrumentingMiddleware) Fetch(context context.Context, request Request) (Response, error) {
	i.ints.Add(1)
	return i.next.Fetch(context, request)
}

package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"spider/layout/flutter"
	"spider/storge/postgres"
)

type Request struct {
	UUID string
	Pos  int32
}

type Response struct {
	Pics []Picture
}

type Picture struct {
	Url   string
	AxisX int
	AxisY int
}

type PictureService interface {
	Fetch(context context.Context, request Request) (Response, error)
}

func NewBasicService() PictureService {
	return basicService{}
}

type basicService struct{}

func (b basicService) Fetch(context context.Context, request Request) (res Response, err error) {

	d := postgres.GetDriver()
	contents, err := d.FetchPictures(int(request.Pos))
	if err != nil {
		return
	}

	var pics []Picture

	for _, c := range contents {

		rate := flutter.Calc([]flutter.Size{
			flutter.Size{
				Width:  c.Width,
				Height: c.Height,
			},
		})

		pics = append(pics, Picture{
			Url:   c.Content,
			AxisX: rate[0].Width,
			AxisY: rate[0].Height,
		})
	}
	return Response{
		Pics: pics,
	}, nil
}

func New(logger log.Logger, ints metrics.Counter) PictureService {
	var svc PictureService
	svc = NewBasicService()
	svc = LoggingMiddleware(logger)(svc)
	svc = InstrumentingMiddleware(ints)(svc)
	return svc
}

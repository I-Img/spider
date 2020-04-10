package main

import (
	context "context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	rpc_v1 "spider/grpc/golang"
	"spider/layout/flutter"
	"spider/storge/postgres"
)

func startGRPC(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hs := health.NewServer()
	hs.SetServingStatus("PictureService", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, hs)

	rpc_v1.RegisterPictureServiceServer(s, &server{
		db: postgres.GetDriver(),
	})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {

	}
}

type server struct {
	db *postgres.Driver
}

func (s server) FetchPicture(ctx context.Context, request *rpc_v1.FetchPicturesRequest) (*rpc_v1.FetchPicturesResponse, error) {
	contents, err := s.db.FetchPictures(int(request.Pos))
	if err != nil {
		return &rpc_v1.FetchPicturesResponse{
			Status: rpc_v1.StatusCode_FAILE,
			Msg:    err.Error(),
			Info:   nil,
		}, nil
	}

	var info []*rpc_v1.PictureInfo

	for _, c := range contents {

		rate := flutter.Calc([]flutter.Size{flutter.Size{
			Width:  int(c.Width),
			Height: int(c.Height),
		}})

		info = append(info, &rpc_v1.PictureInfo{
			Url:    c.Content,
			Axis_X: int32(rate[0].Width),
			Axis_Y: int32(rate[0].Height),
		})
	}

	return &rpc_v1.FetchPicturesResponse{
		Status: rpc_v1.StatusCode_SUCC,
		Info:   info,
	}, nil
}

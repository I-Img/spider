package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	rpc_v1 "spider/grpc/golang"
)

type grpcServer struct {
	fetch grpctransport.Handler
}

func (g grpcServer) FetchPicture(ctx context.Context, request *rpc_v1.FetchPicturesRequest) (*rpc_v1.FetchPicturesResponse, error) {
	res := new(rpc_v1.FetchPicturesResponse)

	_, response, err := g.fetch.ServeGRPC(ctx, request)
	if err != nil {
		res.Status = rpc_v1.StatusCode_FAILE
		res.Msg = err.Error()
		return res, nil
	}

	res.Status = rpc_v1.StatusCode_SUCC
	_rep := response.(Response)

	var info []*rpc_v1.PictureInfo
	for _, p := range _rep.Pics {
		info = append(info, &rpc_v1.PictureInfo{
			Url:    p.Url,
			Axis_X: int32(p.AxisX),
			Axis_Y: int32(p.AxisY),
		})
	}

	res.Info = info

	return res, nil
}

func NewGRPCServer(endpoints Set, logger log.Logger) rpc_v1.PictureServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{
		fetch: grpctransport.NewServer(
			endpoints.FetchEndpoint,
			decodeGRPCFetchRequest,
			encodeGRPCFetchResponse,
			options...,
		),
	}
}

func decodeGRPCFetchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*rpc_v1.FetchPicturesRequest)

	return Request{
		UUID: req.UUID,
		Pos:  req.Pos,
	}, nil
}

func encodeGRPCFetchResponse(_ context.Context, response interface{})(interface{}, error){
	return response, nil
}
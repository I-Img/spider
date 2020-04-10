package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	rpc_v1 "spider/grpc/golang"
)

func main() {
	conn, err := grpc.Dial("150.158.134.74:80", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := rpc_v1.NewPictureServiceClient(conn)

	r, err := c.FetchPicture(context.Background(), &rpc_v1.FetchPicturesRequest{
		UUID: "123",
		Pos:  0,
	})
	if err != nil {
		panic(err)
	}

	if r.Status == rpc_v1.StatusCode_FAILE {
		panic(r.Msg)
	}

	for _, c := range r.Info {
		fmt.Printf("%s %d %d\n", c.Url, c.Axis_X, c.Axis_Y)
	}
}
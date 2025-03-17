package grpcservice

import (
	"crypto-dashboard/pkg/response"

	"google.golang.org/grpc"
)

type GrpcDial struct {
	*grpc.ClientConn
}

func NewGrpcDial(url string, options ...grpc.DialOption) (*GrpcDial, *response.AppError) {
	g, err := grpc.NewClient(url, options...)
	if err != nil {
		return nil, response.ServerError(err.Error())
	}
	return &GrpcDial{g}, nil
}

func (c *GrpcDial) CloseGrpcDial() {
	if c == nil {
		return
	}
	c.ClientConn.Close()
}

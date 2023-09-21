package service

import (
	"context"
	"time"

	"github.com/keenoho/go-core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MicroServiceClient struct {
	core.Logger
}

func (c *MicroServiceClient) Request(targetAddr string, url string, data any) (*ServiceMsgResponse, error) {
	dataBytes, _ := DataToBytes(data)
	requestData := &ServiceMsgRequest{
		Url:  url,
		Data: dataBytes,
	}
	conn, err := grpc.Dial(
		targetAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewServiceMsgHandlerClient(conn)
	grpcCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Send(grpcCtx, requestData)
	if err != nil {
		return nil, err
	}
	return res, nil
}

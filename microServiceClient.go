package core

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type MicroServiceClient struct {
	Pool    any
	Address string
}

func (c *MicroServiceClient) GetConnect() (*grpc.ClientConn, error) {
	return grpc.Dial(
		c.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func (c *MicroServiceClient) Send(address string, requestData *ServiceMsgRequest) (*ServiceMsgResponse, error) {
	conn, err := c.GetConnect()
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

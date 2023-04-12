package core

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func MicroServiceOnceRequest(targetAddr string, requestData *ServiceMsgRequest) (*ServiceMsgResponse, error) {
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

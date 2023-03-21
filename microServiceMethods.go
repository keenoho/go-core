package core

import (
	context "context"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 微服务间发送请求
func SendMicroServiceRequest(request *MicroServiceMsgRequest, targetService ...string) (*MicroServiceMsgResponse, error) {
	conn, err := grpc.Dial("0.0.0.0:8010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewMicroServiceMsgHandlerClient(conn)
	grpcCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := client.Send(grpcCtx, request)
	if err != nil {
		return nil, err
	}
	return r, err
}

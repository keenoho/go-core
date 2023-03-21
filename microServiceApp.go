package core

import (
	context "context"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

// 微服务app
type MicroServiceAppInterface interface {
	Run(addr string)
}

type MicroServiceApp struct {
	MsgHandler MicroServiceMsgHandler
}

type server struct {
	UnimplementedMicroServiceMsgHandlerServer
}

func (s *server) Send(ctx context.Context, in *MicroServiceMsgRequest) (*MicroServiceMsgResponse, error) {
	fmt.Println(ctx, in)
	return &MicroServiceMsgResponse{}, nil
}

func (app *MicroServiceApp) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	ss := server{}
	RegisterMicroServiceMsgHandlerServer(s, &ss)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

package core_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/keenoho/go-core/grpc_engine"
	"github.com/keenoho/go-core/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BaseServiceServer struct {
	protobuf.UnimplementedBaseServiceServer
}

// implements
func (s *BaseServiceServer) BaseRequest(ctx context.Context, in *protobuf.BaseRequestBody) (resp *protobuf.BaseResponseBody, err error) {
	fmt.Println(in)
	return resp, err
}

func TestBaseProtobuf(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		lis.Close()
		t.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterBaseServiceServer(s, &BaseServiceServer{})
	if err := s.Serve(lis); err != nil {
		t.Fatalf("failed to serve: %v", err)
	}
}

func TestGrpcEngine(t *testing.T) {
	engine := grpc_engine.New()
	engine.RegisterService(&protobuf.BaseService_ServiceDesc, &BaseServiceServer{})
	engine.Run("0.0.0.0:1234")
}

func TestGrpcEngineRequest(t *testing.T) {
	conn, err := grpc.NewClient("0.0.0.0:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewBaseClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.BaseRequest(ctx, &protobuf.BaseRequestBody{Action: "foo", Data: []byte("bar")})
	if err != nil {
		log.Fatalf("could not vist: %v", err)
	}
	t.Logf("response: %v", r)
}

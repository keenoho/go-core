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

type RouterServiceServer struct {
	protobuf.UnimplementedRouterServer
}

// implements
func (s *RouterServiceServer) RouterRequest(ctx context.Context, in *protobuf.RouterRequestBody) (resp *protobuf.RouterResponseBody, err error) {
	fmt.Println(in)
	return resp, err
}

func TestRouterProtobuf(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		lis.Close()
		t.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterRouterServer(s, &RouterServiceServer{})
	if err := s.Serve(lis); err != nil {
		t.Fatalf("failed to serve: %v", err)
	}
}

func TestGrpcEngine(t *testing.T) {
	engine := grpc_engine.New()
	engine.RegisterService(&protobuf.Router_ServiceDesc, &RouterServiceServer{})
	engine.Run("0.0.0.0:1234")
}

func TestGrpcEngineRequest(t *testing.T) {
	conn, err := grpc.NewClient("0.0.0.0:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewRouterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RouterRequest(ctx, &protobuf.RouterRequestBody{Url: "some url", Body: []byte("foo")})
	if err != nil {
		log.Fatalf("could not vist: %v", err)
	}
	t.Logf("response: %v", r)
}

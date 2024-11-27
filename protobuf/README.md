# protobuf

## prepare

[official golang tutorial](https://protobuf.dev/getting-started/gotutorial/)

1. download and install [v27.3](https://github.com/protocolbuffers/protobuf/releases/tag/v27.3)
2. run command to install the Go protocol buffers plugin
   ```shell
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
3. run `build.sh`

## develop

Server Side:

```go
// define struct
type RouterServiceServer struct {
	pb.UnimplementedRouterServiceServer
}
// implements
func (s *RouterServiceServer) Request(ctx context.Context, in *pb.RouterRequestBody) (resp *pb.RouterResponseBody, err error) {
	return resp, err
}

// main
// use engine
engine := grpc_engine.New()
engine.RegisterService(&pb.RouterService_ServiceDesc, &RouterServiceServer{})
engine.Run("0.0.0.0:1234")

// or use origin code
lis, err := net.Listen("tcp", "0.0.0.0:1234")
if err != nil {
   lis.Close()
   t.Fatalf("failed to listen: %v", err)
}
s := grpc.NewServer()
pb.RegisterRouterServiceServer(s, &RouterServiceServer{})
if err := s.Serve(lis); err != nil {
   t.Fatalf("failed to serve: %v", err)
}
```

Client Side:

```go
conn, err := grpc.NewClient("0.0.0.0:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
if err != nil {
   t.Fatalf("did not connect: %v", err)
}
defer conn.Close()
c := pb.NewRouterServiceClient(conn)
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
r, err := c.Request(ctx, &pb.RouterRequestBody{Url: "/foo", Data: []byte("bar")})
if err != nil {
   log.Fatalf("could not vist: %v", err)
}
t.Logf("response: %v", r)
```

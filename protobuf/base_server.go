package protobuf

import context "context"

type DefaultServer struct {
	UnimplementedBaseServiceServer
}

func (s *DefaultServer) Request(ctx context.Context, in *RequestBody) (resp *ResponseBody, err error) {
	return resp, err
}

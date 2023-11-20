package service

import (
	"context"
	"encoding/json"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MicroServiceClient struct{}

func (c *MicroServiceClient) OriginRequest(target string, url string, data any) (*ServiceResponse, error) {
	var dataBytes []byte
	_, isBytes := data.([]byte)
	if isBytes {
		dataBytes = data.([]byte)
	} else {
		dataBytes, _ = json.Marshal(data)

	}
	requestData := ServiceRequest{
		Url:  url,
		Data: dataBytes,
	}
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewServiceHandlerClient(conn)
	grpcCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Request(grpcCtx, &requestData)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *MicroServiceClient) Request(target string, url string, data any) (error, resultData any, resultCode int, resultMsg string) {
	resp, err := c.OriginRequest(target, url, data)
	if err != nil {
		return err, nil, 0, ""
	}
	if len(resp.Data) > 0 {
		json.Unmarshal(resp.Data, &resultData)
	}
	resultCode = int(resp.Code)
	resultMsg = resp.Msg
	return err, resultData, resultCode, resultMsg

}

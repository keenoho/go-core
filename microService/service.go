package microService

import (
	"log"
	"net"
	"strings"

	grpc "google.golang.org/grpc"
)

type MicroServiceInterface interface {
	Run(addr string)
}

type MicroService struct {
	Mode string
}

var MicroServiceDebugMode = "debug"
var MicroServiceReleaseMode = "release"

func (s *MicroService) print(printType string, format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	log.Printf("[MicroService-"+printType+"] "+format, values...)
}

func (s *MicroService) debugPrint(format string, values ...any) {
	if s.Mode == MicroServiceDebugMode {
		s.print("debug", format, values...)
	}
}

func (s *MicroService) errorPrint(format string, values ...any) {
	s.print("error", format, values...)
}

func (app *MicroService) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	ss := server{}
	RegisterMicroServiceMsgHandlerServer(s, &ss)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func New() *MicroService {
	ms := new(MicroService)
	ms.Mode = MicroServiceReleaseMode
	return ms
}

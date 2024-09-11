package grpc_engine

import (
	"log"

	"google.golang.org/grpc"
)

func debugPrintRegisterService(serviceDesc *grpc.ServiceDesc) {
	debugPrint("Registering service:", serviceDesc.ServiceName, serviceDesc.HandlerType)
}

func debugPrint(format string, values ...any) {
	if Mode != DebugMode {
		return
	}
	log.Printf("[grpc-engine-debug] "+format, values...)
}

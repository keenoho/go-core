package grpc_engine

import "fmt"

func debugPrint(format string, values ...any) {
	if Mode != DebugMode {
		return
	}
	str := fmt.Sprintf("[grpc-engine-debug] "+format, values...)
	fmt.Println(str)
}

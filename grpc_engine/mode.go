package grpc_engine

const ReleaseMode = "release"
const TestMode = "test"
const DebugMode = "debug"

var Mode string = ReleaseMode

func SetMode(value string) {
	switch value {
	case ReleaseMode:
		Mode = ReleaseMode
	case TestMode:
		Mode = TestMode
	case DebugMode:
		Mode = DebugMode
	default:
		panic("mode unknown: " + value + " (available mode: debug test release)")
	}
}

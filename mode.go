package core

const ReleaseMode = "release"
const TestMode = "test"
const DebugMode = "debug"

var Mode string = ReleaseMode

func SetMode(value string) {
	if len(value) < 1 {
		return
	}

	switch value {
	case ReleaseMode:
		Mode = ReleaseMode
	case TestMode:
		Mode = TestMode
	case DebugMode:
		Mode = DebugMode
	default:
		panic("mode unknown: " + value + " (available mode: debug release)")
	}
}

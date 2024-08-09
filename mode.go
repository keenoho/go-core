package core

const ReleaseMode = "release"
const DebugMode = "debug"

var Mode string = ReleaseMode

func SetMode(value string) {
	if len(value) < 1 {
		return
	}

	switch value {
	case ReleaseMode:
		Mode = ReleaseMode
	case DebugMode:
		Mode = DebugMode
	default:
		panic("mode unknown: " + value + " (available mode: debug release)")
	}
}

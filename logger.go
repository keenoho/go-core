package core

import (
	"log"
	"strings"
)

func Log(format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	log.Printf(format, values...)
}

func LogDebug(format string, values ...any) {
	if Mode != DebugMode {
		return
	}
	Log("[debug] "+format, values...)
}

package core

import (
	"log"
	"strings"
)

type LoggerInterface interface {
	Print()
	PrintDebug()
	PrintError()
}

type Logger struct {
}

func (l *Logger) PrintLog(tag string, format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	log.Printf("["+tag+"] "+format, values...)
}

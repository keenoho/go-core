package core

import (
	"fmt"
	"strings"
)

type LoggerInterface interface {
	Print()
	PrintInfo()
	PrintDebug()
	PrintError()
}

type Logger struct {
	Name string
	Env  string
}

func (l *Logger) SetLoggerName(name string) {
	l.Name = name
}

func (l *Logger) SetLoggerEnv(env string) {
	l.Env = env
}

func (l *Logger) PrintLog(tag string, format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("["+l.Name+" "+tag+"] "+format, values...)
}

func (l *Logger) PrintInfo(format string, values ...any) {
	l.PrintLog("info", format, values...)
}

func (l *Logger) PrintDebug(format string, values ...any) {
	l.PrintLog("debug", format, values...)
}

func (l *Logger) PrintError(format string, values ...any) {
	l.PrintLog("error", format, values...)
}

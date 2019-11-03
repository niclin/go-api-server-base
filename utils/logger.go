package utils

import (
	"fmt"
	"os"
	"time"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

var logger *Logger

type Logger struct {
	level int
}

func (ll *Logger) Println(msg string) {
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:04:05 -0700"), msg)
}

func (ll *Logger) Panic(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] "+format, v...)
	ll.Println(msg)
	os.Exit(0)
}

func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	ll.Println(msg)
}

func (ll *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	ll.Println(msg)
}

func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	ll.Println(msg)
}

func (ll *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	ll.Println(msg)
}

func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}

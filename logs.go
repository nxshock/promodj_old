package main

import "log"

type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelWarning
	LogLevelDebug
)

func Log(minLogLevel LogLevel, v ...interface{}) {
	if config == nil {
		return
	}

	if int(minLogLevel) >= config.LogLevel {
		return
	}

	log.Print(v...)
}

func Logf(minLogLevel LogLevel, f string, v ...interface{}) {
	if config == nil {
		return
	}

	if int(minLogLevel) >= config.LogLevel {
		return
	}

	log.Printf(f, v...)

}

func Logln(minLogLevel LogLevel, v ...interface{}) {
	if config == nil {
		return
	}

	if int(minLogLevel) >= config.LogLevel {
		return
	}

	log.Println(v...)
}

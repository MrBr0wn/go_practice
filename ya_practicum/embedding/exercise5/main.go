package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelError, LogLevelWarning:
		return true
	default:
		return false
	}
}

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO: ", msg)
}

func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERROR: ", msg)
}

func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN: ", msg)
}

func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("No print this message")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

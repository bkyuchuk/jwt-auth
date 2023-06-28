package common

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
)

type LogLevel int8

var CurrentLogLevel LogLevel = LevelInfo

func SetupLogFile(path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Log file can't be accessed.")
		return
	}
	log.SetOutput(file)
}

func SetLogLevel(logLevel LogLevel) {
	CurrentLogLevel = logLevel
}

func Debug(msgs ...string) {
	LogMessage(LevelDebug, msgs...)
}

func Info(msgs ...string) {
	LogMessage(LevelInfo, msgs...)
}

func Warning(msgs ...string) {
	LogMessage(LevelWarning, msgs...)
}

func Error(msgs ...string) {
	LogMessage(LevelError, msgs...)
}

func LogMessage(logLevel LogLevel, msgs ...string) {
	if logLevel < CurrentLogLevel {
		return
	}

	message := strings.Join(msgs, " ")
	log.Println("[" + logLevel.asString() + "] " + message)
}

func (logLevel LogLevel) asString() string {
	switch logLevel {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	default:
		panic(fmt.Sprintf("Log level %v doesn't exist!", logLevel))
	}
}

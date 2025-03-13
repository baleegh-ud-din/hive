package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[36m"
	ColorPurple = "\033[35m"
)

type Logger struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	WarningLog *log.Logger
	SuccessLog *log.Logger
	DebugLog   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		InfoLog:    log.New(os.Stdout, "", 0),
		ErrorLog:   log.New(os.Stdout, "", 0),
		WarningLog: log.New(os.Stdout, "", 0),
		SuccessLog: log.New(os.Stdout, "", 0),
		DebugLog:   log.New(os.Stdout, "", 0),
	}
}

func formatLog(logType string, message string, color string) string {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	return fmt.Sprintf("%s %s%s %s%s",
		timestamp,
		color,
		logType,
		message,
		ColorReset,
	)
}

func (l *Logger) Info(v ...interface{}) {
	message := fmt.Sprint(v...)
	l.InfoLog.Println(formatLog("[ INFO  ] : ", message, ColorBlue))
}

func (l *Logger) Error(v ...interface{}) {
	message := fmt.Sprint(v...)
	l.ErrorLog.Println(formatLog("[ ERROR ] : ", message, ColorRed))
}

func (l *Logger) Warning(v ...interface{}) {
	message := fmt.Sprint(v...)
	l.WarningLog.Println(formatLog("[WARNING] : ", message, ColorYellow))
}

func (l *Logger) Success(v ...interface{}) {
	message := fmt.Sprint(v...)
	l.SuccessLog.Println(formatLog("[SUCCESS] : ", message, ColorGreen))
}

func (l *Logger) Debug(v ...interface{}) {
	message := fmt.Sprint(v...)
	l.SuccessLog.Println(formatLog("[ DEBUG ] : ", message, ColorPurple))
}

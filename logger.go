package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// LogLevel type for setting log levels
type LogLevel int

// Enum values for LogLevel
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	NONE
)

var (
	levelColors = map[LogLevel]string{
		DEBUG: "\033[90m",
		INFO:  "\033[36m",
		WARN:  "\033[33m",
		ERROR: "\033[31m",
	}
	consoleColors = map[string]string{
		"reset": "\033[0m",
	}
	logLevels = map[string]LogLevel{
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"WARN":  WARN,
		"ERROR": ERROR,
		"NONE":  NONE,
	}
	defaultLogger *Logger
	once          sync.Once
)

// Logger struct represents a logger
type Logger struct {
	logLevel   LogLevel
	logFile    *os.File
	mu         sync.Mutex
	fileWriter *log.Logger
}

// initialize sets up the default logger
func initialize(level LogLevel, logFilePath string) {
	defaultLogger = &Logger{logLevel: level}
	if logFilePath != "" {
		var err error
		defaultLogger.logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("error opening log file: %v", err)
		}
		defaultLogger.fileWriter = log.New(defaultLogger.logFile, "", log.LstdFlags)
	}
}

// Setup configures the default logger settings
func Setup(level LogLevel, logFilePath string) {
	once.Do(func() {
		initialize(level, logFilePath)
	})
}

// log outputs messages using the default logger
func (l *Logger) log(module string, msg string, level LogLevel) {
	if level < l.logLevel {
		return
	}

	timeStamp := time.Now().Format(time.RFC3339)
	levelColor := levelColors[level]
	resetColor := consoleColors["reset"]

	logMessage := fmt.Sprintf("%s[%s] %s%s: %s%s", levelColor, module, timeStamp, resetColor, msg, resetColor)

	fmt.Println(logMessage)
	if l.fileWriter != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.fileWriter.Println(logMessage)
	}
}

// logf outputs formatted messages using the default logger
func (l *Logger) logf(module string, level LogLevel, format string, a ...interface{}) {
	if level < l.logLevel {
		return
	}

	timeStamp := time.Now().Format(time.RFC3339)
	levelColor := levelColors[level]
	resetColor := consoleColors["reset"]
	msg := fmt.Sprintf(format, a...)

	logMessage := fmt.Sprintf("%s[%s] %s%s: %s%s", levelColor, module, timeStamp, resetColor, msg, resetColor)

	fmt.Println(logMessage)
	if l.fileWriter != nil {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.fileWriter.Println(logMessage)
	}
}

// Info logs an information message using the default logger
func Info(module, msg string) {
	defaultLogger.log(module, msg, INFO)
}

// Warn logs a warning message using the default logger
func Warn(module, msg string) {
	defaultLogger.log(module, msg, WARN)
}

// Error logs an error message using the default logger
func Error(module, msg string) {
	defaultLogger.log(module, msg, ERROR)
}

// Debug logs a debug message using the default logger
func Debug(module, msg string) {
	defaultLogger.log(module, msg, DEBUG)
}

// Infof logs an information message with format
func Infof(module, format string, a ...interface{}) {
	defaultLogger.logf(module, INFO, format, a...)
}

// Warnf logs a warning message with format
func Warnf(module, format string, a ...interface{}) {
	defaultLogger.logf(module, WARN, format, a...)
}

// Errorf logs an error message with format
func Errorf(module, format string, a ...interface{}) {
	defaultLogger.logf(module, ERROR, format, a...)
}

// Debugf logs a debug message with format
func Debugf(module, format string, a ...interface{}) {
	defaultLogger.logf(module, DEBUG, format, a...)
}

// Close closes the default logger, releasing any resources
func Close() {
	if defaultLogger.logFile != nil {
		defaultLogger.logFile.Close()
	}
}

// SetLogLevel sets the log level for the default logger
func SetLogLevel(level LogLevel) {
	defaultLogger.logLevel = level
}

package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "DEBUG: ", logger.Flags()),
		warning: log.New(writer, "DEBUG: ", logger.Flags()),
		err:     log.New(writer, "DEBUG: ", logger.Flags()),
		writer:  writer,
	}
}

func (log *Logger) Debug(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Info(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Warning(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Error(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}

func (log *Logger) Debugf(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Infof(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Warningf(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Errorf(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}

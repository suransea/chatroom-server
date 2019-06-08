package logs

import (
	"log"
	"os"
)

func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, "[Info] ", log.LstdFlags)
}

func NewDebugLogger() *log.Logger {
	return log.New(os.Stdout, "[Debug] ", log.LstdFlags|log.Lmicroseconds)
}

func NewErrorLogger() *log.Logger {
	return log.New(os.Stderr, "[Error] ", log.LstdFlags|log.Lmicroseconds)
}

func NewLogger() (*log.Logger, *log.Logger, *log.Logger) {
	return NewInfoLogger(), NewDebugLogger(), NewErrorLogger()
}

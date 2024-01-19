package utils

import (
	"bytes"
	"log"
)

var (
	logBuf     bytes.Buffer
	infoLogger = log.New(&logBuf, "INFO: ", log.Lshortfile)
	// debugLogger = log.New(&logBuf, "DEBUG: ", log.Lshortfile)
	// warnLogger  = log.New(&logBuf, "WARN: ", log.Lshortfile)
	// errorLogger = log.New(&logBuf, "ERROR: ", log.Lshortfile)
)

func InfoLog(log string) *bytes.Buffer {
	infoLogger.Print(log)
	return &logBuf
}

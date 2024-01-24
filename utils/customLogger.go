package utils

import (
	"bytes"
	"log"
)

var (
	logBuf      bytes.Buffer
	infoLogger  = log.New(&logBuf, "INFO: ", log.Lshortfile)
	debugLogger = log.New(&logBuf, "DEBUG: ", log.Lshortfile)
	// warnLogger  = log.New(&logBuf, "WARN: ", log.Lshortfile)
	// errorLogger = log.New(&logBuf, "ERROR: ", log.Lshortfile)
)

func INFO(log string) *bytes.Buffer {
	infoLogger.Print(log)
	return &logBuf
}

func DEBUG(log string) *bytes.Buffer {
	debugLogger.Print(log)
	return &logBuf
}

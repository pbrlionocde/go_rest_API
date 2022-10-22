package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var errorLogger *log.Logger

func init() {
	errlogfile, _ := os.Create("logs/error.log")
	errorLogger = log.New(io.MultiWriter(errlogfile), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetErrorLogger() *log.Logger {
	return errorLogger
}

func ErrorToJSON(err error) map[string]string {
	errors := strings.Split(err.Error(), "\n")
	m := make(map[string]string)
	for counter, e := range errors {
		fmt.Println(e)
		parts := strings.Split(e, ": ")
		key := fmt.Sprintf("%s_%d", parts[0], counter)
		m[key] = parts[1]
	}
	return m
}

package utils

import (
	"bytes"
	"errors"
	"log"
)

func PrintLog(logType string, message string) {
	var buffer bytes.Buffer
	switch logType {
		case "error":
			buffer.WriteString("\u001B[31m[ERROR] ")
		case "warning":
			buffer.WriteString("\u001B[33m[WARNING] ")
		case "notice":
			buffer.WriteString("\u001B[36m[NOTICE] ")
		case "debug":
		default:
			buffer.WriteString("\u001B[35m[ERROR] ")
	}
	buffer.WriteString(message)
	buffer.WriteString("\u001B[0m")
	log.Print(buffer.String())
}

func FatalError(message string) {
	var buffer bytes.Buffer
	buffer.WriteString("\u001B[31m[ERROR] ")
	buffer.WriteString(message)
	buffer.WriteString("\u001B[0m")
	log.Fatal(buffer.String())
}

func CreateError(message string) error {
	return errors.New(message)
}
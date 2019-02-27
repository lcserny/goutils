package goutils

import (
	"github.com/pkg/errors"
	"log"
)

const NO_MESSAGE = "---"

func LogFatal(err error) {
	if err != nil {
		log.Fatalf("ERROR: %+v\n\n", errors.Wrap(err, NO_MESSAGE))
	}
}

func LogFatalWithMessage(message string, err error) {
	if err != nil {
		log.Fatalf("ERROR: %+v\n\n", errors.Wrap(err, message))
	}
}

func LogError(err error) {
	if err != nil {
		log.Printf("ERROR: %+v\n\n", errors.Wrap(err, NO_MESSAGE))
	}
}

func LogErrorWithMessage(message string, err error) {
	if err != nil {
		log.Printf("ERROR: %+v\n\n", errors.Wrap(err, message))
	}
}

func LogInfo(message string) {
	log.Printf("INFO: %s\n", message)
}

func LogWarning(message string) {
	log.Printf("WARN: %s\n", message)
}

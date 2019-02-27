package goutils

import (
	"github.com/pkg/errors"
	"log"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatalf("ERROR: %+v\n\n", errors.Wrap(err, ""))
	}
}

func LogFatalWithMessage(message string, err error) {
	if err != nil {
		log.Fatalf("ERROR: %s, %+v\n\n", message, errors.Wrap(err, ""))
	}
}

func LogError(err error) {
	if err != nil {
		log.Printf("ERROR: %+v\n\n", errors.Wrap(err, ""))
	}
}

func LogErrorWithMessage(message string, err error) {
	if err != nil {
		log.Printf("ERROR: %s, %+v\n\n", message, errors.Wrap(err, ""))
	}
}

func LogInfo(message string) {
	log.Printf("INFO: %s\n", message)
}

func LogWarning(message string) {
	log.Printf("WARN: %s\n", message)
}

package goutils

import (
	"github.com/juju/errors"
	"log"
)

func LogFatal(e error) {
	if e != nil {
		log.Fatalf("ERROR: %s", errors.Trace(e))
	}
}

func LogFatalWithMessage(message string, e error) {
	if e != nil {
		log.Fatalf("ERROR: %s\n%s", message, errors.Trace(e))
	}
}

func LogError(err error) {
	if err != nil {
		log.Printf("ERROR: %s\n", errors.Trace(err))
	}
}

func LogErrorWithMessage(message string, err error) {
	if err != nil {
		log.Printf("ERROR: %s: %s\n", message, errors.Trace(err))
	}
}

func LogInfo(message string) {
	log.Printf("INFO: %s\n", message)
}

func LogWarning(message string) {
	log.Printf("WARN: %s\n", message)
}

package goutils

import (
	"github.com/juju/errors"
	"log"
)

func LogFatal(e error) {
	if e != nil {
		log.Fatalf("ERROR: %#v", errors.Trace(e))
	}
}

func LogFatalWithMessage(e error, message string) {
	if e != nil {
		log.Fatalf("ERROR: %s\n%#v", message, errors.Trace(e))
	}
}

func LogError(message string, err error) {
	if err != nil {
		log.Printf("ERROR: %s: %#v\n", message, errors.Trace(err))
	}
}

func LogInfo(message string) {
	log.Printf("INFO: %s\n", message)
}

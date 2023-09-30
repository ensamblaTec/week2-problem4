package utils

import (
	"errors"
	"log"
)

var (
	errFileNotFound = errors.New("cannot find files")
)

func PrintErrorMessage(errMethod string, err error) bool {
	if err != nil {
		log.Printf("[%s]:%s", errMethod, err.Error())
		return false
	}
	return true
}

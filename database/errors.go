package database

import "errors"

var (
	errCannotOpenFile        = errors.New("cannot open file")
	errCannotCreateFile      = errors.New("cannot create file")
	errCannotCreateTempFile  = errors.New("cannot create temporal file")
	errCannotConvertStrToInt = errors.New("cannot convert str value to int")
	errCannotWriteFile       = errors.New("cannot write file")
)

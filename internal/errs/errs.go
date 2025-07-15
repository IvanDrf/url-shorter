package errs

import (
	"fmt"
)

const (
	errConfig = -1

	errDatabaseConn   = -2
	errDatabaseCreate = -3
	errSQL            = -4
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"err"`
}

func (this *Error) Error() string {
	return fmt.Sprintf("code: %v, err: %s", this.Code, this.Msg)
}

// error of env file
func InvalidEnv() Error {
	return Error{Code: errConfig, Msg: "can't read env file"}
}

// error of logger level
func InvalidLevel() Error {
	return Error{Code: errConfig, Msg: "can't set logger level"}
}

// error of connection to database
func InvalidDBConnection() Error {
	return Error{Code: errDatabaseConn, Msg: "can't connect to database"}
}

// error of creation database
func InvalidDBCreation() Error {
	return Error{Code: errDatabaseCreate, Msg: "can't create database"}
}

// error of invalid sql request
func InvalidSQL(msg string) Error {
	return Error{Code: errSQL, Msg: msg}
}

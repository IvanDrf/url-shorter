package errs

import (
	"fmt"
)

const (
	errConfig = -1

	errDatabaseConn   = -2
	errDatabaseCreate = -3
	errSQL            = -4

	errInvalidURL  = -5
	errStartServer = -10
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"err"`
}

func (this Error) Error() string {
	return fmt.Sprintf("code: %v, err: %s", this.Code, this.Msg)
}

// error of env file
func InvalidEnv() error {
	return Error{Code: errConfig, Msg: "can't read env file"}
}

// error of logger level
func InvalidLevel() error {
	return Error{Code: errConfig, Msg: "can't set logger level"}
}

func InvalidStart() error {
	return Error{Code: errStartServer, Msg: "can't start server"}
}

// error of connection to database
func InvalidDBConnection() error {
	return Error{Code: errDatabaseConn, Msg: "can't connect to database"}
}

// error of creation database
func InvalidDBCreation() error {
	return Error{Code: errDatabaseCreate, Msg: "can't create database"}
}

// error of invalid sql request
func InvalidSQL(msg string) error {
	return Error{Code: errSQL, Msg: msg}
}

// error of invalid source url
func InvalidURL() error {
	return Error{Code: errInvalidURL, Msg: "invalid source url"}
}

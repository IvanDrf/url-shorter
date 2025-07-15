package errs

import (
	"fmt"
	"net/http"
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

func (this Error) Byte() []byte {
	return []byte(this.Error())
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

// error of invalid media type in request
func InvalidMediaType() error {
	return Error{Code: http.StatusUnsupportedMediaType, Msg: "invalid media type, wants json"}
}

// error of invalid json in request - cant unmarshal in models.Request
func InvalidJSON() error {
	return Error{Code: http.StatusBadRequest, Msg: "invalid storage of json in request"}
}

// error of invalid short link
func InvalidShortURL() error {
	return Error{Code: http.StatusNotFound, Msg: "can't find url in database"}
}

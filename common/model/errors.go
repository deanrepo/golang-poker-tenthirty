package model

import (
	"errors"
	"log"
)

var (
	// ErrUserNotExists represents user not exists
	ErrUserNotExists = errors.New("user not exists")
	// ErrUserExists represents user exists
	ErrUserExists = errors.New("user exists")
	// ErrWrongPwd represents wrong password
	ErrWrongPwd = errors.New("wrong password")
)

// LogErr logs the error information with specified information.
func LogErr(err error, where string) {
	log.Printf(where+"-> err: %v\n", err)
}

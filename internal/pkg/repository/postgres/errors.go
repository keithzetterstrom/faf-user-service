package db

import (
	"errors"

	"github.com/lib/pq"
)

const ErrCodeConcurrentUpdate = "40001"

func GetPostgresErrorCode(err error) string {
	code := getPGXCode(err)
	if code != "" {
		return code
	}

	return getPQCode(err)
}

func getPQCode(err error) string {
	var pqError *pq.Error
	if errors.As(err, &pqError) {
		return string(pqError.Code)
	}

	return ""
}

func getPGXCode(err error) string {
	var pgxError interface {
		SQLState() string
	}
	if errors.As(err, &pgxError) {
		return pgxError.SQLState()
	}

	return ""
}

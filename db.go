package gosp

import (
	"database/sql"
	"github.com/athlum/gosp/types"
)

const ErrDBPluged = types.Error("db that pluged before was replaced.")

var (
	db     *sql.DB
	logger Logger
)

func PlugDB(v *sql.DB) (err error) {
	if db != nil {
		return ErrDBPluged
	}
	db = v
	return nil
}

func SetLogger(l Logger) {
	logger = l
}

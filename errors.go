package mydb

import "errors"

var (
	ErrAllReadreplicaDied = errors.New("all readreadreplica died")
	ErrMasterDied         = errors.New("master died")
)

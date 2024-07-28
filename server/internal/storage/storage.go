package storage

import "errors"

var (
    ErrRecordNotFound = errors.New("record not found")
)

type Record struct {
    Title string
    Note string
    Author string
}

type Storage interface {
    SaveRecord (rec Record, alias string) (int64, error)
    GetRecord(alias string) (Record, error)
}

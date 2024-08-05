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
    SaveRecord (rec Record, alias string) (error)
    GetRecord(alias string) (Record, error)
}

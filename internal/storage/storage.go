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

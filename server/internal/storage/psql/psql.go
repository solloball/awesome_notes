package psql

import (
    "database/sql"
    "fmt"
    "errors"

    _ "github.com/lib/pq" // add this
    "github.com/solloball/aws_note/internal/storage"
)

type Storage struct {
    db *sql.DB
}

func New(storagePath string) (*Storage, error) {
    const op = "storage.psql.New"

    db, err := sql.Open("postgres", storagePath)
    if err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    stmt, err := db.Prepare(
        `CREATE TABLE IF NOT EXISTS record(
            id INTEGER PRIMARY KEY,
            title TEXT NOT NULL,
            note TEXT NOT NULL,
            author TEXT NOT NULL,
            alias TEXT NOT NULL);
        CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
        `)
    if err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    if _, err := stmt.Exec(); err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }


    return &Storage{db: db}, nil
}


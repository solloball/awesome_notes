package sqlite

import (
    "database/sql"
    "fmt"

    "github.com/mattn/go-sqlite3"
    "github.com/solloball/aws_tg/internal/storage"
)

type Storage struct {
    db *sql.DB
}

func New(storagePath string) (*Storage, error) {
    const op = "storage.sqlite.New"

    db, err := sql.Open("sqlite3", storagePath)
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

func (s *Storage) saveRecord(
    title string,
    body string,
    author string,
    alias string) (int64, error) {
        const op = "storage.sqlite.saveRecord"

        stmt, err := s.db.Prepare(
            "INSERT INTO record(title, note,  author, alias) VALUES(?, ?, ?, ?)")
        if err != nil {
            return 0, fmt.Errorf("%s: %w", op, err)
        }

        res, err := stmt.Exec(title, body, author, alias)
        if err != nil {
            // TODO: refactor this
            if sqliteErr, ok := err.(sqlite3.Error);
                ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
                    return 0, fmt.Errorf("%s: %w", op, storage.ErrUrlExist)
            }
            return 0, fmt.Errorf("%s: %w", op, err)
        }

        id, err := res.LastInsertId()
        if err != nil {
            return 0, fmt.Errorf("%s: %w", op, err)
        }
        
        return id, nil
    }


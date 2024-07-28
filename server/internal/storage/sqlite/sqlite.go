package sqlite

import (
    "database/sql"
    "fmt"
    "errors"

    _ "github.com/mattn/go-sqlite3"
    "github.com/solloball/aws_note/internal/storage"
)

type StorageSQLite struct {
    db *sql.DB
}

func New(storagePath string) (storage.Storage, error) {
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


    return &StorageSQLite{db: db}, nil
}

func (s *StorageSQLite) SaveRecord (rec storage.Record, alias string) (int64, error) {
        const op = "storage.sqlite.saveRecord"

        stmt, err := s.db.Prepare(
            "INSERT INTO record(title, note,  author, alias) VALUES(?, ?, ?, ?)")
        if err != nil {
            return 0, fmt.Errorf("%s: %w", op, err)
        }

        res, err := stmt.Exec(rec.Title, rec.Note, rec.Author, alias)
        if err != nil {
            return 0, fmt.Errorf("%s: %w", op, err)
        }

        id, err := res.LastInsertId()
        if err != nil {
            return 0, fmt.Errorf("%s: %w", op, err)
        }
        
        return id, nil
    }

func (s *StorageSQLite) GetRecord(alias string) (storage.Record, error) {
    const op = "storage.sqlite.GetRecord"


    stmt, err := s.db.Prepare("SELECT author, title, note FROM record WHERE alias = ?")
    if err != nil {
        return storage.Record{}, fmt.Errorf("%s: %w", op, err)
    }

    var res storage.Record
    err = stmt.QueryRow(alias).Scan(&res.Author, &res.Title, &res.Note)
    if errors.Is(err, sql.ErrNoRows) {
        return storage.Record{}, storage.ErrRecordNotFound
    }
    if err != nil {
        return storage.Record{}, fmt.Errorf("%s: %w", op, err)
    }

    return res, nil
}

// TODO:: implement this
// func (s *Storage) DeleteRecord(alias string) (storage.Record, error)

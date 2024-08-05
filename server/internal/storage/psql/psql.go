package psql

import (
    "database/sql"
    "fmt"
    "errors"

    _ "github.com/lib/pq" // add this
    "github.com/solloball/aws_note/internal/storage"
)

type StoragePSQL struct {
    db *sql.DB
}

func New(storagePath string) (storage.Storage, error) {
    const op = "storage.psql.New"

    db, err := sql.Open("postgres", storagePath)
    if err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    stmt, err := db.Prepare(
        `CREATE TABLE IF NOT EXISTS record(
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            note TEXT NOT NULL,
            author TEXT NOT NULL,
            alias TEXT NOT NULL);
        `)
    if err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    if _, err := stmt.Exec(); err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    stmt, err = db.Prepare(
        `CREATE INDEX IF NOT EXISTS idx_alias ON record(alias);`)
    if err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    if _, err := stmt.Exec(); err != nil {
        return nil, fmt.Errorf("%s: :%w", op, err)
    }

    return &StoragePSQL{db: db}, nil
}

func (s *StoragePSQL) SaveRecord (rec storage.Record, alias string) (error) {
        const op = "storage.psql.saveRecord"

        stmt, err := s.db.Prepare(
            "INSERT INTO record(title, note, author, alias) VALUES($1, $2, $3, $4)")
        if err != nil {
            return fmt.Errorf("%s: %w", op, err)
        }

        _ , err = stmt.Exec(rec.Title, rec.Note, rec.Author, alias)
        if err != nil {
            return fmt.Errorf("%s: %w", op, err)
        }

        return nil
    }

func (s *StoragePSQL) GetRecord(alias string) (storage.Record, error) {
    const op = "storage.psql.GetRecord"


    stmt, err := s.db.Prepare("SELECT author, title, note FROM record WHERE alias = $1")
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


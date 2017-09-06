package sql

import (
	"database/sql"
)

type SQLStorageTx interface {
	Query(query string, args ...interface{}) (SQLStorageRows, error)
	QueryRow(query string, args ...interface{}) SQLStorageRow
	Exec(query string, args ...interface{}) (SQLStorageResult, error)

	Commit() error
	Rollback() error
}

type SQLStorageResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type SQLStorageRow interface {
	Scan(dest ...interface{}) error
}

type SQLStorageRows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type SQLStorage interface {
	Query(query string, args ...interface{}) (SQLStorageRows, error)
	QueryRow(query string, args ...interface{}) SQLStorageRow
	Exec(query string, args ...interface{}) (SQLStorageResult, error)

	Begin() (SQLStorageTx, error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) SQLStorage {
	return &storage{
		db: db,
	}
}

func (s *storage) Begin() (SQLStorageTx, error) {
	tx, err := s.db.Begin()
	return &txWrap{tx}, err
}

func (s *storage) Query(query string, args ...interface{}) (SQLStorageRows, error) {
	return s.db.Query(query, args...)
}

func (s *storage) QueryRow(query string, args ...interface{}) SQLStorageRow {
	return s.db.QueryRow(query, args...)
}

func (s *storage) Exec(query string, args ...interface{}) (SQLStorageResult, error) {
	return s.db.Exec(query, args...)
}

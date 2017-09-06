package sql

import (
	"database/sql"
)

type txWrap struct {
	tx *sql.Tx
}

func (t *txWrap) Query(query string, args ...interface{}) (SQLStorageRows, error) {
	return t.tx.Query(query, args...)
}
func (t *txWrap) QueryRow(query string, args ...interface{}) SQLStorageRow {
	return t.tx.QueryRow(query, args...)
}
func (t *txWrap) Exec(query string, args ...interface{}) (SQLStorageResult, error) {
	return t.tx.Exec(query, args...)
}

func (t *txWrap) Commit() error {
	return t.tx.Commit()
}

func (t *txWrap) Rollback() error {
	return t.tx.Rollback()
}

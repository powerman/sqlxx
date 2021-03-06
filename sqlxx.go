// Package sqlxx extends github.com/jmoiron/sqlx.
package sqlxx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// DB is a wrapper around sqlx.DB.
type DB struct {
	*sqlx.DB
}

// NewDB returns a new sqlxx DB wrapper for a pre-existing *sqlx.DB.
func NewDB(db *sqlx.DB) *DB {
	return &DB{DB: db}
}

// NamedIn expands slice values in arg returning the modified query string
// and a new arg list that can be executed by a database.
func (db *DB) NamedIn(queryIn string, arg interface{}) (query string, args []interface{}, err error) {
	query, args, err = sqlx.Named(queryIn, arg)
	if err == nil {
		query, args, err = sqlx.In(query, args...)
	}
	return db.Rebind(query), args, err
}

// NamedGetContext using this DB.
// Any named placeholder parameters are replaced with fields from arg.
// An error is returned if the result set is empty.
func (db *DB) NamedGetContext(ctx Ctx, dest interface{}, query string, arg interface{}) error {
	query, args, err := db.BindNamed(query, arg)
	if err == nil {
		err = db.GetContext(ctx, dest, query, args...)
	}
	return err
}

// NamedSelectContext using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func (db *DB) NamedSelectContext(ctx Ctx, dest interface{}, query string, arg interface{}) error {
	query, args, err := db.BindNamed(query, arg)
	if err == nil {
		err = db.SelectContext(ctx, dest, query, args...)
	}
	return err
}

// BeginTxx begins a transaction and returns an *sqlxx.Tx instead of an *sqlx.Tx.
func (db *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, err := db.DB.BeginTxx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx}, nil
}

// Tx is an sqlxx wrapper around sqlx.Tx with extra functionality.
type Tx struct {
	*sqlx.Tx
}

// NamedIn expands slice values in arg returning the modified query string
// and a new arg list that can be executed within a transaction.
func (tx *Tx) NamedIn(queryIn string, arg interface{}) (query string, args []interface{}, err error) {
	query, args, err = sqlx.Named(queryIn, arg)
	if err == nil {
		query, args, err = sqlx.In(query, args...)
	}
	return tx.Rebind(query), args, err
}

// NamedGetContext within a transaction.
// Any named placeholder parameters are replaced with fields from arg.
// An error is returned if the result set is empty.
func (tx *Tx) NamedGetContext(ctx Ctx, dest interface{}, query string, arg interface{}) error {
	query, args, err := tx.BindNamed(query, arg)
	if err == nil {
		err = tx.GetContext(ctx, dest, query, args...)
	}
	return err
}

// NamedSelectContext within a transaction.
// Any named placeholder parameters are replaced with fields from arg.
func (tx *Tx) NamedSelectContext(ctx Ctx, dest interface{}, query string, arg interface{}) error {
	query, args, err := tx.BindNamed(query, arg)
	if err == nil {
		err = tx.SelectContext(ctx, dest, query, args...)
	}
	return err
}

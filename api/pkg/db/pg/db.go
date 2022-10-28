package pg

import (
	"context"
	"database/sql"
)

// go_coreDB wraps the *sql.DB provided
type go_coreDB struct {
	*sql.DB
	info InstanceInfo
}

// BeginTx begins a transaction with the database in receiver and returns a Transactor
func (i go_coreDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transactor, error) {
	return i.DB.BeginTx(ctx, opts)
}

// ExecContext wraps the base connector
func (i go_coreDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return i.DB.ExecContext(ctx, query, args...)
}

// QueryContext wraps the base connector
func (i go_coreDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return i.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext wraps the base connector
func (i go_coreDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return i.DB.QueryRowContext(ctx, query, args...)
}

// InstanceInfo returns info about the DB
func (i go_coreDB) InstanceInfo() InstanceInfo {
	return i.info
}

// go_coreTx wraps the Transactor provided
type go_coreTx struct {
	Transactor
	info InstanceInfo
	ctx  context.Context
}

// ExecContext wraps the base connector
func (i go_coreTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return i.Transactor.ExecContext(ctx, query, args...)
}

// QueryContext wraps the base connector
func (i go_coreTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return i.Transactor.QueryContext(ctx, query, args...)
}

// QueryRowContext wraps the base connector
func (i go_coreTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return i.Transactor.QueryRowContext(ctx, query, args...)
}

// Commit commits the transaction
func (i go_coreTx) Commit() error {
	return i.Transactor.Commit()
}

// Rollback aborts the transaction
func (i go_coreTx) Rollback() error {
	return i.Transactor.Rollback()
}

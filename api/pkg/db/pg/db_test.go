package pg

import (
	"context"
	"database/sql"
	"testing"

	"go_core/api/pkg/env"

	"github.com/stretchr/testify/require"
)

func Testgo_coreDB(t *testing.T) {

	pool, err := sql.Open("postgres", env.GetAndValidateF("DB_URL"))
	require.NoError(t, err)

	db := &go_coreDB{DB: pool}

	_, err = db.Exec("DROP TABLE IF EXISTS instrumented_test_transactions")
	require.NoError(t, err)
	_, err = db.Exec("CREATE TABLE instrumented_test_transactions (id SERIAL PRIMARY KEY)")
	require.NoError(t, err)
	defer func() {
		_, err = db.Exec("DROP TABLE IF EXISTS instrumented_test_transactions")
		require.NoError(t, err)
	}()

	_, err = db.QueryContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	require.NoError(t, err)

	var p interface{}
	row := db.QueryRowContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	err = row.Scan(&p)
	require.Equal(t, sql.ErrNoRows, err)

	row = db.QueryRow("SELECT * FROM instrumented_test_transactions")
	err = row.Scan(&p)
	require.Equal(t, sql.ErrNoRows, err)

	_, err = db.ExecContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	require.NoError(t, err)

	tx, err := db.Begin()
	require.NoError(t, err)
	require.NoError(t, tx.Rollback())
}

func Testgo_coreTx(t *testing.T) {

	pool, err := sql.Open("postgres", env.GetAndValidateF("DB_URL"))
	require.NoError(t, err)

	db := &go_coreDB{DB: pool}

	_, err = db.Exec("DROP TABLE IF EXISTS instrumented_test_transactions")
	require.NoError(t, err)
	_, err = db.Exec("CREATE TABLE instrumented_test_transactions (id SERIAL PRIMARY KEY)")
	require.NoError(t, err)
	defer func() {
		_, err = db.Exec("DROP TABLE IF EXISTS instrumented_test_transactions")
		require.NoError(t, err)
	}()

	transactor, err := db.BeginTx(context.Background(), nil)
	require.NoError(t, err)
	tx := &go_coreTx{Transactor: transactor}

	defer func() {
		_ = tx.Rollback()
	}()
	_, err = tx.QueryContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	require.NoError(t, err)

	var p interface{}
	row := tx.QueryRowContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	err = row.Scan(&p)
	require.Equal(t, sql.ErrNoRows, err)

	row = tx.QueryRow("SELECT * FROM instrumented_test_transactions")
	err = row.Scan(&p)
	require.Equal(t, sql.ErrNoRows, err)

	_, err = tx.ExecContext(context.Background(), "SELECT * FROM instrumented_test_transactions")
	require.NoError(t, err)

	_, err = tx.Exec("SELECT * FROM instrumented_test_transactions")
	require.NoError(t, err)

	require.NoError(t, tx.Commit())
}

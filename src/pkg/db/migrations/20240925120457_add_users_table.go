package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUsersTable, downAddUsersTable)
}

func upAddUsersTable(ctx context.Context, tx *sql.Tx) error {
	query := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE users (
        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
    `
    _, err := tx.ExecContext(ctx, query)
    return err
}

func downAddUsersTable(ctx context.Context, tx *sql.Tx) error {
	 query := `DROP TABLE IF EXISTS users;`
    _, err := tx.ExecContext(ctx, query)
    return err
}

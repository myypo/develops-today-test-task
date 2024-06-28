package sql

import "database/sql"

type SqlDbProvider interface {
	Conn() *sql.DB
	MigrateUp() error
}

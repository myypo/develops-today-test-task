package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sca/internal/config"

	provider "sca/internal/provider/db/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type postgresProvider struct {
	conn *sql.DB
}

func NewPostgresProvider(
	config *config.Config,
) (provider.SqlDbProvider, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Europe/London",
		config.DB.Host,
		config.DB.Port,
		config.DB.UserName,
		config.DB.Password,
		config.DB.Name,
		config.DB.SslMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &postgresProvider{conn}, nil
}

func (p *postgresProvider) Conn() *sql.DB {
	return p.conn
}

func (p *postgresProvider) MigrateUp() error {
	goose.SetDialect("postgres")
	return goose.RunContext(context.Background(), "up", p.conn, "./migration")
}

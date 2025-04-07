package database

import (
	"context"
	"fmt"

	"github.com/his-vita/patients-service/internal/config"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	PgCon *pgx.Conn
}

func NewPostgresConnect(dbCfg *config.Db) *Postgres {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName, dbCfg.SSLMode)

	pgCon, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	if err := pgCon.Ping(context.Background()); err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	return &Postgres{
		PgCon: pgCon,
	}
}

func (p *Postgres) CloseConnect() {
	p.PgCon.Close(context.Background())
}

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/his-vita/patients-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

const timeout = 30 * time.Second

type PgContext struct {
	Pool    *pgxpool.Pool
	Context context.Context
	cancel  context.CancelFunc
}

func NewPostgresConnect(dbCfg *config.Db) *PgContext {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	return &PgContext{
		Pool:    pool,
		Context: ctx,
		cancel:  cancel,
	}
}

func (p *PgContext) Close() {
	p.cancel()
	p.Pool.Close()
}

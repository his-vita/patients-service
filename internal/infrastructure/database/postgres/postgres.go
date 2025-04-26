package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/his-vita/patients-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgContext struct {
	Pool        *pgxpool.Pool
	connTimeout time.Duration
}

func NewPostgresConnect(dbCfg *config.Db) (*PgContext, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &PgContext{
		Pool:        pool,
		connTimeout: 30 * time.Second,
	}, nil
}

func (pg *PgContext) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func (pg *PgContext) DefaultTimeoutCtx() (context.Context, context.CancelFunc) {
	return pg.WithTimeout(pg.connTimeout)
}

func (p *PgContext) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

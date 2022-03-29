package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	faflogger "github.com/keithzetterstrom/faf-user-service/utils/logger"
)

type dbWrap struct {
	dbPool *pgxpool.Pool
	logger faflogger.Logger
}

type Handler interface {
	DBPool() *pgxpool.Pool
	GetTx(ctx context.Context, level pgx.TxIsoLevel) (pgx.Tx, error)
	Logger() faflogger.Logger
	Close()
}

func New(ctx context.Context, cfg Config, logger faflogger.Logger) (Handler, error) {
	conf := cfg.withDefaults()

	pool, err := pgxpool.Connect(ctx, conf.DSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	wrap := &dbWrap{
		dbPool: pool,
		logger: logger,
	}

	return wrap, nil
}

func (d *dbWrap) DBPool() *pgxpool.Pool {
	return d.dbPool
}

func (d *dbWrap) GetTx(ctx context.Context, level pgx.TxIsoLevel) (pgx.Tx, error) {
	tx, err := d.dbPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: level})
	if err != nil {
		return nil, fmt.Errorf("pool begin tx failed: %w", err)
	}

	return tx, nil
}

func (d *dbWrap) Logger() faflogger.Logger {
	return d.logger
}

func (d *dbWrap) Close() {
	d.dbPool.Close()
}

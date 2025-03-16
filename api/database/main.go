package database

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/celso-alexandre/api/environment"
	"github.com/celso-alexandre/api/query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomDB struct {
	pool    *pgxpool.Pool
	queries *query.Queries
}

var (
	pool      *pgxpool.Pool
	queryPool *query.Queries
	muPool    sync.Mutex
)

var cfg = environment.GetConfig()

func init() {
	var err error
	if !strings.Contains(cfg.DatabaseURL, "application_name=") {
		if strings.Contains(cfg.DatabaseURL, "?") {
			cfg.DatabaseURL += "&"
		} else {
			cfg.DatabaseURL += "?"
		}
		cfg.DatabaseURL += "application_name=web"
	}
	pool, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	queryPool = query.New(pool)
}

func NewCustomDB() (*CustomDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := pool.Ping(ctx)
	if err != nil {
		muPool.Lock()
		defer muPool.Unlock()
		fmt.Println("Reconnecting to the database")
		pool, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
		if err != nil {
			return nil, err
		}
		queryPool = query.New(pool)
	}

	return &CustomDB{
		pool:    pool,
		queries: queryPool,
	}, nil
}

// NewQuery returns a new query instance
// func (db *CustomDB) NewQuery(ctx context.Context) *query.Queries {
// 	// return db.queries
// 	return query.New(db.pool)
// }

// txWithContext runs a pre-script in order to set user context configuration
func (db *CustomDB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (db *CustomDB) NewTxQuery(ctx context.Context) (pgx.Tx, *query.Queries, error) {
	tx, err := db.BeginTx(ctx)
	if err != nil {
		return nil, nil, err
	}
	query := db.queries.WithTx(tx)
	return tx, query, nil
}

func (db *CustomDB) UseTxQuery(tx pgx.Tx) *query.Queries {
	return db.queries.WithTx(tx)
}

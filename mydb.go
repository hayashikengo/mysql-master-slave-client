package mydb

import (
	"context"
	"database/sql"
	"time"
)

type FallbackType int

const (
	None FallbackType = iota
	UseMaster
)

const (
	DefaultHealthCheckIntervalMilli = 5000
	DefaultBalanceAlgorithm         = Random
	DefaultFallbackType             = UseMaster
)

type DB struct {
	ctx            context.Context
	cancel         context.CancelFunc
	master         *sql.DB
	masterHealth   error
	readreplicas   []*sql.DB
	readDbBalancer *dbBalancer
	fallbackType   FallbackType
}

func New(master *sql.DB, readreplicas ...*sql.DB) *DB {
	ctx := context.Background()
	db := &DB{
		ctx:            ctx,
		master:         master,
		readreplicas:   readreplicas,
		readDbBalancer: NewDbBalancer(ctx, readreplicas),
		fallbackType:   DefaultFallbackType,
	}

	// setup context
	db.ctx, db.cancel = context.WithCancel(ctx)

	if db.master != nil {
		// master health check
		db.masterHealthCheck()

		// run master health check worker
		go db.masterHealthCheckWorker()
	}

	return db
}

func (db *DB) masterHealthCheck() {
	db.masterHealth = db.master.Ping()
}

func (db *DB) masterHealthCheckWorker() {
	for {
		select {
		case <-db.ctx.Done():
			return
		default:
			time.Sleep(DefaultHealthCheckIntervalMilli * time.Millisecond)
			db.masterHealthCheck()
		}
	}
}

func (db *DB) getMaster() (*sql.DB, error) {
	if db.masterHealth == nil {
		return db.master, nil
	} else {
		return nil, ErrMasterDied
	}
}

func (db *DB) getReadReplica() (*sql.DB, error) {
	if db.readDbBalancer.IsAlive() {
		return db.readDbBalancer.Get(), nil
	} else {
		// Fallback. Use master for read, if all replica died
		switch db.fallbackType {
		case UseMaster:
			if db.masterHealth == nil {
				return db.master, nil
			} else {
				return nil, ErrMasterDied
			}
		default:
			return nil, ErrAllReadreplicaDied
		}
	}
}

func (db *DB) allDbList() []*sql.DB {
	return append(db.readreplicas, db.master)
}

func (db *DB) Ping() error {
	allDbList := db.allDbList()
	return goFuncs(len(allDbList), func(i int) error {
		return allDbList[i].Ping()
	})
}

func (db *DB) PingContext(ctx context.Context) error {
	allDbList := db.allDbList()
	return goFuncs(len(allDbList), func(i int) error {
		return allDbList[i].PingContext(ctx)
	})
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d, err := db.getReadReplica()
	if err != nil {
		return nil, err
	}
	return d.Query(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	d, err := db.getReadReplica()
	if err != nil {
		return nil, err
	}

	return d.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	d, err := db.getReadReplica()
	if err != nil {
		return nil
	}

	return d.QueryRow(query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	d, err := db.getReadReplica()
	if err != nil {
		return nil
	}

	return d.QueryRowContext(ctx, query, args...)
}

func (db *DB) Begin() (*sql.Tx, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.Begin()
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.BeginTx(ctx, opts)
}

func (db *DB) Close() error {
	// stop master health check
	db.cancel()
	// stop readDbBalancer
	db.readDbBalancer.Destroy()

	allDbList := db.allDbList()
	return goFuncs(len(allDbList), func(i int) error {
		return allDbList[i].Close()
	})
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.Exec(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.ExecContext(ctx, query, args...)
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.Prepare(query)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	d, err := db.getMaster()
	if err != nil {
		return nil, err
	}

	return d.PrepareContext(ctx, query)
}

func (db *DB) SetConnMaxLifetime(d time.Duration) {
	allDbList := db.allDbList()
	for i := range allDbList {
		allDbList[i].SetConnMaxLifetime(d)
	}
}

func (db *DB) SetMaxIdleConns(n int) {
	allDbList := db.allDbList()
	for i := range allDbList {
		allDbList[i].SetMaxIdleConns(n)
	}
}

func (db *DB) SetMaxOpenConns(n int) {
	allDbList := db.allDbList()
	for i := range allDbList {
		allDbList[i].SetMaxOpenConns(n)
	}
}

func (db *DB) GetHealthCheckIntervalMilli() int {
	return db.readDbBalancer.GetHealthCheckIntervalMilli()
}

func (db *DB) SetHealthCheckIntervalMilli(i int) {
	db.readDbBalancer.SetHealthCheckIntervalMilli(i)
}

func (db *DB) GetBalanceAlgorithm() BalanceAlgorithm {
	return db.readDbBalancer.GetBalanceAlgorithm()
}

func (db *DB) SetBalanceAlgorithm(balanceAlgorithm BalanceAlgorithm) {
	db.readDbBalancer.SetBalanceAlgorithm(balanceAlgorithm)
}

func (db *DB) GetFallbackType() FallbackType {
	return db.fallbackType
}

func (db *DB) SetFallbackType(fallbackType FallbackType) {
	db.fallbackType = fallbackType
}

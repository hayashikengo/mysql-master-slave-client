package mydb

import (
	"context"
	"database/sql"
	"time"
)

var _ interface {
	IsAlive() bool
	Get() *sql.DB
	Destroy()
	GetHealthCheckIntervalMilli() int
	SetHealthCheckIntervalMilli(i int)
	GetBalanceAlgorithm() BalanceAlgorithm
	SetBalanceAlgorithm(balanceAlgorithm BalanceAlgorithm)
} = NewDbBalancer(context.Background(), []*sql.DB{})

type BalanceAlgorithm int

const (
	RoundRobin BalanceAlgorithm = iota
	Random
)

type dbBalancer struct {
	ctx                      context.Context
	cancel                   context.CancelFunc
	dbs                      []*sql.DB
	availableDbs             *dbList
	isMulti                  bool
	healthCheckIntervalMilli int
	balanceAlgorithm         BalanceAlgorithm
}

func NewDbBalancer(ctx context.Context, dbs []*sql.DB) *dbBalancer {
	d := &dbBalancer{
		dbs:                      dbs,
		availableDbs:             NewDbList(),
		isMulti:                  len(dbs) > 1,
		healthCheckIntervalMilli: DefaultHealthCheckIntervalMilli,
		balanceAlgorithm:         DefaultBalanceAlgorithm,
	}

	// setup context
	d.ctx, d.cancel = context.WithCancel(ctx)

	// health check
	d.healthCheck()

	// run health check worker
	go d.healthCheckWorker()

	return d
}

func (d *dbBalancer) healthCheck() {
	// OPTIMIZE: allocate times
	// Not critical, Because this method called by only health check.
	availableDbs := make([]*sql.DB, 0)
	for i := range d.dbs {
		db := d.dbs[i]
		if db.Ping() == nil {
			availableDbs = append(availableDbs, db)
		}
	}

	d.availableDbs.Replace(availableDbs)
}

func (d *dbBalancer) healthCheckWorker() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			time.Sleep(DefaultHealthCheckIntervalMilli * time.Millisecond)
			d.healthCheck()
		}
	}
}

func (d *dbBalancer) IsAlive() bool {
	return !d.availableDbs.IsEmpty()
}

func (d *dbBalancer) Get() *sql.DB {
	if d.availableDbs.IsEmpty() {
		return nil
	}

	switch d.balanceAlgorithm {
	case RoundRobin:
		return d.availableDbs.Next()
	case Random:
		return d.availableDbs.Random()
	default:
		return nil
	}
}

func (d *dbBalancer) Destroy() {
	d.cancel()
}

func (d *dbBalancer) GetHealthCheckIntervalMilli() int {
	return d.healthCheckIntervalMilli
}

func (d *dbBalancer) SetHealthCheckIntervalMilli(i int) {
	d.healthCheckIntervalMilli = i
}

func (d *dbBalancer) GetBalanceAlgorithm() BalanceAlgorithm {
	return d.balanceAlgorithm
}

func (d *dbBalancer) SetBalanceAlgorithm(balanceAlgorithm BalanceAlgorithm) {
	d.balanceAlgorithm = balanceAlgorithm
}

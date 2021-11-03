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
} = NewDbBalancer(context.Background(), []*sql.DB{})

type BalaceAlgorithm int

const (
	RoundRobin BalaceAlgorithm = iota
	Random
)

type dbBalancer struct {
	ctx                      context.Context
	cancel                   context.CancelFunc
	dbs                      []*sql.DB
	availableDbs             *dbList
	isMulti                  bool
	healthCheckIntervalMilli int
	balaceAlgorithm          BalaceAlgorithm
}

func NewDbBalancer(ctx context.Context, dbs []*sql.DB) *dbBalancer {
	d := &dbBalancer{
		dbs:                      dbs,
		availableDbs:             NewDbList(),
		isMulti:                  len(dbs) > 1,
		healthCheckIntervalMilli: DefaultHealthCheckIntervalMilli,
		balaceAlgorithm:          DefaultBalaceAlgorithm,
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
	availableDbs := make([]*sql.DB, len(d.dbs))
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
	return d.availableDbs.IsEmpty()
}

func (d *dbBalancer) Get() *sql.DB {
	if d.availableDbs.IsEmpty() {
		return nil
	}

	switch d.balaceAlgorithm {
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

func (d *dbBalancer) GetBalaceAlgorithm() BalaceAlgorithm {
	return d.balaceAlgorithm
}

func (d *dbBalancer) SetBalaceAlgorithm(balaceAlgorithm BalaceAlgorithm) {
	d.balaceAlgorithm = balaceAlgorithm
}

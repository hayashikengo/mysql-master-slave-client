package mydb

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TextNewBalancer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db0, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		d := make([]*sql.DB, 0)
		d = append(d, db0)
		dbBalancer := NewDbBalancer(context.Background(), d)
		dbBalancer.Destroy()
		defer dbBalancer.Destroy()
	})
}

func TestGet(t *testing.T) {
	t.Run("success with RoundRobin", func(t *testing.T) {
		db0, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db1, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		d := make([]*sql.DB, 0)
		d = append(d, db0)
		d = append(d, db1)
		dbBalancer := NewDbBalancer(context.Background(), d)
		defer dbBalancer.Destroy()
		dbBalancer.SetBalanceAlgorithm(RoundRobin)

		if dbBalancer.IsAlive() == false {
			t.Errorf("dbBalancer IsAlive() want %t, but get %t", true, dbBalancer.IsAlive())
		}

		if !reflect.DeepEqual(dbBalancer.Get(), db1) {
			t.Error("dbBalancer Get() want db1")
		}

		if !reflect.DeepEqual(dbBalancer.Get(), db0) {
			t.Error("dbBalancer Get() want db0")
		}
	})

	t.Run("success with Random", func(t *testing.T) {
		db0, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db1, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		d := make([]*sql.DB, 0)
		d = append(d, db0)
		d = append(d, db1)
		dbBalancer := NewDbBalancer(context.Background(), d)
		defer dbBalancer.Destroy()
		dbBalancer.SetBalanceAlgorithm(Random)

		if dbBalancer.IsAlive() == false {
			t.Errorf("dbBalancer IsAlive() want %t, but get %t", true, dbBalancer.IsAlive())
		}

		if dbBalancer.Get() == nil {
			t.Error("dbBalancer Get() want db")
		}

		if dbBalancer.Get() == nil {
			t.Error("dbBalancer Get() want db")
		}
	})
}

func TestSetHealthCheckIntervalMilli(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dbBalancer := NewDbBalancer(context.Background(), nil)
		defer dbBalancer.Destroy()

		if dbBalancer.GetHealthCheckIntervalMilli() != 5000 {
			t.Error("GetHealthCheckIntervalMilli() want 5000")
		}

		dbBalancer.SetHealthCheckIntervalMilli(2000)
		if dbBalancer.GetHealthCheckIntervalMilli() != 2000 {
			t.Error("GetHealthCheckIntervalMilli() want 2000")
		}
	})
}

func TestSetBalanceAlgorithm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dbBalancer := NewDbBalancer(context.Background(), nil)
		defer dbBalancer.Destroy()

		if dbBalancer.GetBalanceAlgorithm() != Random {
			t.Error("GetBalanceAlgorithm() want Random")
		}

		dbBalancer.SetBalanceAlgorithm(RoundRobin)
		if dbBalancer.GetBalanceAlgorithm() != RoundRobin {
			t.Error("GetBalanceAlgorithm() want RoundRobin")
		}
	})
}

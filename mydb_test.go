package mydb

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPing(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()

		db := New(master)
		defer db.Close()

		err = db.Ping()
		if err != nil {
			t.Error(err)
		}
		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("success with master and readreplica", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()

		db := New(master, readreplica)
		defer db.Close()

		err = db.Ping()
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("error with master", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		// health check worker
		masterMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master)
		defer db.Close()

		// should raise error
		err = db.Ping()
		if err == nil || err.Error() != "ping error" {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("error with readreplica", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master, readreplica)
		defer db.Close()

		// should raise error
		err = db.Ping()
		if err == nil || err.Error() != "ping error" {
			t.Error(err)
		}
		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestPingContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()

		db := New(master)
		defer db.Close()

		err = db.PingContext(context.Background())
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("success with master and readreplica", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()

		db := New(master, readreplica)
		defer db.Close()

		err = db.PingContext(context.Background())
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("error with master", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		// health check worker
		masterMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master)
		defer db.Close()

		// should raise error
		err = db.PingContext(context.Background())
		if err == nil || err.Error() != "ping error" {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("error with readreplica", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing()
		// mock ping
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master, readreplica)
		defer db.Close()

		// should raise error
		err = db.PingContext(context.Background())
		if err == nil || err.Error() != "ping error" {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestQuery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplicaMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.Query("select 1")
		if err != nil {
			t.Error(err)
		}

		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestQueryContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplicaMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.QueryContext(context.Background(), "select 1")
		if err != nil {
			t.Error(err)
		}

		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestQueryRow(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplicaMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		db := New(master, readreplica)
		defer db.Close()

		var id int
		err = db.QueryRow("select 1").Scan(&id)
		if err != nil {
			t.Error(err)
		}
		if id != 1 {
			t.Error("id want 1")
		}

		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestQueryRowContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplicaMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		db := New(master, readreplica)
		defer db.Close()

		var id int
		err = db.QueryRowContext(context.Background(), "select 1").Scan(&id)
		if err != nil {
			t.Error(err)
		}
		if id != 1 {
			t.Error("id want 1")
		}

		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestBegin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectBegin()

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.Begin()
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestBeginTx(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectBegin()

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.BeginTx(context.Background(), nil)
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestExec(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectExec("select 1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.Exec("select 1")
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestExecContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectExec("select 1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.ExecContext(context.Background(), "select 1")
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestPrepare(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectPrepare("select 1").
			ExpectExec().
			WillReturnResult(sqlmock.NewResult(1, 1))
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master, readreplica)
		defer db.Close()

		ctx := context.Background()
		stmt, err := db.Prepare("select 1")
		if err != nil {
			t.Error(err)
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx)
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestPrepareContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		masterMock.ExpectPrepare("select 1").
			ExpectExec().
			WillReturnResult(sqlmock.NewResult(1, 1))
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master, readreplica)
		defer db.Close()

		ctx := context.Background()
		stmt, err := db.PrepareContext(ctx, "select 1")
		if err != nil {
			t.Error(err)
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx)
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestSetConnMaxLifetime(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master, readreplica)
		defer db.Close()

		db.SetConnMaxLifetime(10)
		// FIXME: assert equal value
	})
}

func TestSetMaxIdleConns(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master, readreplica)
		defer db.Close()

		db.SetMaxIdleConns(10)
		// FIXME: assert equal value
	})
}

func TestSetMaxOpenConns(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master, readreplica)
		defer db.Close()

		if db.master.Stats().MaxOpenConnections != 0 {
			t.Error("Stats().MaxOpenConnections want 0")
		}
		if db.readreplicas[0].Stats().MaxOpenConnections != 0 {
			t.Error("Stats().MaxOpenConnections want 0")
		}

		db.SetMaxOpenConns(10)
		if db.master.Stats().MaxOpenConnections != 10 {
			t.Error("Stats().MaxOpenConnections want 10")
		}
		if db.readreplicas[0].Stats().MaxOpenConnections != 10 {
			t.Error("Stats().MaxOpenConnections want 0")
		}
	})
}

func TestHealthCheckIntervalMilli(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master)
		defer db.Close()

		if db.GetHealthCheckIntervalMilli() != 5000 {
			t.Error("GetHealthCheckIntervalMilli() want 5000")
		}

		db.SetHealthCheckIntervalMilli(2000)
		if db.GetHealthCheckIntervalMilli() != 2000 {
			t.Error("GetHealthCheckIntervalMilli() want 2000")
		}
	})
}

func TestBalaceAlgorithm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master)
		defer db.Close()

		if db.GetBalaceAlgorithm() != Random {
			t.Error("GetBalaceAlgorithm() want Random")
		}

		db.SetBalaceAlgorithm(RoundRobin)
		if db.GetBalaceAlgorithm() != RoundRobin {
			t.Error("GetBalaceAlgorithm() want RoundRobin")
		}
	})
}

func TestFallbackType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		db := New(master)
		defer db.Close()

		if db.GetFallbackType() != UseMaster {
			t.Error("GetFallbackType() want UseMaster")
		}

		db.SetFallbackType(None)
		if db.GetFallbackType() != None {
			t.Error("GetFallbackType() want None")
		}
	})
}

func TestFallbackWithUseMaster(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing().WillReturnError(errors.New("ping error"))

		// mock query
		masterMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		db := New(master, readreplica)
		defer db.Close()

		_, err = db.Query("select 1")
		if err != nil {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("error with die master and all readreplicas", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing().WillReturnError(errors.New("ping error"))
		readreplicaMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master, readreplica)
		defer db.Close()

		// should raise error
		_, err = db.Query("select 1")
		if err == nil || err != ErrMasterDied {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestFallbackWithNone(t *testing.T) {
	t.Run("error with die all readreplicas", func(t *testing.T) {
		master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}
		readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		if err != nil {
			t.Error(err.Error())
		}

		// initial health check
		masterMock.ExpectPing()
		readreplicaMock.ExpectPing().WillReturnError(errors.New("ping error"))

		db := New(master, readreplica)
		db.SetFallbackType(None)
		defer db.Close()

		// should raise error
		_, err = db.Query("select 1")
		if err == nil || err != ErrAllReadreplicaDied {
			t.Error(err)
		}

		if err := masterMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		if err := readreplicaMock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

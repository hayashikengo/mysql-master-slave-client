package mydb

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func BenchmarkSqlSelect(b *testing.B) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		b.Error(err.Error())
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		dbMock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Query("select 1")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMydbSelectWithRandom(b *testing.B) {
	master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica0, readreplica0Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica1, readreplica1Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica2, readreplica2Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica3, readreplica3Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}

	// initial health check
	masterMock.ExpectPing()
	readreplica0Mock.ExpectPing()
	readreplica1Mock.ExpectPing()
	readreplica2Mock.ExpectPing()
	readreplica3Mock.ExpectPing()

	// mock query
	for i := 0; i < b.N; i++ {
		readreplica0Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica1Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica2Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica3Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	}

	db := New(master, readreplica0, readreplica1, readreplica2, readreplica3)
	db.SetBalaceAlgorithm(Random)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = db.Query("select 1")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMydbSelectWithRoundRobin(b *testing.B) {
	master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica0, readreplica0Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica1, readreplica1Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica2, readreplica2Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica3, readreplica3Mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}

	// initial health check
	masterMock.ExpectPing()
	readreplica0Mock.ExpectPing()
	readreplica1Mock.ExpectPing()
	readreplica2Mock.ExpectPing()
	readreplica3Mock.ExpectPing()

	// mock query
	for i := 0; i < b.N; i++ {
		readreplica0Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica1Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica2Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		readreplica3Mock.ExpectQuery("select 1").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	}

	db := New(master, readreplica0, readreplica1, readreplica2, readreplica3)
	db.SetBalaceAlgorithm(RoundRobin)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = db.Query("select 1")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSqlExec(b *testing.B) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		b.Error(err.Error())
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		dbMock.ExpectExec("insert into code values (100), (200)").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err := db.Exec("insert into code values (100), (200)")
		if err != nil {
			b.Error(err)
		}
		_, err = result.RowsAffected()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMydbExec(b *testing.B) {
	master, masterMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}
	readreplica, readreplicaMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		b.Error(err.Error())
	}

	// initial health check
	masterMock.ExpectPing()
	readreplicaMock.ExpectPing()

	// mock query
	for i := 0; i < b.N; i++ {
		masterMock.ExpectExec("insert into code values (100), (200)").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	db := New(master, readreplica)
	db.SetBalaceAlgorithm(Random)
	defer db.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, err := db.Exec("insert into code values (100), (200)")
		if err != nil {
			b.Error(err)
		}
		_, err = result.RowsAffected()
		if err != nil {
			b.Error(err)
		}
	}
}

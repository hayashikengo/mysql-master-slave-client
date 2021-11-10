package main

import (
	"database/sql"
	"fmt"
	"mydb"
	"testing"
	"time"
)

func BenchmarkSqlSelect(b *testing.B) {
	db, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, err := db.Query("select * from code")
		if err != nil {
			b.Error(err)
		}
		for rows.Next() {
			c := Code{}
			rows.Scan(&c.Code)
		}
	}
}

func BenchmarkMydbSelectRandom(b *testing.B) {
	master, err := sql.Open("mysql", "mydb_user:mydb_pwd@tcp(127.0.0.1:4406)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave1, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:6606)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave2, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:7706)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db := mydb.New(master, slave, slave1, slave2)
	db.SetBalaceAlgorithm(mydb.Random)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query("select * from code")
		if err != nil {
			b.Error(err)
		}
		for rows.Next() {
			c := Code{}
			rows.Scan(&c.Code)
		}
	}
}

func BenchmarkMydbSelectRoundrobin(b *testing.B) {
	master, err := sql.Open("mysql", "mydb_user:mydb_pwd@tcp(127.0.0.1:4406)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave1, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:6606)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave2, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:7706)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db := mydb.New(master, slave, slave1, slave2)
	db.SetBalaceAlgorithm(mydb.RoundRobin)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		rows, err := db.Query("select * from code")
		if err != nil {
			b.Error(err)
		}
		for rows.Next() {
			c := Code{}
			rows.Scan(&c.Code)
		}
	}
}

func BenchmarkSqlExec(b *testing.B) {
	db, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
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

func BenchmarkMydbExec(b *testing.B) {
	master, err := sql.Open("mysql", "mydb_user:mydb_pwd@tcp(127.0.0.1:4406)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave1, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:6606)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave2, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:7706)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db := mydb.New(master, slave, slave1, slave2)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
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

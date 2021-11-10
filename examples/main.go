package main

import (
	"database/sql"
	"fmt"
	"mydb"

	_ "github.com/go-sql-driver/mysql"
)

type Code struct {
	Code int
}

func main() {
	master, err := sql.Open("mysql", "mydb_user:mydb_pwd@tcp(127.0.0.1:4406)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave1, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	slave2, err := sql.Open("mysql", "mydb_slave_user:mydb_slave_pwd@tcp(127.0.0.1:5506)/mydb?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	// new mydb instance
	db := mydb.New(master, slave1, slave2)

	// close db connections
	defer db.Close()

	// setting
	db.SetFallbackType(mydb.UseMaster)
	db.SetBalanceAlgorithm(mydb.Random)
	db.SetHealthCheckIntervalMilli(1000)
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	// exec by master
	result, err := db.Exec("insert into code values (100), (200)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	// query by readreplica
	rows, err := db.Query("select * from code")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		c := Code{}
		rows.Scan(&c.Code)
		fmt.Println(c)
	}
}

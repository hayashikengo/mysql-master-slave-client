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
	db := mydb.New(master, slave1, slave2)

	rows, err := db.Query("select * from code")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		c := Code{}
		rows.Scan(&c.Code)
		fmt.Println(c)
	}

	db.Close()
}

# Mydb
Mydb is client library to abstracts access to master-slave physical Sql server.

Inspired by [nap (MIT License)](https://github.com/tsenart/nap) and [mssqlx (MIT License)](https://github.com/linxGnu/mssqlx)

## Overview
- manage mysql master slave access
- slave traffic balancing
- health check master and slave
- thread safe

## Getting Started
### Install
<!-- TODO: add install command -->
```shell
$ go get ...
```

### Usage
<!-- TODO: add usage -->
```go
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
```
### Benchmark
<!-- TODO: add Benchmark -->
```
```

## Dev tool
### Test
```bash
$ make test-all
```
### Build
```bash
$ make build
```

### Run master slave db

```bash
# initialize db setting on the first time
$ make docker-master-slave-db-initialize

# run master slave db
$ make docker-master-slave-db-run
```

### Run example code

```bash
$ cd ./examples
$ go run main.go
{100}
{200}
```

## License
© hkengo, 2021~time.Now
<!-- TODO: add License -->

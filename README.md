# Mydb
Mydb is client library to abstracts access to master-slave physical Sql server.

Inspired by [nap (MIT License)](https://github.com/tsenart/nap) and [mssqlx (MIT License)](https://github.com/linxGnu/mssqlx)

[![test status](https://github.com/m-rec/6685f2732193a4a5a150add9369734dbf0f260c9/workflows/test/badge.svg?branch=master "test status")](https://github.com/m-rec/6685f2732193a4a5a150add9369734dbf0f260c9/actions)

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

```go
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

	// configuration
	db.SetFallbackType(mydb.UseMaster)
	db.SetBalanceAlgorithm(mydb.Random)

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
```

### Configuration

#### Readreplica Balancing Algorithm configuration
```go
db.SetBalanceAlgorithm(mydb.Random) // default Random
```
- Random
- RoundRobin

#### Fallback type configuration
```go
db.SetFallbackType(mydb.UseMaster) // default UseMaster
```
- None
  - Return `ErrAllReadreplicaDied` if all readreplica died.
- UseMaster
  - Use Master if all readreplica died.
  - Return `ErrMasterDied` if all readreplica and master died.

#### Health check interval configuration
```go
db.SetHealthCheckIntervalMilli(1000) // default 5000
```

#### DB connection configuration
```go
db.SetConnMaxLifetime(10)
db.SetMaxIdleConns(10)
db.SetMaxOpenConns(10)
```

## Benchmark
```
$ make bench-with-docker
cd ./examples && \
	go test -bench=.
goos: darwin
goarch: amd64
pkg: mydb_examples
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkSqlSelect-8              	     441	   2741500 ns/op
BenchmarkMydbSelectRandom-8       	     634	   1612263 ns/op
BenchmarkMydbSelectRoundrobin-8   	     800	   1600085 ns/op
BenchmarkSqlExec-8                	     393	   2876589 ns/op
BenchmarkMydbExec-8               	     344	   3090725 ns/op
PASS
ok  	mydb_examples	8.230s
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

### Run master and 3 readreplica dbs by docker-compose
```bash
# initialize db setting on the first time
$ make docker-master-slave-db-initialize

# run master slave db
$ make docker-master-slave-db-run
```
### Benchmark (mydb vs database/sql)
```bash
# run with sqlmock
$ make bench-with-sqlmock

# run with docker
$ make bench-with-docker
```

## License
© hkengo, 2021~time.Now
<!-- TODO: add License -->

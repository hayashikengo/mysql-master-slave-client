package mydb

import (
	"database/sql"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var _ interface {
	IsEmpty() bool
	Current() *sql.DB
	Next() *sql.DB
	Random() *sql.DB
	Replace(dbs []*sql.DB)
} = NewDbList()

type dbList struct {
	lk           sync.RWMutex
	list         []*sql.DB
	currentIndex int
}

func NewDbList() *dbList {
	return &dbList{
		currentIndex: 0,
	}
}

func (d *dbList) IsEmpty() bool {
	return len(d.list) == 0
}

func (d *dbList) Current() (res *sql.DB) {
	if d.IsEmpty() {
		return nil
	}

	d.lk.RLock()
	defer d.lk.RUnlock()

	res = d.list[d.currentIndex]

	return
}

func (d *dbList) Next() (res *sql.DB) {
	d.lk.RLock()
	defer d.lk.RUnlock()

	len := len(d.list)
	if len > 0 {
		d.currentIndex = (d.currentIndex + 1) % len
		res = d.list[d.currentIndex]
	}

	return
}

func (d *dbList) Random() (res *sql.DB) {
	d.lk.RLock()
	defer d.lk.RUnlock()

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(d.list))
	res = d.list[num]

	return
}

func (d *dbList) isSame(dbs []*sql.DB) bool {
	if len(d.list) != len(dbs) {
		return false
	}

	// OPTIMIZE: reflect.DeepEqual is slow?? want get bench.
	return reflect.DeepEqual(d.list, dbs)
}

func (d *dbList) Replace(dbs []*sql.DB) {
	// If same dbs, not replace
	if d.isSame(dbs) {
		return
	}

	d.lk.Lock()
	defer d.lk.Unlock()

	d.list = dbs
}

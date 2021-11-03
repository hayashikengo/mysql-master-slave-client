package mydb

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewDbList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dbs := NewDbList()
		if dbs.currentIndex != 0 {
			t.Errorf("NewDbList() want currentIndex is %d but get %d", 0, dbs.currentIndex)
		}
		if dbs.IsEmpty() == false {
			t.Errorf("NewDbList() want list is Empty but not Empty")
		}
	})
}

func TestIsEmpty(t *testing.T) {

	t.Run("success true", func(t *testing.T) {
		dbs := NewDbList()
		if dbs.IsEmpty() != true {
			t.Errorf("dbList.IsEmpty() want %t, bad got %t", true, false)
		}
	})

	t.Run("success false", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		d := make([]*sql.DB, 1)
		d = append(d, db)
		dbs := dbList{
			list: d,
		}

		if dbs.IsEmpty() != false {
			t.Errorf("dbList.IsEmpty")
		}
	})
}

func TestCurrent(t *testing.T) {

	t.Run("success empty", func(t *testing.T) {
		dbs := NewDbList()
		if dbs.Current() != nil {
			t.Errorf("Current want nil, but not nil")
		}
	})

	t.Run("success current", func(t *testing.T) {
		db, _, err := sqlmock.New()
		if err != nil {
			t.Error(err.Error())
		}
		d := make([]*sql.DB, 0)
		d = append(d, db)
		dbs := dbList{
			list: d,
		}
		if !reflect.DeepEqual(dbs.Current(), db) {
			t.Errorf("Current is want db, but not exist")
		}
	})
}

func TestNext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
		dbs := dbList{
			list: d,
		}
		if !reflect.DeepEqual(dbs.Next(), db1) {
			t.Errorf("Next() is want db1, but not db1")
		}
		if !reflect.DeepEqual(dbs.Next(), db0) {
			t.Errorf("Next() is want db0, but not db0")
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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
		dbs := dbList{
			list: d,
		}
		result := dbs.Random()
		if !reflect.DeepEqual(result, db0) && !reflect.DeepEqual(result, db1) {
			t.Errorf("Random() is want db0 or db1, but not db0 or db1")
		}
	})
}

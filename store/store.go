package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Store interface {
}

type datastore struct {
	Db *xorm.Engine
}

func New(driver, url string) (Store, error) {
	db, err := xorm.NewEngine(driver, url)
	if err != nil {
		return datastore{}, err
	}
	return datastore{Db: db}, err

}

package store

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type Store interface {
	Close()
}

type datastore struct {
	Db *gorm.DB
}

func New(driver, url string) (Store, error) {
	db, err := gorm.Open(driver, url)
	if err != nil {
		return datastore{}, err
	}
	return datastore{Db: db}, err
}

func (ds datastore) Close() {
	ds.Db.Close()
}

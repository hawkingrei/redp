package datastore

import (
	"reflect"

	"github.com/hawkingrei/redp/store"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type datastore struct {
	Db *gorm.DB
}

func New(driver, url string) (store.Store, error) {
	db, err := gorm.Open(driver, url)
	if err != nil {
		return datastore{}, err
	}
	return datastore{Db: db}, err
}

func (ds datastore) Close() {
	ds.Db.Close()
}

func (ds datastore) CreateTable(models ...interface{}) {
	for _, model := range models {
		if !ds.Db.HasTable(model) {
			logrus.Info("create table ", reflect.TypeOf(model).Elem().Name())
			ds.Db.CreateTable(model)
		}
	}
}

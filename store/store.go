package store

import (
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type Store interface {
	CreateTable(models ...interface{})
	Close()
}

type datastore struct {
	Db *gorm.DB
}

func New(driver, url string) (Store, error) {
	db, err := gorm.Open(driver, url)
	//if !db.HasTable(&model.User{}) {
	//	db.CreateTable(&model.User{})
	//}
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

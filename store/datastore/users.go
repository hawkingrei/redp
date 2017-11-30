package datastore

import (
	"github.com/hawkingrei/redp/model"
	"github.com/sirupsen/logrus"
)

func (ds datastore) GetUser(username string) (*model.User, error) {
	var usr = new(model.User)
	err := ds.Db.Where("username = ?", username).Find(&usr).Error
	return usr, err
}

func (ds datastore) CreateUser(username string) {
	user := model.User{Username: username, Memory: 0}
	logrus.Debug("ds create user ", user)
	ds.Db.Create(&user)
}

func (ds datastore) HasUser(username string) bool {
	_, err := ds.GetUser(username)
	if err != nil {
		return false
	}
	return true
}

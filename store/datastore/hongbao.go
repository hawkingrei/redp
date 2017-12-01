package datastore

import (
	"errors"
	"math/rand"
	"time"

	"github.com/hawkingrei/redp/internal/hongbao"
	"github.com/hawkingrei/redp/model"
	"github.com/sirupsen/logrus"
)

func randStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ds datastore) CreateSendedHongbao(username string, money float32, num int) (*model.SendedHongbao, error) {
	if money/float32(num) < 0.01 {
		return &model.SendedHongbao{}, errors.New("money is too small,but share num is too large")
	}
	hb := model.SendedHongbao{Username: username, Num: num, Money: money, CreateTime: time.Now(), Password: randStringRunes(8)}
	err := ds.Db.Create(&hb).Error
	hongbaos := hongbao.GenerateMoneyVector(money, num)
	logrus.Debug("ds create SendedHongbao ", hb)
	for _, v := range hongbaos {
		gothb := model.GotHongbao{Hbid: hb.Hbid, Money: v}
		ds.Db.Create(&gothb)
		logrus.Debug("ds create gotHongbao", gothb)
	}
	return &hb, err
}

func (ds datastore) GrabHongbao(hid int64, username string, password string) (*model.GotHongbao, error) {
	var shd model.SendedHongbao
	var ghd model.GotHongbao
	tx := ds.Db.Begin()
	err := tx.Where("hbid = ? AND closed = 0", hid).Find(&shd).Error
	if err != nil {
		tx.Commit()
		logrus.Debug("ds GrabHongbao query sened hongbao ", err.Error())
		return &ghd, err
	}
	if password == shd.Password {
		tx.Commit()
		return &ghd, errors.New("Password ERROR")
	}
	err = tx.Where("hbid = ? And username = ?", hid, username).Find(&ghd).Error
	if err == nil {
		tx.Commit()
		logrus.Debug("ds GrabHongbao username has got hongbao ")
		return &ghd, err
	}
	err = tx.Where("hbid = ? And username = \"\"", hid).First(&ghd).Error
	if err != nil {
		tx.Commit()
		logrus.Debug("ds GrabHongbao get hongbao ", err.Error())
		return &ghd, err
	}
	ghd.Username = username
	err = tx.Save(&ghd).Error
	if err != nil {
		tx.Rollback()
		logrus.Debug("ds GrabHongbao save got hongbao ", err.Error())
		return &ghd, err
	}
	var myaccount model.User
	err = tx.Where("username = ?", username).Find(&myaccount).Error
	if err != nil {
		tx.Rollback()
		logrus.Debug("ds GrabHongbao get account ", err.Error())
		return &ghd, err
	}
	myaccount.Memory = myaccount.Memory + ghd.Money
	err = tx.Save(&myaccount).Error
	if err != nil {
		tx.Rollback()
		logrus.Debug("ds GrabHongbao save account ", err.Error())
		return &ghd, err
	}
	tx.Commit()
	return &ghd, err
}

func (ds datastore) ListGotHongbao(username string) ([]model.GotHongbao, error) {
	var gotHongbaos []model.GotHongbao
	err := ds.Db.Where(" username = ?", username).Find(&gotHongbaos).Error
	if err != nil {
		return gotHongbaos, err
	}
	return gotHongbaos, err
}

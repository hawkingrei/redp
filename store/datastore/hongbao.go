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



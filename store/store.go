package store

import (
	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hawkingrei/redp/model"
)

type Store interface {
	// GetUser gets a user by unique Username.
	GetUser(string) (*model.User, error)
	CreateUser(username string)
	HasUser(string) bool

	CreateSendedHongbao(username string, money float32, num int) (*model.SendedHongbao, error)
	GrabHongbao(hid int64, username string, password string) (*model.GotHongbao, error)
	ListGotHongbao(username string) ([]model.GotHongbao, error)
	//CreateGotHongbao(username string, hid int64) (*model.GotHongbao, error)
	CreateTable(models ...interface{})
	Close()
}

func GetUser(c context.Context, username string) (*model.User, error) {
	return FromContext(c).GetUser(username)
}

func CreateUser(c context.Context, username string) {
	FromContext(c).CreateUser(username)
}

func HasUser(c context.Context, username string) bool {
	return FromContext(c).HasUser(username)
}

func CreateSendedHongbao(c context.Context, username string, money float32, num int) (*model.SendedHongbao, error) {
	return FromContext(c).CreateSendedHongbao(username, money, num)
}

func GrabHongbao(c context.Context, hid int64, username string, password string) (*model.GotHongbao, error) {
	return FromContext(c).GrabHongbao(hid, username, password)
}

func ListGotHongbao(c context.Context, username string) ([]model.GotHongbao, error) {
	return FromContext(c).ListGotHongbao(username)
}

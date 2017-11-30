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

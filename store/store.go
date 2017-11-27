package store

import "github.com/go-xorm/xorm"

func New(driver, url string) (*xorm.Engine, error) {
	return xorm.NewEngine(driver, url)
}

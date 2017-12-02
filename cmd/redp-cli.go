package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/hawkingrei/redp/conf"
	"github.com/hawkingrei/redp/internal/version"
	"github.com/hawkingrei/redp/model"
	"github.com/hawkingrei/redp/routers"
	"github.com/hawkingrei/redp/routers/middleware"
	"github.com/hawkingrei/redp/routers/middleware/header"
	"github.com/hawkingrei/redp/store"
	"github.com/hawkingrei/redp/store/datastore"
	"github.com/sirupsen/logrus"
)

func setupStore(c *conf.Configure) (store.Store, error) {
	store, err := datastore.New(
		c.DbDriver,
		c.DbURL,
	)
	return store, err
}

func redpFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("redp", flag.ExitOnError)
	flagSet.Bool("version", false, "print version string")
	flagSet.String("config", "", "path to config file")
	return flagSet
}

func loadmeta(configFile string) (meta *conf.Configure, err error) {
	if configFile != "" {
		_, err = toml.DecodeFile(configFile, &meta)
		if err != nil {
			return
		}
	}
	return
}

func CreateStote(config *conf.Configure) store.Store {
	store_, err := setupStore(config)
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
	store_.CreateTable(&model.User{})
	store_.CreateTable(&model.SendedHongbao{})
	store_.CreateTable(&model.GotHongbao{})
	return store_
}

func CreateHttpHandler(store store.Store, config *conf.Configure) http.Handler {
	handler := routers.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Version,
		header.NoCache,
		middleware.Store(config, store),
	)
	return handler
}

func main() {
	flagSet := redpFlagSet()
	flagSet.Parse(os.Args[1:])
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) || len(os.Args) == 1 {
		fmt.Println(version.String())
		os.Exit(0)
	}
	configFile := flagSet.Lookup("config").Value.String()
	config, err := loadmeta(configFile)
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
	}
	if config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	quit := make(chan os.Signal)
	store_ := CreateStote(config)
	handler := CreateHttpHandler(store_, config)
	signal.Notify(quit, os.Interrupt)
	serve := &http.Server{
		Addr:    ":9000",
		Handler: handler,
	}
	var g errgroup.Group
	g.Go(func() error {
		return serve.ListenAndServe()
	})
	g.Go(func() error {
		ticker := time.NewTicker(time.Duration(1) * time.Second)
		for {
			select {
			case <-quit:
				logrus.Info("receive interrupt signal")
				store_.Close()
				return serve.Close()

			case <-ticker.C:
				store_.Background(config.HBtimeout)
			}
		}
	})
	g.Wait()
}

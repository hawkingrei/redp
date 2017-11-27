package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/hawkingrei/redp/conf"
	"github.com/hawkingrei/redp/internal/version"
)

func redpFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("redp", flag.ExitOnError)
	flagSet.Bool("version", false, "print version string")
	flagSet.String("config", "", "path to config file")
	return flagSet
}

func loadmeta(configFile string) (meta conf.Configure, err error) {
	if configFile != "" {
		_, err = toml.DecodeFile(configFile, &meta)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	flagSet := redpFlagSet()
	flagSet.Parse(os.Args[1:])

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String())
		os.Exit(0)
	}

}

package main

import "github.com/BurntSushi/toml"

type (
	config struct {
		Starling starling
		Monzo    Monzo
	}

	starling struct {
		StarlingToken string
	}

	Monzo struct {
		CurrentAccount CurrentAccount
	}

	CurrentAccount struct {
		CurrentAccount string
	}
)

var (
	conf config
)

func initConfig() {
	if _, err := toml.Decode(loadConfig(), &conf); err != nil {
		Error.Fatalln("Could not load configuration", err)
	}
}

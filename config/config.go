package config

import "github.com/tkanos/gonfig"


type Configuration struct {
	PathImages string
}

func Config() (Configuration, error) {
	configuration := Configuration{}
	err := gonfig.GetConf("config/dcontainer.json", &configuration)
	return configuration,err
}

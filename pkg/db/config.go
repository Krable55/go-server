package db

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	API_KEY           string
	DATABASE          string
	HOST_NAME         string
	PORT              int
	USER_NAME         string
	PASSWORD          string
	CONNECTION_STRING string
}

func GetConfig(env string) Configuration {
	configuration := Configuration{}

	fileName := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

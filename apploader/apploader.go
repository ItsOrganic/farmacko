package apploader

import (
	"flag"
	"log"
	"os"

	"github.com/itsorganic/farmacko-assignment/cache"
	"github.com/itsorganic/farmacko-assignment/constants"
	"github.com/itsorganic/farmacko-assignment/globals"
	"github.com/itsorganic/farmacko-assignment/models"
	"gopkg.in/yaml.v3"
)

func Init() {
	cache.Init()
	configPath := getConfigPath()
	LoadConfig(configPath)
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "config.yaml", "path to config file")
	flag.Parse()
	if path == "" {
		log.Fatal(constants.ERR_CONFIG_PATH_NOT_FOUND)
	}
	return path
}

func LoadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config := &models.Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	globals.Config = config
}

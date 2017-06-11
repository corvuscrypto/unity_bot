package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	BotName string `yaml:"botName"`
	UIPort  int    `yaml:"uiPort"`
}

// GlobalConfig is the golang representation of the configuration of this bot.
var GlobalConfig *config

// LoadConfig opens a file at the given filepath and parses it to load into the global config structure
// If no filepath is given then defaults are used.
func LoadConfig(filepath string) {
	if filepath == "" {
		return
	}

	failAndExit := func(err error) {
		log.Fatalf("Unable to load configuration from %s: %s\n", filepath, err.Error())
	}

	file, err := os.Open(filepath)
	if err != nil {
		failAndExit(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		failAndExit(err)
	}

	err = yaml.Unmarshal(fileBytes, GlobalConfig)
	if err != nil {
		failAndExit(err)
	}
}

func init() {
	GlobalConfig = new(config)
	setDefaults() //Set the defaults
}

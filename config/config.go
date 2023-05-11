package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func WriteConfig(config map[string]interface{}) error {
	configToWrite, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalln("Config marshalling failed!")
	}
	home, _ := os.UserHomeDir()
	log.Println("Updating config file...")
	os.RemoveAll(home + "/.telegramcloner.yaml")
	writeConfigFileError := os.WriteFile(home+"/.telegramcloner.yaml", configToWrite, 0644)
	if writeConfigFileError != nil {
		log.Printf("Error while saving configuration file %s", writeConfigFileError)
	}
	return writeConfigFileError
}

func CheckConfigFileExistence() {
	home, _ := os.UserHomeDir()
	_, configFileNotFoundError := os.Stat(home + "/.telegramcloner.yaml")
	if configFileNotFoundError != nil {
		log.Fatalln("Config file does not exists!")
	}
}

package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Conf -> config.yaml structure
type Conf struct {
	AppleCountries string `yaml:"appleCountries"`
	AppleSubregion string `yaml:"appleSubregion"`
	Test           string `yaml:"test"`
}

//ReadConf -> Function to read configs from yaml
func (c *Conf) ReadConf() *Conf {
	yamlFile, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println("Configs loaded successfully")
	return c
}

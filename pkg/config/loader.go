package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Load(configFile string) (*ServerConfig, error) {

	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}

	// NOTE: ExpandEnv will substitute any variable defined as
	// $VAR_NAME (example) in the config file with corresponding env variable
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	c := ServerConfig{}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

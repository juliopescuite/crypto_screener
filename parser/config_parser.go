package parser

import (
	"github.com/juliopescuite/crypto_screener/data_types"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseScreenerConfig() data_types.ScreeningConfig {
	var config data_types.ScreeningConfig
	yamlFile, err := os.ReadFile("config/screener_config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

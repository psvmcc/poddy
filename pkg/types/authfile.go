package types

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AuthFile []struct {
	User  string `yaml:"user"`
	Role  string `yaml:"role"`
	Token string `yaml:"token"`
}

func (c *AuthFile) Load(cfgFile string) {
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Printf("Read config file error: %v", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal config file error: %v", err)
		os.Exit(1)
	}
}

package types

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AuthFile []struct {
	User  string `yaml:"user"`
	Role  string `yaml:"role"`
	Token string `yaml:"token"`
}

func (c *AuthFile) Load(cfgFile string) {
	yamlFile, err := os.ReadFile(filepath.Clean(cfgFile))
	if err != nil {
		log.Fatalf("Read config file error: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal config file error: %v", err)
	}
}

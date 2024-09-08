package types

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	AuthFile string `yaml:"auth_file"`
	DataPath string `yaml:"data_path"`
	Roles    map[string]struct {
		NamespacesAccess map[string]string `yaml:"namespace_access"`
	} `yaml:"roles"`
	Namespaces map[string]struct {
		Network string            `yaml:"network"`
		ENV     map[string]string `yaml:"env"`
	} `yaml:"namespaces"`
}

func (c *ConfigFile) Load(cfgFile string) {
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("Read config file error: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal config file error: %v", err)
	}
}

package internal

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Endpoint string `yaml:"connection_endpoint"`
	Token    string `yaml:"jwt_token"`
}

func NewConfig(configPath string) (*AppConfig, error) {
	config := &AppConfig{}

	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

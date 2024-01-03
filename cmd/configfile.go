package cmd

import (
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Instances []Instance `yaml:"instances"`
}

type Instance struct {
	URL      string `yaml:"url"`
	Instance IP     `yaml:"instance"`
}

type IP struct {
	IP string `yaml:"ip"`
}

// generateConfigFile generates a configuration file and saves it to the user's home directory.
func generateConfigFile() {
	config := Config{
		Instances: []Instance{
			{
				URL: "http://192.168.1.35:8001",
				Instance: IP{
					IP: "192.168.1.171",
				},
			},
			{
				URL: "http://192.168.1.35:8002",
				Instance: IP{
					IP: "192.168.1.172",
				},
			},
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	// Get the user's home directory
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	configDir := filepath.Join(usr.HomeDir, ".config", "coffeeburn")
	configFile := filepath.Join(configDir, "config.yaml")

	// Create config directory if it doesn't exist
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Write the file
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

package config

import (
	"fmt"
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

func GenerateConfigFile(cfgFile string) {

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

	configFile := ""

	if cfgFile != "" {
		configFile = filepath.Join(cfgFile)

	} else {
		if _, err := os.Stat(filepath.Join(usr.HomeDir, ".config", "burntcoffee", "config.yaml")); err == nil {
			fmt.Print("Config file already exists\n")
			return
		}
		configDir := filepath.Join(usr.HomeDir, ".config", "burntcoffee")
		configFile = filepath.Join(configDir, "config.yaml")

		err = os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		panic(err)
	}
}

func GetConfig(cfgFile string) Config {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	if cfgFile == "" {
		cfgFile = filepath.Join(usr.HomeDir, ".config", "burntcoffee", "config.yaml")
	}

	fmt.Println("Using config file:", cfgFile)
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		fmt.Printf("Config file not found: %s \n You can generate a config-file using ./burntcoffee gen-config \n", err)
		os.Exit(1)
	}

	// Parse config file
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

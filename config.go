package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	defaultConfigDir  = ".config/csa"
	defaultConfigFile = "archive.yml"
)

type Target struct {
	Path    string   `yaml:"path"`
	Exclude []string `yaml:"exclude,omitempty"`
}

type Remote struct {
	Type            string `yaml:"type"`
	Location        string `yaml:"location"`
	CredentialsFile string `yaml:"credentialsFile"`
	NumArchives     int    `yaml:"numArchives"`
}

type Config struct {
	Targets []Target `yaml:"targets"`
	Remotes []Remote `yaml:"remotes"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, defaultConfigDir, defaultConfigFile), nil
}

func getConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

func saveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func initConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("config file already exists at %s", configPath)
	}

	defaultConfig := &Config{
		Targets: []Target{},
		Remotes: []Remote{},
	}

	return saveConfig(defaultConfig)
}

func addTarget(path string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	for _, target := range config.Targets {
		if target.Path == absPath {
			return fmt.Errorf("target %s already exists", absPath)
		}
	}

	config.Targets = append(config.Targets, Target{Path: absPath})
	return saveConfig(config)
}

package config

import (
	"os"
	"path/filepath"
)

type UserConfigValue interface {
	Value() string
	Message() string
	Validate() error
}

type UserConfig struct {
	GithubUsername GithubUsername `yaml:"github_username"`
	DefaultURL     string         `yaml:"default_url"`
	AccessToken    string
}

// GetConfigPath returns the appropriate config file path based on OS
func GetConfigPath() (string, error) {
	appName := "GPM"
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appName, "config.yaml"), nil
}

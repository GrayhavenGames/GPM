package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// GetConfigPath returns the appropriate config file path based on OS
func GetConfigPath() (string, error) {
	appName := "GPM"
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appName, "config.yaml"), nil
}

type UserConfig struct {
	GithubUsername string `yaml:"github_username"`
	DefaultURL     string `yaml:"default_url"`
}

// SaveConfig saves UserConfig to file
func SaveConfig(path string, cfg UserConfig) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

// LoadConfig reads UserConfig from file
func LoadConfig(path string) (UserConfig, error) {
	var cfg UserConfig
	data, err := os.ReadFile(path)
	if err != nil {
		// If file does not exist, return an empty config
		if os.IsNotExist(err) {
			return UserConfig{}, nil
		}
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}

// UpdateGithubUsername updates the GithubUsername field in the config
func UpdateGithubUsername(username string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Load existing config
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Update the username
	cfg.GithubUsername = username

	// Save updated config
	return SaveConfig(configPath, cfg)
}

// UpdateDefaultURL updates the DefaultURL field in the config
func UpdateDefaultURL(url string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Load existing config
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Update the default URL
	cfg.DefaultURL = url

	// Save updated config
	return SaveConfig(configPath, cfg)
}

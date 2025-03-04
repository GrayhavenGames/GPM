package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type UserConfigValue interface {
	Value() string
	Message() string
	Update(newValue string) error
	Validate() bool
}

type UserConfig struct {
	GithubUsername       `yaml:"github_username"`
	DefaultRepositoryURL `yaml:"default_repository_url"`
	AccessToken          string
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

// AddConfigValue dynamically adds or updates a key-value pair in the YAML config
func AddConfigValue(key string, value string) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Load existing config (or create a new map)
	configData, err := LoadRawConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Update or update the key-value pair
	configData[key] = value

	// Save updated config
	return SaveRawConfig(configPath, configData)
}

// LoadRawConfig loads the configuration file into a map[string]interface{}
func LoadRawConfig(path string) (map[string]interface{}, error) {
	configData := make(map[string]interface{})

	// Check if file exists
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return configData, nil // Return empty map if file doesn't exist
		}
		return nil, err
	}

	// Unmarshal YAML into map
	err = yaml.Unmarshal(data, &configData)
	if err != nil {
		return nil, err
	}

	return configData, nil
}

// SaveRawConfig saves the map[string]interface{} to a YAML file
func SaveRawConfig(path string, configData map[string]interface{}) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(configData)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

// ValueExists checks if a key exists in the configuration file
func ValueExists(key string) (bool, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return false, fmt.Errorf("failed to get config path: %w", err)
	}

	// Load existing config
	configData, err := LoadRawConfig(configPath)
	if err != nil {
		return false, fmt.Errorf("failed to load config: %w", err)
	}

	// Check if key exists
	_, exists := configData[key]
	return exists, nil
}

// GetConfigValue retrieves a specific value from the configuration file
func GetConfigValue(key string) (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", fmt.Errorf("failed to get config path: %w", err)
	}

	// Load existing config
	configData, err := LoadRawConfig(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	// Check if key exists
	value, exists := configData[key]
	if !exists {
		return "", fmt.Errorf("config key '%s' not found", key)
	}

	// Convert value to string if necessary
	valueStr, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("config key '%s' is not a string", key)
	}

	return valueStr, nil
}

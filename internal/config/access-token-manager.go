package config

import (
	"fmt"
	"github.com/zalando/go-keyring"
)

const serviceName = "GPM" // The application name used for keyring storage

// AccessTokenManager handles secure storage and retrieval of access tokens
type AccessTokenManager struct{}

// SaveToken securely stores an access token for a given username
func (m *AccessTokenManager) SaveToken(username, token string) error {
	err := keyring.Set(serviceName, username, token)
	if err != nil {
		return fmt.Errorf("failed to store access token: %w", err)
	}
	return nil
}

// GetToken retrieves the stored access token for a given username
func (m *AccessTokenManager) GetToken(username string) (string, error) {
	token, err := keyring.Get(serviceName, username)
	if err == keyring.ErrNotFound {
		return "", fmt.Errorf("no access token found for user %s", username)
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve access token: %w", err)
	}
	return token, nil
}

// DeleteToken removes the stored access token for a given username
func (m *AccessTokenManager) DeleteToken(username string) error {
	err := keyring.Delete(serviceName, username)
	if err == keyring.ErrNotFound {
		return fmt.Errorf("no access token found to delete for user %s", username)
	} else if err != nil {
		return fmt.Errorf("failed to delete access token: %w", err)
	}
	return nil
}

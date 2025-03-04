package config

import "fmt"

type DefaultRepositoryURL struct {
	value        string
	defaultValue string
	message      string
}

func NewDefaultRepositoryURL() *DefaultRepositoryURL {
	value, _ := GetConfigValue("default_repository_url")
	return &DefaultRepositoryURL{
		value:   value,
		message: "Invalid / No Default Repository URL found in config.  You can set this value in the config with\n\n gpm config update --default-repo <Your Default Repo>",
	}
}

func (d DefaultRepositoryURL) Value() string {
	return d.value
}

func (d DefaultRepositoryURL) Message() string {
	return d.message
}

func (d DefaultRepositoryURL) Update(newValue string) error {
	err := AddConfigValue("default_repository_url", newValue)
	if err != nil {
		fmt.Println("⚠ Error:", err)
	} else {
		fmt.Println("✔ Default Repository URL updated successfully!")
	}
	return nil
}

func (d DefaultRepositoryURL) Validate() bool {
	exists, err := ValueExists("default_repository_url")
	if err != nil {
		fmt.Println("⚠ Error:", err)
	}

	if !exists {
		fmt.Println(d.message)
	}
	return exists
}

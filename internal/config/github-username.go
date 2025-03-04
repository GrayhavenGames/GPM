package config

import "fmt"

type GithubUsername struct {
	value   string
	message string
}

func NewGithubUsername() *GithubUsername {
	value, _ := GetConfigValue("github_username")
	return &GithubUsername{
		value:   value,
		message: "Invalid / No Github Username found in config.  You can set this value in the config with\n\n gpm config update --github-username <Your Username>",
	}
}

func (g GithubUsername) Value() string {
	return g.value
}

func (g GithubUsername) Message() string {
	return g.message
}

func (g GithubUsername) Update(newValue string) error {
	err := AddConfigValue("github_username", newValue)
	if err != nil {
		fmt.Println("⚠ Error:", err)
	} else {
		fmt.Println("✔ GitHub username updated successfully!")
	}
	return nil
}

func (g GithubUsername) Validate() bool {
	exists, err := ValueExists("github_username")
	if err != nil {
		fmt.Println("⚠ Error:", err)
	}
	if !exists {
		fmt.Println(g.message)
	}
	return exists
}

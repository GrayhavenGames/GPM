package config

import "fmt"

type GithubUsername struct {
	value   string
	message string
}

func (g GithubUsername) Validate() error {
	//TODO implement me
	panic("implement me")
}

func NewGithubUsername(username string) *GithubUsername {
	message := fmt.Sprintf("Invalid / No Github Username found in config.  You can set this value in the config with\n\n gpm config update --github-username <Your Username>")
	return &GithubUsername{
		value:   username,
		message: message,
	}
}

func (g GithubUsername) Value() string {
	return g.value
}

func (g GithubUsername) Message() string {
	return g.message
}

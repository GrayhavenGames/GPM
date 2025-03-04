/*
Copyright © 2025 Paxton Terry <info@grayhavengames.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"GPM/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

// Flags for updating configuration
var githubUsername string
var defaultRepositoryURL string

// configUpdateCmd represents the `gpm config update` command
var configUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update user configuration values",
	Long: `Allows updating stored user configuration values such as:
  - GitHub username
  - Default repository URL`,
	Run: func(cmd *cobra.Command, args []string) {
		updated := false

		// Update GitHub Username
		if githubUsername != "" {
			n := config.NewGithubUsername()
			err := n.Update(githubUsername)
			if err != nil {
				fmt.Println("Could not update config file: ", err)
			}
			updated = true
		}

		// Update Default URL
		if defaultRepositoryURL != "" {
			n := config.NewDefaultRepositoryURL()
			err := n.Update(defaultRepositoryURL)
			if err != nil {
				fmt.Println("Could not update config file: ", err)
			}
			updated = true
		}

		if !updated {
			fmt.Println("⚠ No updates were provided. Use --github-username or --default-url.")
		}
	},
}

// configCmd represents the base `gpm config` command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage user configuration",
	Long:  "Manage and update the stored user configuration for GPM.",
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Update `update` as a subcommand of `config`
	configCmd.AddCommand(configUpdateCmd)

	// Define flags for `gpm config update`
	configUpdateCmd.Flags().StringVar(&githubUsername, "github-username", "", "Update GitHub username")
	configUpdateCmd.Flags().StringVar(&defaultRepositoryURL, "default-url", "", "Update default repository URL")
}

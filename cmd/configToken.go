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

var token string
var deleteToken bool

// configTokenCmd represents the `gpm config token` command
var configTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage the stored access token",
	Long: `Store, retrieve, or delete your GitHub access token securely.

Examples:
  gpm config token --set <your-token>     # Store a new token
  gpm config token --get                  # Retrieve the stored token
  gpm config token --delete               # Delete the stored token
  `,
	Run: func(cmd *cobra.Command, args []string) {
		manager := &config.AccessTokenManager{}
		username := "your-username" // You can replace this with dynamic retrieval

		if deleteToken {
			err := manager.DeleteToken(username)
			if err != nil {
				fmt.Println("⚠ Error deleting token:", err)
			} else {
				fmt.Println("✔ Token deleted successfully!")
			}
			return
		}

		if token != "" {
			err := manager.SaveToken(username, token)
			if err != nil {
				fmt.Println("⚠ Error storing token:", err)
			} else {
				fmt.Println("✔ Token stored successfully!")
			}
			return
		}

		// Default case: Retrieve the token
		retrievedToken, err := manager.GetToken(username)
		if err != nil {
			fmt.Println("⚠ Error retrieving token:", err)
		} else {
			fmt.Println("✔ Retrieved token:", retrievedToken)
		}
	},
}

func init() {
	configCmd.AddCommand(configTokenCmd)

	// Update flags for storing and deleting tokens
	configTokenCmd.Flags().StringVar(&token, "set", "", "Set a new access token")
	configTokenCmd.Flags().BoolVar(&deleteToken, "delete", false, "Delete the stored access token")
}

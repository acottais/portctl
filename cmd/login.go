// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"syscall"

	"github.com/acottais/portctl/portainer"
	portainerapi "github.com/acottais/portctl/portainerapi/client"
	"github.com/acottais/portctl/portainerapi/client/auth"
	"github.com/acottais/portctl/portainerapi/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login API_URL",
	Short: "log in to a portainer api server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Portainer api server is required")
		}

		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			reader := bufio.NewReader(os.Stdin)
			print("Username: ")
			readname, _ := reader.ReadString('\n')
			username = strings.TrimRight(readname, "\r\n")
		}
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			print("Password: ")
			bytesPassword, _ := terminal.ReadPassword(int(syscall.Stdin))
			password = string(bytesPassword)
			fmt.Println()
		}

		apiURL, err := url.Parse(args[0])
		basetransport := httptransport.New(apiURL.Host, apiURL.Path, []string{apiURL.Scheme})
		if unsecure {
			basetransport.Transport, _ = httptransport.TLSTransport(httptransport.TLSClientOptions{
				InsecureSkipVerify: true,
			})
		}
		portainercli := portainerapi.New(basetransport, strfmt.Default)
		basic, _ := cmd.Flags().GetBool("basic")
		var mode, token string
		if basic {
			basetransport.DefaultAuthentication = httptransport.BasicAuth(username, password)
			_, err := portainercli.Status.StatusInspect(nil, nil)
			if err != nil {
				fmt.Println("Login failed")
				fmt.Println(err)
				return nil
			}
			mode = "basic"
			token = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		} else {
			body := new(models.AuthenticateUserRequest)
			body.Password = &password
			body.Username = &username
			params := auth.NewAuthenticateUserParams()
			params.SetBody(body)
			auth, err := portainercli.Auth.AuthenticateUser(params)
			if err != nil {
				fmt.Println("Login failed")
				fmt.Println(err)
				return nil
			}
			mode = "token"
			token = auth.Payload.Jwt

		}
		// login successfull, save the token
		authContext := portainer.NewAuthContext()
		authContext.SetToken(args[0], mode, token)

		if err = authContext.SaveContext(); err != nil {
			fmt.Println("unable to save token.")
			fmt.Println(err)
			return nil
		}
		fmt.Println("Login successfull")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.Flags().Bool("basic", false, "force the HTTP Basic authentication mode. Use this if Portainer is behind a reverse proxy handling the authentication")
}

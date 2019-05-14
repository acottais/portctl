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
	"fmt"
	"io/ioutil"

	"github.com/acottais/portctl/portainer"
	"github.com/acottais/portctl/portainerapi/models"

	"github.com/acottais/portctl/portainerapi/client/stacks"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// createStackCmd represents the createStack command
var createStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Create a new stack",
	Long:  `Create a new Stack`,
	RunE: func(cmd *cobra.Command, args []string) error {

		portainercli, err := GetPortainerCli()
		if err != nil {

			fmt.Println(err)
			return nil

		}
		dockercli, err := GetDockerCli(portainercli)
		if err != nil {

			fmt.Println("unable to get docker cli")
			return nil

		}
		fmt.Println("creating stack")

		// retrieving swarm context
		swarm, err := portainer.SwarmLookup(dockercli)
		if err != nil {
			fmt.Println("unable to get swarm context")
			return err
		}
		// get stack content
		stackContent, err := ioutil.ReadFile(stackFile)
		if err != nil {
			return err
		}

		// getEnvironment variables
		// env := make([]*models.StackEnv, 2)
		// for i := range env {
		// 	env[i] = new(models.StackEnv)
		// }
		// env[0].Name = "VERSION"
		// env[0].Value = "10.5"
		// env[1].Name = "PASSWORD"
		// env[1].Value = "example"
		var env []*models.StackEnv = nil
		envMap, err := cmd.Flags().GetStringToString("env")
		if err != nil {
			return err
		}
		if envMap != nil {
			env = make([]*models.StackEnv, len(envMap))
			i := 0
			for key, value := range envMap {
				env[i] = new(models.StackEnv)
				env[i].Name = key
				env[i].Value = value
				i++
			}
		}

		// format request body and params
		body := new(models.StackCreateRequest)
		body.Name = &stackName
		body.StackFileContent = string(stackContent)
		body.Env = env
		body.SwarmID = swarm.ID
		stackParams := stacks.NewStackCreateParams()
		stackParams.SetEndpointID(endpoint.ID)
		stackParams.SetMethod("string")
		stackParams.SetBody(body)
		stackParams.SetType(1)
		// validate body
		err = body.Validate(strfmt.Default)
		if err != nil {
			println(err)
			return err
		}
		fmt.Println("calling create stack")
		result, err := portainercli.Stacks.StackCreate(stackParams, nil)

		rc, err := portainer.CheckPortainerError(err)
		switch rc {
		case 409:
			fmt.Printf("Error: Stack with name %s already exist\n", stackName)
			return nil
		case 0:
		default:
			if err != nil {
				return err
			}
		}

		fmt.Printf("Stack %s successfully created\n", result.Payload.Name)
		return nil
	},
}

func init() {
	createCmd.AddCommand(createStackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createStackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createStackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createStackCmd.Flags().StringVar(&stackName, "name", "", "name of the stack")
	cobra.MarkFlagRequired(createStackCmd.Flags(), "name")

	createStackCmd.Flags().StringVar(&stackFile, "file", "", "file containing the stack")
	cobra.MarkFlagRequired(createStackCmd.Flags(), "file")

	createStackCmd.Flags().StringToStringP("env", "e", nil, "environment variable to set when applying the stack content")

}

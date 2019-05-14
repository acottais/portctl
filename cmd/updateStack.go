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
	"github.com/acottais/portctl/portainerapi/client/stacks"
	"github.com/acottais/portctl/portainerapi/models"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// updateStackCmd represents the updateStack command
var updateStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "update a stack",
	Long:  `update a stack.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		envMap, err := cmd.Flags().GetStringToString("env")
		if err != nil {
			return err
		}
		// get stack content
		stackContent, err := ioutil.ReadFile(stackFile)
		if err != nil {
			return err
		}
		// get stack content
		prune, err := cmd.Flags().GetBool("prune")
		if err != nil {
			return err
		}
		return updateStack(string(stackContent), prune, envMap)
	},
}

func init() {
	updateCmd.AddCommand(updateStackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateStackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateStackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateStackCmd.Flags().StringVar(&stackName, "name", "", "name of the stack")
	cobra.MarkFlagRequired(updateStackCmd.Flags(), "name")

	updateStackCmd.Flags().Bool("prune", false, "removes old services")

	updateStackCmd.Flags().StringVar(&stackFile, "file", "", "file containing the stack")
	cobra.MarkFlagRequired(updateStackCmd.Flags(), "file")

	updateStackCmd.Flags().StringToStringP("env", "e", nil, "environment variable to set when applying the stack content")

}

func updateStack(stackContent string, prune bool, envMap map[string]string) error {
	portainercli, err := GetPortainerCli()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	dockercli, err := GetDockerCli(portainercli)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// retrieving swarm context
	swarm, err := portainer.SwarmLookup(dockercli)
	if err != nil {
		fmt.Println("unable to get swarm context")
		fmt.Println(err)
		return nil
	}
	if swarm == nil {
		fmt.Printf("Error: Swarm context not found the endpoint %s, only swarm stacks are supported", endpoint.Name)
		fmt.Println()
		return nil
	}
	stack, err := portainer.StackLookup(portainercli, swarm.ID, stackName)
	if err != nil {
		fmt.Println("unable to get stack context")
		fmt.Println(err)
		return nil
	}
	if stack == nil {
		fmt.Printf("Error: Stack %s not found in the endpoint %s", stackName, endpoint.Name)
		fmt.Println()
		return nil
	}

	var env []*models.StackEnv
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
	body := new(models.StackUpdateRequest)
	body.StackFileContent = stackContent
	body.Env = env
	stackParams := stacks.NewStackUpdateParams()
	stackParams.SetEndpointID(&endpoint.ID)
	stackParams.SetID(stack.ID)
	stackParams.SetBody(body)

	// validate body
	err = body.Validate(strfmt.Default)
	if err != nil {
		println(err)
		return nil
	}
	result, err := portainercli.Stacks.StackUpdate(stackParams, nil)

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

	fmt.Printf("Stack %s successfully updated\n", result.Payload.Name)
	return nil
}

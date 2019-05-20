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
	"html"
	"io"
	"os"
	"text/template"

	"github.com/acottais/portctl/portainerapi/client/stacks"
	"github.com/acottais/portctl/portainerapi/models"

	"github.com/acottais/portctl/portainer"
	"github.com/spf13/cobra"
)

// getStackCmd represents the getStack command
var describeStackCmd = &cobra.Command{
	Use:   "stack",
	Short: "inspect the stack deployed to the instance",
	Long:  `inspect the stack deployed to the instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return printStack(os.Stdout)
	},
}

var stackTemplate = `
Name: {{ .Stack.Name }}
Type: {{ .Stack.Type }}
Environment:{{ range .Stack.Env}}
  {{ .Name }} = {{ .Value }}{{ end }}
Content: 
{{ .Content }}
`

func init() {
	describeCmd.AddCommand(describeStackCmd)
	describeStackCmd.Flags().StringVar(&stackName, "name", "", "name of the stack")
	cobra.MarkFlagRequired(updateStackCmd.Flags(), "name")

}

func printStack(writer io.Writer) error {
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
	fileParams := stacks.NewStackFileInspectParams()
	fileParams.SetID(stack.ID)
	stackcontent, err := portainercli.Stacks.StackFileInspect(fileParams, nil)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("stack").Parse(stackTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(writer, struct {
		Stack   *models.Stack
		Content string
	}{
		Stack:   stack,
		Content: html.UnescapeString(stackcontent.Payload.StackFileContent)})
	if err != nil {
		panic(err)
	}
	return nil
}

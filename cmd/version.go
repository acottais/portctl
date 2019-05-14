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
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Version of portctl
	Version = "devel"
	// BuildDate of the binary
	BuildDate string
	// PortainerVersion used
	PortainerVersion string
	// DockerVersion used
	DockerVersion string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version information",
	Long:  `Display informations ont he cli version and the api versions used`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
portctl: a control tool for Portainer
version: %s
build date: %s
portainer api: %s
docker api: %s
runtime version : %s
		`, Version, BuildDate, PortainerVersion, DockerVersion, runtime.Version())
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/acottais/portctl/portainerapi/models"

	dockerapi "github.com/acottais/portctl/dockerapi/client"
	"github.com/acottais/portctl/portainer"
	portainerapi "github.com/acottais/portctl/portainerapi/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var apiURL *url.URL
var endpoint *models.Endpoint
var portainerToken runtime.ClientAuthInfoWriter
var stackName string
var stackFile string

// Unsecure specify that the connection will not validate certificate
var unsecure bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.portctl.yaml)")
	rootCmd.PersistentFlags().String("host", "", "Portainer api server")
	rootCmd.PersistentFlags().String("endpoint", "", "name of the docker endpoint")
	rootCmd.PersistentFlags().BoolVar(&unsecure, "unsecure", false, "do not validate server certificate")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".portctl" (without extension).
		viper.AddConfigPath(path.Join(home, ".portctl"))
		viper.SetConfigName("config")

		// create token files to keep track of current logins

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// GetPortainerCli returns a portainer client for the current cmd.
func GetPortainerCli() (*portainerapi.Portainer, error) {
	host, err := rootCmd.PersistentFlags().GetString("host")
	if err != nil {

		fmt.Println(err)
		return nil, errors.New("error reading host parameter")

	}
	endpointName, err := rootCmd.PersistentFlags().GetString("endpoint")
	if err != nil {

		fmt.Println(err)
		return nil, errors.New("error reading endpoint parameter")

	}
	authContext := portainer.NewAuthContext()

	apiURL, err = url.Parse(host)
	portainerToken, err = authContext.GetToken(host)
	if err != nil {
		return nil, err
	}
	basetransport := httptransport.New(apiURL.Host, apiURL.Path, []string{apiURL.Scheme})
	basetransport.DefaultAuthentication = portainerToken
	if unsecure {
		basetransport.Transport, _ = httptransport.TLSTransport(httptransport.TLSClientOptions{
			InsecureSkipVerify: true,
		})
	}
	portainercli := portainerapi.New(basetransport, strfmt.Default)
	endpoint, err = portainer.EndpointLookup(portainercli, endpointName)
	_, err = portainer.CheckPortainerError(err)
	if err != nil {
		return nil, err
	}
	if endpoint == nil {
		return nil, errors.New("endpoint not found")
	}
	return portainercli, nil
}

// GetDockerCli returns a docker client based on the current portainer context
func GetDockerCli(portainercli *portainerapi.Portainer) (*dockerapi.DockerEngine, error) {
	path := portainer.GetDockerEndpoint(apiURL.Path, endpoint.ID)

	transport := httptransport.New(apiURL.Host, path, []string{apiURL.Scheme})
	transport.DefaultAuthentication = portainerToken
	if unsecure {
		transport.Transport, _ = httptransport.TLSTransport(httptransport.TLSClientOptions{
			InsecureSkipVerify: true,
		})
	}
	return dockerapi.New(transport, strfmt.Default), nil

}

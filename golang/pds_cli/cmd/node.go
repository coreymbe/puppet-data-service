/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"log"

	client "github.com/puppetlabs/puppet-data-service/golang/pkg/pds_go_client"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Operations on nodes",
}

var listNodesCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := pdsClient.GetAllNodesWithResponse(context.Background())
		if err != nil {
			log.Fatalf("Couldn't get nodes %s", err)
		}
		dump(response.JSON200)
	},
}

var getNodeCmd = &cobra.Command{
	Use:   "get NODENAME",
	Args:  cobra.ExactArgs(1),
	Short: "Retrieve node with nodename NODENAME",
	Run: func(cmd *cobra.Command, args []string) {
		// username, _ := cmd.Flags().GetString("username")
		nodename := args[0]
		response, err := pdsClient.GetNodeByNameWithResponse(context.Background(), client.NodeName(nodename))
		if err != nil {
			log.Fatalf("Couldn't get user %s: %s", nodename, err)
		}
		if response.HTTPResponse.StatusCode > 299 {
			log.Fatalf("Request failed with status code: %d and\nbody: %s\n", response.HTTPResponse.StatusCode, response.Body)
		}
		dump(response.JSON200)
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(listNodesCmd)
	nodeCmd.AddCommand(getNodeCmd)
}

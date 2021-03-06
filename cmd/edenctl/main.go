/*
Copyright 2018 Alexander Trost. All rights reserved.

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

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	nodes_v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha"
	"github.com/galexrt/edenconfmgmt/pkg/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "edenctl",
		Short: "Configuration management with automatic clustering, events and stuff.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger = common.GetLogger()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := []grpc.DialOption{
				grpc.WithInsecure(),
			}
			grpcClient, err := grpc.Dial("127.0.0.1:1337", opts...)
			if err != nil {
				log.Fatalf("fail to dial: %v", err)
			}
			defer grpcClient.Close()
			nodesClient := nodes_v1alpha.NewNodesClient(grpcClient)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			in := &nodes_v1alpha.ListRequest{}
			listResponse, err := nodesClient.List(ctx, in)
			fmt.Printf("Nodes.List(): %+v - %+v\n", listResponse, err)
			return nil
		},
	}
	// Logger for logging.
	logger *zap.Logger
)

func main() {
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("failed to execute root cmd", zap.Error(err))
		os.Exit(1)
	}
	os.Exit(0)
}

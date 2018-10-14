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
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/spf13/cobra"
)

var logger = capnslog.NewPackageLogger("github.com/galexrt/edenconfmgmt", "cmd/edenconfmgmt")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "edenconfmgmt",
	Short: "Configuration management with automatic clustering, events and stuff.",
}

func main() {
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	os.Exit(0)
}

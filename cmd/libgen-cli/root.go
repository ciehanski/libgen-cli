// Copyright © 2019 Antoine Chiny <antoine.chiny@inria.fr>
// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
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

package libgen_cli

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "libgen",
	Short: "A command line interface to access Library Genesis' library.",
	Long: `libgen-cli queries Library Genesis, lists all results of a specific query, 
	and makes them available for download. Simple and easy.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	// Add all subcommands to root cmd
	rootCmd.AddCommand(dbdumpsCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(downloadAllCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(statusCmd)

	if len(os.Args) < 2 {
		if err := rootCmd.Help(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	if os.Args[1] == "-v" || os.Args[1] == "version" || os.Args[1] == "--version" {
		fmt.Printf("libgen-cli %v\n", libgen.Version)
		os.Exit(0)
	}

	// Execute libgen-cli cmd
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

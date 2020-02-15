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
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var downloadAllOutput string

var downloadAllCmd = &cobra.Command{
	Use:   "download-all",
	Short: "",
	Long:  ``,
	Example: "libgen download-all kubernetes",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		searchQuery := strings.Join(args, " ")
		fmt.Printf("++ Searching for: %s\n", searchQuery)

		books, err := libgen.Search(searchQuery, 10, false, false, "")
		if err != nil {
			log.Fatalf("error completing search query: %v", err)
		}

		var wg sync.WaitGroup
		for _, book := range books {
			wg.Add(1)
			go func() {
				if err := libgen.DownloadBook(book, downloadAllOutput); err != nil {
					fmt.Printf("error downloading %v: %v\n", book.Title, err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(downloadAllCmd)
	downloadAllCmd.Flags().StringVarP(&downloadAllOutput, "output", "o", "", "where you want " +
		"libgen-cli to save your download.")
}

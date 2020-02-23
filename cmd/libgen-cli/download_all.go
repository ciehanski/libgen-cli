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

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var downloadAllOutput string
var downloadAllResults int

var downloadAllCmd = &cobra.Command{
	Use:     "download-all",
	Short:   "Downloads all found resources for a specified query.",
	Long:    `Searches for a specific query and downloads all the results found.`,
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

		books, err := libgen.Search(
			searchQuery,
			libgen.GetWorkingMirror(libgen.SearchMirrors),
			downloadAllResults,
			false,
			false,
			"",
		)
		if err != nil {
			log.Fatalf("error completing search query: %v", err)
		}

		// TODO: fix; works outside of goroutine when run synchronously
		var wg sync.WaitGroup
		for _, book := range books {
			wg.Add(1)
			if err := libgen.GetDownloadURL(&book); err != nil {
				log.Println(err)
				continue
			}
			go func() {
				defer wg.Done()
				if err := libgen.DownloadBook(book, downloadAllOutput); err != nil {
					fmt.Printf("error downloading %v: %v\n", book.Title, err)
				}
			}()
		}
		wg.Wait()

		fmt.Printf("\n%s\n", color.GreenString("[OK]"))
	},
}

func init() {
	downloadAllCmd.Flags().StringVarP(&downloadAllOutput, "output", "o", "", "where "+
		"you want libgen-cli to save your download.")
	downloadAllCmd.Flags().IntVarP(&downloadAllResults, "results", "r", 10, "controls "+
		"how many query results are displayed.")
}

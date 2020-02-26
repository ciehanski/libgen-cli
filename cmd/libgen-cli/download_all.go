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
	"os"
	"runtime"
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
				fmt.Printf("error displaying CLI help: %v\n", err)
			}
			os.Exit(0)
		}

		// Join args for complete search query in case
		// it contains spaces
		searchQuery := strings.Join(args, " ")
		fmt.Printf("++ Downloading all for: %s\n", searchQuery)

		books, err := libgen.Search(&libgen.SearchOptions{
			Query:        searchQuery,
			SearchMirror: libgen.GetWorkingMirror(libgen.SearchMirrors),
			Results:      downloadAllResults,
		})
		if err != nil {
			fmt.Printf("error completing search query: %v\n", err)
		}

		// TODO: fix; works outside of goroutine when run synchronously
		var wg sync.WaitGroup
		for _, book := range books {
			if err := libgen.GetDownloadURL(book); err != nil {
				fmt.Printf("error getting download DownloadURL: %v\n", err)
				continue
			}
			fmt.Println(book.DownloadURL)
			wg.Add(1)
			go func() {
				if err := libgen.DownloadBook(*book, downloadAllOutput); err != nil {
					fmt.Printf("error downloading %v: %v\n", book.Title, err)
				}
				wg.Done()
			}()
		}
		wg.Wait()

		if runtime.GOOS == "windows" {
			_, err = fmt.Fprintf(color.Output, "\n%s\n", color.GreenString("[DONE]"))
			if err != nil {
				fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
			}
		} else {
			fmt.Printf("\n%s\n", color.GreenString("[DONE]"))
		}
	},
}

func init() {
	downloadAllCmd.Flags().StringVarP(&downloadAllOutput, "output", "o", "", "where "+
		"you want libgen-cli to save your download.")
	downloadAllCmd.Flags().IntVarP(&downloadAllResults, "results", "r", 10, "controls "+
		"how many query results are displayed.")
}

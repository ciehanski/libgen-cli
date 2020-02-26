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
			os.Exit(1)
		}

		// Get flags
		results, err := cmd.Flags().GetInt("results")
		if err != nil {
			fmt.Printf("error getting results flag: %v\n", err)
		}
		requireAuthor, err := cmd.Flags().GetBool("require-author")
		if err != nil {
			fmt.Printf("error getting require-author flag: %v\n", err)
		}
		extension, err := cmd.Flags().GetString("extension")
		if err != nil {
			fmt.Printf("error getting extension flag: %v\n", err)
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("error getting output flag: %v\n", err)
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Printf("error getting output flag: %v\n", err)
		}

		// Join args for complete search query in case
		// it contains spaces
		searchQuery := strings.Join(args, " ")
		fmt.Printf("++ Downloading all for: %s\n", searchQuery)

		books, err := libgen.Search(&libgen.SearchOptions{
			Query:         searchQuery,
			SearchMirror:  libgen.GetWorkingMirror(libgen.SearchMirrors),
			Results:       results,
			RequireAuthor: requireAuthor,
			Extension:     extension,
			Year:          year,
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
			wg.Add(1)
			go func() {
				if err := libgen.DownloadBook(book, output); err != nil {
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
	downloadAllCmd.Flags().IntP("results", "r", 10, "controls "+
		"how many query results are displayed.")
	searchCmd.Flags().BoolP("require-author", "a", false, "controls "+
		"if the query results will return any media without a listed author.")
	searchCmd.Flags().StringP("extension", "e", "", "controls if the query "+
		"results will return any media with a certain file extension.")
	downloadAllCmd.Flags().StringP("output", "o", "", "where "+
		"you want libgen-cli to save your download.")
	downloadAllCmd.Flags().IntP("year", "y", 0, "filters search query results by the "+
		"year provided.")
}

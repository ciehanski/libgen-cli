// Copyright © 2019 Ryan Ciehanski <ryan@ciehanski.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
		extension, err := cmd.Flags().GetStringSlice("extension")
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
		publisher, err := cmd.Flags().GetString("publisher")
		if err != nil {
			fmt.Printf("error getting publisher flag: %v\n", err)
		}
		language, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Printf("error getting language flag: %v\n", err)
		}
		useIpfs, err := cmd.Flags().GetBool("ipfs-mirrors")
		if err != nil {
			fmt.Printf("error getting ipfs-mirrors flag: %v\n", err)
		}
		sortBy, err := cmd.Flags().GetString("sort-by")
		if err != nil {
			fmt.Printf("error getting sort-by flag: %v\n", err)
		}
		sortASC, err := cmd.Flags().GetBool("sort-asc")
		if err != nil {
			fmt.Printf("error getting sort-asc flag: %v\n", err)
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
			Publisher:     publisher,
			Language:      language,
			SortBy:        sortBy,
			SortASC:       sortASC,
		})
		if err != nil {
			fmt.Printf("error completing search query: %v\n", err)
			os.Exit(1)
		}

		var wg sync.WaitGroup
		bChan := make(chan *libgen.Book, results)
		for _, book := range books {
			if err := libgen.GetDownloadURL(book, useIpfs); err != nil {
				fmt.Printf("error getting download DownloadURL: %v\n", err)
				continue
			}
			wg.Add(1)
			bChan <- book
			go func() {
				curBook := <-bChan
				if useIpfs {
					if err := libgen.DownloadBookIPFS(curBook, output); err != nil {
						fmt.Printf("error downloading %v: %v\n", curBook.Title, err)
					}
				} else {
					if err := libgen.DownloadBook(curBook, output); err != nil {
						fmt.Printf("error downloading %v: %v\n", curBook, err)
					}
				}

				wg.Done()
			}()
		}
		wg.Wait()
		close(bChan)

		if runtime.GOOS == "windows" {
			_, err = fmt.Fprintf(color.Output, "%s\n", color.GreenString("[DONE]"))
			if err != nil {
				fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("%s\n", color.GreenString("[DONE]"))
		}
	},
}

func init() {
	downloadAllCmd.Flags().IntP("results", "r", 10, "controls "+
		"how many query results are displayed.")
	downloadAllCmd.Flags().BoolP("require-author", "a", false, "controls if the query "+
		"results will return any media without a listed author.")
	downloadAllCmd.Flags().StringSliceP("extension", "e", []string{""}, "controls if the query results "+
		"will return any media with a certain file extension.")
	downloadAllCmd.Flags().StringP("output", "o", "", "where you want libgen-cli to "+
		"save your download.")
	downloadAllCmd.Flags().IntP("year", "y", 0, "filters search query results by the "+
		"year provided.")
	downloadAllCmd.Flags().StringP("publisher", "p", "", "filters search query "+
		"results by the publisher provided")
	downloadAllCmd.Flags().StringP("language", "l", "", "filters search query "+
		"results by the language provided")
	downloadAllCmd.Flags().BoolP("ipfs-mirrors", "i", false, "enforces libgen-cli to download "+
		"results via IPFS mirrors instead of HTTP(S) mirrors.")
	downloadAllCmd.Flags().StringP("sort-by", "s", "", "sorts the queried results "+
		"by the specified string. (id, title, author, pub, year, lang, size, ext)")
	downloadAllCmd.Flags().Bool("sort-asc", true, "sorts the queried results "+
		"by ascension or descension.")
}

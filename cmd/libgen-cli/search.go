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
	"strings"

	"github.com/ciehanski/libgen-cli/libgen"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for content hosted by Library Genesis",
	Long: `
	Search pattern and get a list of hash map urls to it, and show
	formatted title + link`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			books         []libgen.Book
			bookSelection []string
			pBookFormat   string
			selectedBook  libgen.Book
		)

		if len(args) < 1 {
			log.Fatal("Error: Search needs a pattern for the command")
		}

		// Parse additional flags
		queryResults, err := rootCmd.Flags().GetInt("results")
		if err != nil {
			log.Fatalf("error parsing additional flags: %v", err)
		}

		searchQuery := strings.Join(args, " ")
		log.Printf("++ Searching: %s\n", searchQuery)

		hashes, err := libgen.Search(searchQuery, queryResults)
		if err != nil {
			log.Fatalf("error completing search query: %v", err)
		}

		books, err = libgen.GetDetails(hashes)
		if err != nil {
			log.Fatalf("error retrieving results from LibGen API: %v", err)
		}

		for _, b := range books {
			selectChoice := fmt.Sprintf("%8s ", b.Id)
			selectChoice += fmt.Sprintf("%-4s ", b.Extension)
			if len(b.Title) > libgen.TitleMaxLength {
				pBookFormat = b.Title[:libgen.TitleMaxLength]
			} else {
				pBookFormat = b.Title
			}
			selectChoice += fmt.Sprintf("%s", pBookFormat)
			bookSelection = append(bookSelection, selectChoice)
		}

		prompt := promptui.Select{
			Label: "Select Book: ",
			Items: bookSelection,
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("error selecting book: %v\n", err)
		}

		for i, b := range bookSelection {
			if b == result {
				selectedBook = books[i]
				break
			}
		}

		if err := libgen.DownloadBook(selectedBook); err != nil {
			log.Fatalf("error downloading book: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.Flags().IntP("results", "r", 10, "Controls how many "+
		"query results are displayed.")
}

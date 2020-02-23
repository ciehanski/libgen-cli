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
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var (
	//mediaType string
	resultsFlag   int
	requireAuthor bool
	extension     string
	searchOutput  string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Query all content hosted by Library Genesis.",
	Long:    `Searches for all resources that result from the provided query and then provides them for download.`,
	Example: "libgen search kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			books         []libgen.Book
			bookSelection []string
			pBookFormat   string
			selectedBook  libgen.Book
		)

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
			resultsFlag,
			true,
			requireAuthor,
			extension,
		)
		if err != nil {
			log.Fatalf("error completing search query: %v", err)
		}
		if len(books) == 0 {
			fmt.Print("\nNo results found.\n")
			os.Exit(0)
		}

		for _, b := range books {
			selectChoice := fmt.Sprintf("%8s ", color.New(color.FgHiBlue).Sprintf(b.ID))
			selectChoice += fmt.Sprintf("%-4s ", color.New(color.FgRed).Sprintf(b.Extension))
			if len(b.Title) > libgen.TitleMaxLength {
				pBookFormat = b.Title[:libgen.TitleMaxLength]
			} else {
				pBookFormat = b.Title
			}
			selectChoice += fmt.Sprintf("%s", pBookFormat)
			bookSelection = append(bookSelection, selectChoice)
		}

		promptTemplate := &promptui.SelectTemplates{
			Active: `▸ {{ .Title | cyan | bold }}{{ if .Title }} ({{ .Title }}){{end}}`,
			//Inactive: `  {{ .Title | cyan }}{{ if .Title }} ({{ .Title }}){{end}}`,
			Selected: `{{ "✔" | green }} %s: {{ .Title | cyan }}{{ if .Title }} ({{ .Title }}){{end}}`,
		}

		prompt := promptui.Select{
			Label:     "Select Book",
			Items:     bookSelection,
			Templates: promptTemplate,
			Size:      resultsFlag,
		}

		fmt.Println(strings.Repeat("-", 80))

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}

		for i, b := range bookSelection {
			if b == result {
				selectedBook = books[i]
				break
			}
		}

		fmt.Printf("Download started for: %s by %s\n", selectedBook.Title, selectedBook.Author)

		if err := libgen.GetDownloadURL(&selectedBook); err != nil {
			log.Fatalf("error retrieving valid download URL: %v", err)
		}
		if err := libgen.DownloadBook(selectedBook, searchOutput); err != nil {
			log.Fatalf("error downloading %v: %v", selectedBook.Title, err)
		}

		fmt.Printf("\n%s %s by %s.%s\n", color.GreenString("[OK]"),
			selectedBook.Title, selectedBook.Author, selectedBook.Extension)
	},
}

func init() {
	//searchCmd.Flags().StringVarP(&mediaType, "media", "m", "libgen", "controls what "+
	//	"type of media will be queried for. Ex: fiction, comics, scientific papers, etc.")
	searchCmd.Flags().IntVarP(&resultsFlag, "results", "r", 10, "controls how many "+
		"query results are displayed.")
	searchCmd.Flags().BoolVarP(&requireAuthor, "require-author", "a", false, "controls "+
		"if the query results will return any media without a listed author.")
	searchCmd.Flags().StringVarP(&extension, "ext", "e", "", "controls if the query "+
		"results will return any media with a certain file extension.")
	searchCmd.Flags().StringVarP(&searchOutput, "output", "o", "", "where you want "+
		"libgen-cli to save your download.")
}

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

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var downloadOutput string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a specific resource by hash.",
	Long:  ``,
	Example: "libgen download 2F2DBA2A621B693BB95601C16ED680F8",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		book, err := libgen.GetDetails(args, true, false, "")
		if err != nil {
			log.Fatalf("error retrieving results from LibGen API: %v", err)
		}

		fmt.Printf("Download started for: %s by %s\n", book[0].Title, book[0].Author)

		if err := libgen.DownloadBook(book[0], downloadOutput); err != nil {
			log.Fatalf("error downloading %v: %v", book[0].Title, err)
		}

		fmt.Printf("%s %s", color.GreenString("[OK]"), book[0].Title+book[0].Extension)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&downloadOutput, "output", "o", "", "where you want " +
		"libgen-cli to save your download.")
}

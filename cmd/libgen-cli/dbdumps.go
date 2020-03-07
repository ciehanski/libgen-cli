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
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var dbdumpsCmd = &cobra.Command{
	Use:     "dbdumps",
	Short:   "Allows users to download any selection of Library Genesis' database dumps.",
	Long:    `A collection of Library Genesis' compressed SQL database dumps can be downloaded using this command.`,
	Example: "libgen dbdumps",
	Run: func(cmd *cobra.Command, args []string) {

		// Don't allow args
		if len(args) != 0 {
			if err := cmd.Help(); err != nil {
				fmt.Printf("error displaying CLI help: %v\n", err)
			}
			os.Exit(1)
		}

		// Get flags
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("error getting output flag: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("++ Retrieving all database dumps...")

		mirror := libgen.GetWorkingMirror(libgen.SearchMirrors)

		r, err := http.Get(mirror.String() + "/dbdumps/")
		if err != nil {
			fmt.Printf("error reaching mirror: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error reading response: %v\n", err)
			os.Exit(1)
		}

		dbdumps := libgen.ParseDbdumps(b)
		if dbdumps == nil {
			fmt.Println("\nerror parsing dbdumps. No dbdumps found.")
			os.Exit(1)
		}
		if len(dbdumps) == 0 {
			fmt.Println("\nNo results found.")
			os.Exit(1)
		}

		promptTemplate := &promptui.SelectTemplates{
			Active: `▸ {{ .Title | cyan | bold }}{{ if .Title }} ({{ .Title }}){{end}}`,
			//Inactive: `  {{ .Title | cyan }}{{ if .Title }} ({{ .Title }}){{end}}`,
			Selected: `{{ "✔" | green }} %s: {{ .Title | cyan }}{{ if .Title }} ({{ .Title }}){{end}}`,
		}

		prompt := promptui.Select{
			Label:     "Select Database Dump",
			Items:     dbdumps,
			Templates: promptTemplate,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		var selectedDbdump string
		for i, b := range dbdumps {
			if b == result {
				selectedDbdump = dbdumps[i]
				break
			}
		}

		fmt.Printf("Download started for: %s\n", libgen.RemoveQuotes(selectedDbdump))

		if err := libgen.DownloadDbdump(selectedDbdump, output); err != nil {
			fmt.Printf("error download dbdump: %v\n", err)
			os.Exit(1)
		}

		if runtime.GOOS == "windows" {
			_, err = fmt.Fprintf(color.Output, "\n%s %s\n", color.GreenString("[OK]"), selectedDbdump)
			if err != nil {
				fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
			}
		} else {
			fmt.Printf("\n%s %s\n", color.GreenString("[OK]"), selectedDbdump)
		}
	},
}

func init() {
	dbdumpsCmd.Flags().StringP("output", "o", "", "where you want libgen-cli to "+
		"save your download.")
}

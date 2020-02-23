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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var dbdumpsOutput string

var dbdumpsCmd = &cobra.Command{
	Use:     "dbdumps",
	Short:   "Allows users to download any selection of Library Genesis' database dumps.",
	Long:    `A collection of Library Genesis' compressed SQL database dumps can be downloaded using this command.`,
	Example: "libgen dbdumps",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := http.Get("http://gen.lib.rus.ec/dbdumps/")
		if err != nil {
			log.Fatalf("error reaching mirror: %v", err)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("error reading response: %v", err)
		}

		dbdumps := parseDbdumps(string(b))
		if dbdumps == nil {
			log.Fatal("error parsing dbdumps. No dbdumps found.")
		}
		if len(dbdumps) == 0 {
			fmt.Print("\nNo results found.\n")
			os.Exit(0)
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

		fmt.Println(strings.Repeat("-", 80))

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}

		var selectedDbDump string
		for i, b := range dbdumps {
			if b == result {
				selectedDbDump = dbdumps[i]
				break
			}
		}

		fmt.Printf("Download started for: %s\n", selectedDbDump)

		if err := downloadDbDump(selectedDbDump); err != nil {
			log.Fatalf("error download dbdump: %v", err)
		}

		fmt.Printf("\n%s %s\n", color.GreenString("[OK]"), selectedDbDump)
	},
}

func downloadDbDump(filename string) error {
	filename = removeQuotes(filename)
	r, err := http.Get(fmt.Sprintf("http://gen.lib.rus.ec/dbdumps/%s", filename))
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		filesize := r.ContentLength
		bar := pb.Full.Start64(filesize)

		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		out, err := os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
		if err != nil {
			return err
		}

		_, err = io.Copy(out, bar.NewProxyReader(r.Body))
		if err != nil {
			return err
		}

		bar.Finish()

		if err := out.Close(); err != nil {
			return err
		}
		if err := r.Body.Close(); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unable to reach mirror: HTTP %v", r.StatusCode)
	}

	return nil
}

func parseDbdumps(response string) []string {
	re := regexp.MustCompile(`(["])(.*?\.(rar|sql.gz))"`)
	dbdumps := re.FindAllString(response, -1)

	for _, dbdump := range dbdumps {
		dbdump = removeQuotes(dbdump)
	}

	return dbdumps
}

func removeQuotes(s string) string {
	s = s[1:]
	s = s[:len(s)-1]
	return s
}

func init() {
	dbdumpsCmd.Flags().StringVarP(&dbdumpsOutput, "output", "o", "", "where you want "+
		"libgen-cli to save your download.")
}

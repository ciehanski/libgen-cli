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
	"regexp"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "Download a specific resource by hash.",
	Long:    `Use this command if you already know the hash of the specific resource you'd like to download.'`,
	Example: "libgen download 2F2DBA2A621B693BB95601C16ED680F8",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			if err := cmd.Help(); err != nil {
				fmt.Printf("error displaying CLI help: %v\n", err)
			}
			os.Exit(1)
		}
		// Ensure provided entry is valid MD5 hash
		re := regexp.MustCompile(libgen.SearchMD5)
		if !re.MatchString(args[0]) {
			fmt.Printf("Please provide a valid MD5 hash\n")
			os.Exit(1)
		}

		// Get flags
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("error getting output flag: %v\n", err)
		}
		useIpfs, err := cmd.Flags().GetBool("ipfs-mirrors")
		if err != nil {
			fmt.Printf("error getting ipfs-mirrors flag: %v\n", err)
		}

		if len(args) == 1 {
			fmt.Printf("++ Searching for: %s\n", args[0])
		} else {
			fmt.Printf("++ Searching for: MD5s\n")
		}

		searchMirror := libgen.GetWorkingMirror(libgen.SearchMirrors)
		bookDetails, err := libgen.GetDetails(&libgen.GetDetailsOptions{
			Hashes:       args,
			SearchMirror: searchMirror,
			Print:        true,
		})
		if err != nil {
			// If error, try another mirror before exiting
			secondaryMirror := libgen.GetWorkingMirror(libgen.SearchMirrors)
			for secondaryMirror == searchMirror {
				secondaryMirror = libgen.GetWorkingMirror(libgen.SearchMirrors)
			}
			bookDetails, err = libgen.GetDetails(&libgen.GetDetailsOptions{
				Hashes:       args,
				SearchMirror: secondaryMirror,
				Print:        true,
			})
			if err != nil {
				log.Fatalf("error retrieving results from LibGen API: %v", err)
			}
		}

		for _, book := range bookDetails {

			fmt.Println(strings.Repeat("-", 80))
			fmt.Printf("Download started for: %s by %s\n", book.Title, book.Author)

			if err := libgen.GetDownloadURL(book, useIpfs); err != nil {
				fmt.Printf("error getting download URL: %v\n", err)
				os.Exit(1)
			}
			if useIpfs {
				if err := libgen.DownloadBookIPFS(book, output); err != nil {
					fmt.Printf("error downloading %v: %v\n", book.Title, err)
					os.Exit(1)
				}
			} else {
				if err := libgen.DownloadBook(book, output); err != nil {
					fmt.Printf("error downloading %v: %v\n", book.Title, err)
					os.Exit(1)
				}
			}

			if runtime.GOOS == "windows" {
				_, err = fmt.Fprintf(color.Output, "%s %s by %s.%s", color.GreenString("[OK]"),
					book.Title, book.Author, book.Extension)
				if err != nil {
					fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
					os.Exit(1)
				}
			} else {
				fmt.Printf("%s %s by %s.%s\n", color.GreenString("[OK]"),
					book.Title, book.Author, book.Extension)
			}

		}

	},
}

func init() {
	downloadCmd.Flags().StringP("output", "o", "", "where you want "+
		"libgen-cli to save your download.")
	downloadCmd.Flags().BoolP("ipfs-mirrors", "i", false, "enforces libgen-cli to download "+
		"results via IPFS mirrors instead of HTTP(S) mirrors.")
}

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
	"net/http"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/ciehanski/libgen-cli/libgen"
)

var mirrorSwitch string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Checks the status of Library Genesis' mirrors.",
	Long:    `Checks the status of all Library Genesis search mirrors as well as all download mirrors.`,
	Example: `libgen status`,
	Run: func(cmd *cobra.Command, args []string) {
		switch mirrorSwitch {
		case "download":
			for _, url := range libgen.DownloadMirrors {
				status := libgen.CheckMirror(url)
				if status == http.StatusOK {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.GreenString("[OK]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.GreenString("[OK]"), url.String())
					}
				} else {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.RedString("[FAIL]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.RedString("[FAIL]"), url.String())
					}
				}
			}
		case "search":
			for _, url := range libgen.SearchMirrors {
				status := libgen.CheckMirror(url)
				if status == http.StatusOK {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.GreenString("[OK]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.GreenString("[OK]"), url.String())
					}
				} else {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.RedString("[FAIL]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.RedString("[FAIL]"), url.String())
					}
				}
			}
		default:
			for _, url := range libgen.SearchMirrors {
				status := libgen.CheckMirror(url)
				if status == http.StatusOK {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.GreenString("[OK]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.GreenString("[OK]"), url.String())
					}
				} else {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.RedString("[FAIL]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.RedString("[FAIL]"), url.String())
					}
				}
			}
			for _, url := range libgen.DownloadMirrors {
				status := libgen.CheckMirror(url)
				if status == http.StatusOK {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.GreenString("[OK]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.GreenString("[OK]"), url.String())
					}
				} else {
					if runtime.GOOS == "windows" {
						_, err := fmt.Fprintf(color.Output, "%s %s\n", color.RedString("[FAIL]"), url.String())
						if err != nil {
							fmt.Printf("error writing to Windows os.Stdout: %v\n", err)
						}
					} else {
						fmt.Printf("%s %s\n", color.RedString("[FAIL]"), url.String())
					}
				}
			}
		}
	},
}

func init() {
	statusCmd.Flags().StringVarP(&mirrorSwitch, "mirror", "m", "", "Choose a specific "+
		"collection of mirrors to check status.")
}

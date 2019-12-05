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
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks status of Library Genesis' mirrors.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		checkMirror("http://libgen.lc")
		checkMirror("http://gen.lib.rus.ec")
		checkMirror("https://93.174.95.29")
		checkMirror("http://booksdl.org")
		checkMirror("https://b-ok.cc")
	},
}

func checkMirror(url string) {
	r, err := http.Get(url)
	if err != nil || r.StatusCode != http.StatusOK {
		log.Printf("[FAIL] %s\n", url)
	} else {
		log.Printf("[OK] %s\n", url)
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

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

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/ciehanski/libgen-cli/cmd/libgen-cli"
)

func main() {
	_, err := net.DialTimeout("tcp", "golang.org:80", time.Second*10)
	if err != nil {
		fmt.Println("\nYou need an internet connection to run libgen-cli.")
		os.Exit(1)
	}

	if err := libgen_cli.Execute(); err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

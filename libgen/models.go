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

package libgen

import "net/url"

const (
	SearchHref = "<a href='book/index.php.+</a>"
	SearchMD5  = "[A-Z0-9]{32}"
	booksdlReg = "http://booksdl.org/get\\.php\\?md5=\\w{32}\\&key=\\w{16}"
	bokReg     = `/book/\d{6}/\d{6}`
	//libgenPwReg     = `http://libgen.pw/item/detail/id/\d*$`
	JSONQuery       = "id,title,author,filesize,extension,md5,year"
	TitleMaxLength  = 65
	AuthorMaxLength = 25
)

// Book contains all the pertinent information required for a resource queried by libgen.
type Book struct {
	ID        string
	Title     string
	Author    string
	Filesize  string
	Extension string
	Md5       string
	Year      string
	URL       string
}

// SearchMirrors is used by libgen-cli strictly for querying resources.
var SearchMirrors = []url.URL{
	{
		Scheme: "http",
		Host:   "gen.lib.rus.ec",
	},
	{
		Scheme: "http",
		Host:   "libgen.lc",
	},
	{
		Scheme: "http",
		Host:   "libgen.li",
	},
	{
		Scheme: "https",
		Host:   "libgen.is",
	},
	{
		Scheme: "http",
		Host:   "185.39.10.101",
	},
}

// DownloadMirrors is used by libgen-cli strictly for downloading resources.
var DownloadMirrors = []url.URL{
	SearchMirrors[1],
	SearchMirrors[3],
	{
		Scheme: "https",
		Host:   "93.174.95.29",
	},
	{
		Scheme: "http",
		Host:   "booksdl.org",
	},
	{
		Scheme: "https",
		Host:   "b-ok.cc",
	},
}

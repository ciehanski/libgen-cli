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

const (
	SearchHref     = "<a href='book/index.php.+</a>"
	SearchMD5      = "[A-Z0-9]{32}"
	searchUrl      = "http://booksdl.org/get\\.php\\?md5=\\w{32}\\&key=\\w{16}"
	TitleMaxLength = 60
)

type Book struct {
	Id        string
	Title     string
	Author    string
	Filesize  string
	Extension string
	Md5       string
	Year      string
	Url       string
}

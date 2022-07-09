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

import "time"

const (
	Version           = "v1.0.8"
	SearchHref        = "<a href='book/index.php.+</a>"
	SearchMD5         = "[A-Z0-9]{32}"
	booksdlReg        = `https:\/\/libgen\.rocks\/get\.php\?md5=\w{32}&key=\w{16}`
	libraryLolReg     = `http:\/\/62\.182\.86\.140\/main\/\d{7}\/\w{32}\/.+?(gz|pdf|rar|djvu|epub|chm)`
	dbdumpReg         = `(["])(.*?\.(rar|sql.gz))"`
	JSONQuery         = "id,title,author,filesize,extension,md5,year,language,pages,publisher,edition,coverurl"
	TitleMaxLength    = 68
	AuthorMaxLength   = 25
	HTTPClientTimeout = time.Second * 15
	//UploadUsername    = "genesis"
	//UploadPassword    = "upload"
	//libgenPwReg     = `http://libgen.pw/item/detail/id/\d*$`
)

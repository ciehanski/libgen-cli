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
	Version             = "v1.1.0"
	SearchHref          = "<a href='book/index.php.+</a>"
	SearchMD5           = "[A-Za-z0-9]{32}"
	libgenPMReg         = `get\.php\?md5=\w{32}&key=\w{16}`
	libraryLolReg       = `https://download\.library\.lol/main/\d+/[A-Za-z0-9]+/[^"]+`
	libraryLolIPFSReg   = `https:\/\/gateway\.ipfs\.io\/ipfs\/[A-Za-z0-9_-]+(\?[^"]*)?`
	libraryLolIPFSCFReg = `https:\/\/cloudflare-ipfs\.com\/ipfs\/[A-Za-z0-9_-]+(\?[^"]*)?`
	dbdumpReg           = `(["])(.*?\.(rar|sql.gz))"`
	JSONQuery           = "id,title,author,filesize,extension,md5,year,language,pages,publisher,edition,coverurl"
	TitleMaxLength      = 68
	AuthorMaxLength     = 25
	HTTPClientTimeout   = time.Second * 5
	ipfsReg             = `/ipfs/([a-z0-9]+)`
	//UploadUsername    = "genesis"
	//UploadPassword    = "upload"
	//libgenPwReg     = `http://libgen.pw/item/detail/id/\d*$`
)

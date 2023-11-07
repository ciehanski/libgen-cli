// Copyright © 2023 Ryan Ciehanski <ryan@ciehanski.com>
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

import (
	"strings"
	"testing"
)

func TestDownloadIPFSBook(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getLibraryLolURL(book[0], true); err != nil {
		t.Error(err)
	}
	if err := DownloadBook(book[0], ""); err != nil {
		t.Error(err)
	}
}

func TestGetDownloadIPFSURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := GetDownloadURL(book[0], true); err != nil {
		t.Error(err)
	}
	if book[0].DownloadURL == "" {
		t.Error("download URL empty")
	}
}

func TestGetLibraryLolIPFSURL(t *testing.T) {
	book, err := GetDetails(&GetDetailsOptions{
		Hashes:       []string{"1794743BB21D72736FFE64D66DCA9F0E"},
		SearchMirror: GetWorkingMirror(SearchMirrors),
		Print:        false,
	})
	if err != nil {
		t.Error(err)
	}

	if err := getLibraryLolURL(book[0], true); err != nil {
		t.Error(err)
	}

	if book[0].DownloadURL == "" {
		t.Error("no valid url found")
	}
	if !strings.Contains(book[0].DownloadURL, "https://gateway.ipfs.io/ipfs/bafykbzacectwnzckgcrnozlrkx7j5fbdwlf6qo7whmf2sksafwfwvunazyl4e?filename=") {
		t.Errorf(`got: %s, expected: https://gateway.ipfs.io/ipfs/bafykbzacectwnzckgcrnozlrkx7j5fbdwlf6qo7whmf2sksafwfwvunazyl4e?filename=`, book[0].DownloadURL)
	}
}

func TestGetHrefIPFS(t *testing.T) {
	results := findMatch(libraryLolIPFSReg, []byte(`
<!DOCTYPE HTML>
<html lang="en">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title></title>
<style type="text/css">
table td {
	vertical-align: top;
}
#message {
	width: 400px;
	margin: 0px auto;
	padding: 10px 20px;
	text-align: center;	
	background-color: #0f9d58;
	color: #fff;
	border-radius: 3px;
}
#info {
	max-width: 700px;
	padding: 10px;
	border: 1px solid #C0C0C0;
	font-family: "Arial", "Helvetica", sans-serif;
	font-size: 0.8em;
}
#info img {
	display: block;
	width: 240px;
	max-width: 240px;
	margin: 3px auto;
}
#download {
	text-align: center;
}
#download ul {
	margin: 0.8em 0 0 0;
}
#download ul li {
	display: inline-block;
}
#download ul li a {
	display: block !important;
	min-width: 5.5em;
	margin: 0 5px;
	padding: 4px;
	border: 1px solid blue;
	border-radius: 7px;
	text-decoration: none;
	text-align: center;
}
#download ul li sup>a {
	width: auto;
	border: 0;
}
.adsbygoogle {
	margin: 7px;
}
</style>
<script src="/jquery-latest.min.js"></script>
</head>
<body>
<table width="100%" align="center" border="0">
<tr>
	<td class="ad"></td>
	<td id="info">				<div id="download">
		<h2><a href="https://download.library.lol/main/2596000/a87ede7392897082324a9ac30ffc1999/America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">GET</a></h2>
				<div><em>FASTER</em> Download from an IPFS distributed storage, choose any gateway:</div>
		<ul>
		<li><a href="https://cloudflare-ipfs.com/ipfs/bafykbzacebrrexkkvcb5dmswgoi4c33t4ztocnulzpt4dmg5vifleijmglnvk?filename=America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">Cloudflare</a></li><li><a href="https://gateway.ipfs.io/ipfs/bafykbzacebrrexkkvcb5dmswgoi4c33t4ztocnulzpt4dmg5vifleijmglnvk?filename=America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">IPFS.io</a></li><li><a href="https://gateway.pinata.cloud/ipfs/bafykbzacebrrexkkvcb5dmswgoi4c33t4ztocnulzpt4dmg5vifleijmglnvk?filename=America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">Pinata</a></li><li><a href="http://localhost:8080/ipfs/bafykbzacebrrexkkvcb5dmswgoi4c33t4ztocnulzpt4dmg5vifleijmglnvk?filename=America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">local gateway</a></li>		</ul>
				<h2><a href="http://libgenfrialc7tguyjywa36vtrdcplwpxaw43h6o63dmmwhvavo5rqqd.onion/LG/02596000/a87ede7392897082324a9ac30ffc1999/America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub">download from the Tor mirror</a><br><span style="font-size:70%">(make sure you're accessing via <a href="https://www.howtogeek.com/272049/how-to-access-.onion-sites-also-known-as-tor-hidden-services/">Tor</a>)</span></h2>
		</div>
				<h1>The complete America's test kitchen TV show cookbook 2015</h1>
		<div><img src="/covers/2596000/a87ede7392897082324a9ac30ffc1999-g.jpg" alt="cover"></div>
		<p>Author(s): America's Test Kitchen</p>				<p>Publisher: America's Test Kitchen, Year: 2014</p>
		<p>ISBN: 9781940352169,1940352169</p>		<p style="text-align:center"><a href="https://www.worldcat.org/search?qt=worldcat_org_bks&q=The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015&fq=dt%3Abks">Search in WorldCat</a> | <a href="https://www.goodreads.com/search?utf8=✓&query=The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015">Search in Goodreads</a> | <a href="https://www.abebooks.com/servlet/SearchResults?tn=The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015&pt=book&cm_sp=pan-_-srp-_-ptbook">Search in AbeBooks</a> | <a href="https://www.amazon.com/s/?url=search-alias%3Dstripbooks&field-keywords=The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015">Search in Amazon.com</a></p>
				<div>Description:<br>The ultimate collection of recipes from your favorite TV show. This newly revised edition of The Complete America's Test Kitchen TV Show Cookbook includes all 15 seasons (including 2015) of the hit TV show in a lively collection featuring more than 950 foolproof recipes and dozens of tips and techniques.</div>		</td>
	<td class="ad"></td>
</tr>
</table>
</body>
</html>
    `))
	if results == nil {
		t.Error("empty result")
	}
	if string(results) != "https://gateway.ipfs.io/ipfs/bafykbzacebrrexkkvcb5dmswgoi4c33t4ztocnulzpt4dmg5vifleijmglnvk?filename=America%27s%20Test%20Kitchen%20-%20The%20complete%20America%27s%20test%20kitchen%20TV%20show%20cookbook%202015-America%27s%20Test%20Kitchen%20%282014%29.epub" {
		t.Errorf("incorrect DownloadURL returned. got %s", string(results))
	}
}

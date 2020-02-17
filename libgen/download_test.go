// Copyright © 2020 Ryan Ciehanski <ryan@ciehanski.com>
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

func TestDownloadBook(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := DownloadBook(book[0], ""); err != nil {
		t.Error(err)
	}
}

func TestGetDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := getDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}

	if book[0].URL == "" {
		t.Error("Unable to retrieve download URL")
	}
}

func TestGetBokDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := getBokDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}
	if book[0].URL != "https://b-ok.cc/dl/436993/659204" {
		t.Errorf("got: %s, expected: https://b-ok.cc/dl/436993/659204", book[0].URL)
	}
}

func TestGetBooksdlDownloadURL(t *testing.T) {
	book, err := GetDetails([]string{"2F2DBA2A621B693BB95601C16ED680F8"}, GetWorkingMirror(SearchMirrors), false, false, "")
	if err != nil {
		t.Error(err)
	}

	if err := getBooksdlDownloadURL(&book[0]); err != nil {
		t.Error(err)
	}
	if !strings.Contains(book[0].URL, "http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=") {
		t.Errorf("got: %s, expected: http://booksdl.org/get.php?md5=2f2dba2a621b693bb95601c16ed680f8&key=", book[0].URL)
	}
}

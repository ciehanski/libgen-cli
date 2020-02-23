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

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/cheggaaa/pb/v3"
)

// DownloadBook grabs the download URL for the book requested. First, it queries Booksdl.org and then
// b-ok.cc for valid URL. Then, the download process is initiated with a progress bar displayed to
// the user's CLI.
func DownloadBook(book Book, output string) error {
	var filesize int64
	filename := getBookFilename(book)

	if book.URL == "" {
		if err := GetDownloadURL(&book); err != nil {
			return err
		}
	}

	//client := http.Client{Timeout: time.Second * 5}
	r, err := http.Get(book.URL)
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		var (
			out *os.File
			err error
		)
		filesize = r.ContentLength
		bar := pb.Full.Start64(filesize)

		// check if output folder was provided. If not, create
		// one at the current directory called "libgen."
		var osErr error
		if output == "" {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			if stat, err := os.Stat(fmt.Sprintf("%s/libgen", wd)); err == nil && stat.IsDir() {
				out, osErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			} else {
				if err := os.Mkdir(fmt.Sprintf("%s/libgen", wd), 0700); err != nil {
					return err
				}
				out, osErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			}
			if osErr != nil {
				return err
			}
		} else {
			if stat, err := os.Stat(output); err == nil && stat.IsDir() {
				out, err = os.Create(fmt.Sprintf("%s/%s", output, filename))
				if err != nil {
					return err
				}
			} else {
				return errors.New("invalid output path")
			}
		}

		_, err = io.Copy(out, bar.NewProxyReader(r.Body))
		if err != nil {
			return err
		}

		bar.Finish()
		if err := out.Close(); err != nil {
			return err
		}
		if err := r.Body.Close(); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unable to reach mirror: HTTP %v", r.StatusCode)
	}

	return nil
}

// GetDownloadURL picks a random download mirror to download the specified
// resource from.
func GetDownloadURL(book *Book) error {
	chosenMirror := DownloadMirrors[rand.Intn(2)]

	switch chosenMirror.String() {
	case "http://booksdl.org":
		if err := getBooksdlDownloadURL(book); err != nil {
			if err := getBokDownloadURL(book); err != nil {
				return err
			}
		}
	case "https://b-ok.cc":
		if err := getBokDownloadURL(book); err != nil {
			if err := getBooksdlDownloadURL(book); err != nil {
				return err
			}
		}
	}

	if book.URL == "" {
		return fmt.Errorf("unable to retrieve download link for desired resource")
	}

	return nil
}

func getBooksdlDownloadURL(book *Book) error {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   "libgen.lc",
		Path:   "ads.php",
	}

	q := baseURL.Query()
	q.Set("md5", book.Md5)
	baseURL.RawQuery = q.Encode()

	r, err := http.Get(baseURL.String())
	if err != nil {
		log.Printf("http.Get(%q) error: %v", baseURL, err)
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to connect to mirror: %v", r.StatusCode)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	book.URL = getHref(booksdlReg, string(b))

	if err := r.Body.Close(); err != nil {
		return err
	}

	return nil
}

func getBokDownloadURL(book *Book) error {
	baseURL := url.URL{
		Scheme: "https",
		Host:   "b-ok.cc",
		Path:   "md5/",
	}

	queryURL := baseURL.String() + book.Md5

	r, err := http.Get(queryURL)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to connect to mirror: %v", r.StatusCode)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	downloadURL := getHref(bokReg, string(b))
	if downloadURL == "" {
		return errors.New("no valid download URL found")
	}
	book.URL = "https://b-ok.cc" + downloadURL

	if err := r.Body.Close(); err != nil {
		return err
	}

	return nil
}

func getHref(reg string, response string) string {
	re := regexp.MustCompile(reg)
	matches := re.FindAllString(response, -1)

	if len(matches) > 0 {
		return matches[0]
	}

	return ""
}

func getBookFilename(book Book) string {
	var tmp []string
	tmp = append(tmp, book.Title)
	tmp = append(tmp, fmt.Sprintf(" by %s", book.Author))
	tmp = append(tmp, fmt.Sprintf(".%s", book.Extension))
	return strings.Join(tmp, "")
}

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

// DownloadBook grabs the download DownloadURL for the book requested.
// First, it queries Booksdl.org and then b-ok.cc for valid DownloadURL.
// Then, the download process is initiated with a progress bar displayed to
// the user's CLI.
func DownloadBook(book *Book, output string) error {
	var filesize int64
	filename := getBookFilename(book)

	req, err := http.NewRequest("GET", book.DownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Encoding", "*")
	if strings.Contains(book.PageURL, "b-ok.cc") {
		req.Header.Add("Referer", book.PageURL)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		filesize = r.ContentLength
		bar := pb.Full.Start64(filesize)

		// check if output folder was provided. If not, create
		// one at the current directory called "libgen."
		var mkErr error
		var out *os.File
		if output == "" {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			if stat, err := os.Stat(fmt.Sprintf("%s/libgen", wd)); err == nil && stat.IsDir() {
				out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			} else {
				if err := os.Mkdir(fmt.Sprintf("%s/libgen", wd), 0755); err != nil {
					return err
				}
				out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			}
			if mkErr != nil {
				return mkErr
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
		return fmt.Errorf("unable to reach mirror %v: HTTP %v", req.Host, r.StatusCode)
	}

	return nil
}

// DownloadDbdump downloads the selected database dump from
// Library Genesis.
func DownloadDbdump(filename string, output string) error {
	filename = RemoveQuotes(filename)
	mirror := GetWorkingMirror(SearchMirrors)
	client := http.Client{Timeout: httpClientTimeout}
	r, err := client.Get(fmt.Sprintf("%s/dbdumps/%s", mirror.String(), filename))
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		filesize := r.ContentLength
		bar := pb.Full.Start64(filesize)

		var mkErr error
		var out *os.File
		if output == "" {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			if stat, err := os.Stat(fmt.Sprintf("%s/libgen", wd)); err == nil && stat.IsDir() {
				out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			} else {
				if err := os.Mkdir(fmt.Sprintf("%s/libgen", wd), 0755); err != nil {
					return err
				}
				out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
			}
			if mkErr != nil {
				return mkErr
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
			if err = getBokDownloadURL(book); err != nil {
				return err
			}
		}
	case "https://b-ok.cc":
		if err := getBokDownloadURL(book); err != nil {
			if err = getBooksdlDownloadURL(book); err != nil {
				return err
			}
		}
	}

	if book.DownloadURL == "" {
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
	book.PageURL = baseURL.String()

	client := http.Client{Timeout: httpClientTimeout}
	r, err := client.Get(baseURL.String())
	if err != nil {
		log.Printf("http.Get(%q) error: %v", baseURL, err)
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to reach to mirror %v: %v", baseURL.Host, r.StatusCode)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	book.DownloadURL = findMatch(booksdlReg, b)

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
	book.PageURL = queryURL

	client := http.Client{Timeout: httpClientTimeout}
	resp, err := client.Get(queryURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to reach to mirror %v: %v", baseURL.Host, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	downloadURL := findMatch(bokReg, b)
	if downloadURL == "" {
		return errors.New("no valid download DownloadURL found")
	}

	book.DownloadURL = "https://b-ok.cc" + downloadURL

	if err := checkBokDownloadLimit(book); err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return nil
}

// checkBokDownloadLimit checks the response from the b-ok.cc
// download page and scans it for text stating there have
// been more than 5 downloads from your IP in the past 24
// hours and returns an error if so.
func checkBokDownloadLimit(book *Book) error {
	req, err := http.NewRequest("GET", book.DownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Referer", book.PageURL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	re := regexp.MustCompile("WARNING: There are more than 5 downloads from your IP")
	matches := re.FindAllString(string(b), -1)

	if len(matches) > 0 {
		return errors.New("download limit reached for b-ok.cc")
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return nil
}

// findMatch is a helper function that searches an []byte
// for a specified regex and returns the matches.
func findMatch(reg string, response []byte) string {
	re := regexp.MustCompile(reg)
	match := re.FindString(string(response))

	if match != "" {
		return match
	}

	return ""
}

func getBookFilename(book *Book) string {
	var tmp []string
	tmp = append(tmp, book.Title)
	tmp = append(tmp, fmt.Sprintf(" by %s", book.Author))
	tmp = append(tmp, fmt.Sprintf(".%s", book.Extension))
	return strings.Join(tmp, "")
}

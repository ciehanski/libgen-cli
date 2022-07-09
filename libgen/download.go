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
func DownloadBook(book *Book, outputPath string) error {
	var filesize int64
	filename := getBookFilename(book)

	req, err := http.NewRequest("GET", book.DownloadURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Encoding", "*")
	client := http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	r, err := client.Do(req)
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		filesize = r.ContentLength
		bar := pb.Full.Start64(filesize)

		out, err := makeFile(outputPath, filename)
		if err != nil {
			return err
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
func DownloadDbdump(filename string, outputPath string) error {
	filename = RemoveQuotes(filename)
	mirror := GetWorkingMirror(SearchMirrors)
	client := http.Client{Timeout: HTTPClientTimeout, Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	r, err := client.Get(fmt.Sprintf("%s/dbdumps/%s", mirror.String(), filename))
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusOK {
		filesize := r.ContentLength
		bar := pb.Full.Start64(filesize)

		out, err := makeFile(outputPath, filename)
		if err != nil {
			return err
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
// This is a hack that I don't like and needs to be revisited.
func GetDownloadURL(book *Book) error {
	chosenMirror := DownloadMirrors[rand.Intn(len(DownloadMirrors))]

	var x int
	tries := 3
	for tries >= x {
		switch chosenMirror.Hostname() {
		case "62.182.86.140":
			if err := getLibraryLolURL(book); err != nil {
				if err := getBooksdlDownloadURL(book); err != nil {
					return err

				}
			}
		case "libgen.rocks":
			if err := getBooksdlDownloadURL(book); err != nil {
				if err = getLibraryLolURL(book); err != nil {
					return err
				}
			}
		}
		if book.DownloadURL != "" {
			break
		}
		// Increment tries
		x++
	}

	if book.DownloadURL == "" {
		return fmt.Errorf("unable to retrieve download link for desired resource")
	}
	return nil
}

func getLibraryLolURL(book *Book) error {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   "library.lol",
		Path:   "main/",
	}
	queryURL := baseURL.String() + book.Md5
	book.PageURL = queryURL

	b, err := getBody(queryURL)
	if err != nil {
		return err
	}

	downloadURL := findMatch(libraryLolReg, b)
	if downloadURL == nil {
		return errors.New("no valid download LibraryLol download URL found")
	}

	book.DownloadURL = string(downloadURL)

	return nil
}

func getBooksdlDownloadURL(book *Book) error {
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "cdn1.booksdl.org",
		Path:   "ads.php",
	}
	q := baseURL.Query()
	q.Set("md5", book.Md5)
	baseURL.RawQuery = q.Encode()
	book.PageURL = baseURL.String()

	b, err := getBody(baseURL.String())
	if err != nil {
		return err
	}

	downloadURL := findMatch(booksdlReg, b)
	if downloadURL == nil {
		return errors.New("no valid download Booksdl download URL found")
	}
	book.DownloadURL = string(downloadURL)
	fmt.Printf("%s", book.DownloadURL)

	return nil
}

func makeFile(outputPath, filename string) (*os.File, error) {
	var out *os.File
	var mkErr error

	// Handle long titles
	if len(filename) >= 256 {
		filename = filename[:256]
	}

	// if output path was not provided
	if outputPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		if stat, err := os.Stat(fmt.Sprintf("%s/libgen", wd)); err == nil && stat.IsDir() {
			out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
		} else {
			if err := os.Mkdir(fmt.Sprintf("%s/libgen", wd), 0755); err != nil {
				return nil, err
			}
			out, mkErr = os.Create(fmt.Sprintf("%s/libgen/%s", wd, filename))
		}
		if mkErr != nil {
			return nil, mkErr
		}
		// If output path was provided
	} else {
		if stat, err := os.Stat(outputPath); err == nil && stat.IsDir() {
			out, err = os.Create(fmt.Sprintf("%s/%s", outputPath, filename))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("invalid output path")
		}
	}

	return out, nil
}

// findMatch is a helper function that searches an []byte
// for a specified regex and returns the matches.
func findMatch(reg string, response []byte) []byte {
	re := regexp.MustCompile(reg)
	match := re.FindString(string(response))

	if match != "" {
		return []byte(match)
	}

	return nil
}

func getBookFilename(book *Book) string {
	var tmp []string
	tmp = append(tmp, book.Title)
	tmp = append(tmp, fmt.Sprintf(" by %s", book.Author))
	tmp = append(tmp, fmt.Sprintf(".%s", book.Extension))
	return strings.Join(tmp, "")
}

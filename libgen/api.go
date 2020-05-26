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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

// Book is the struct of resources on Library Genesis.
type Book struct {
	ID          string
	Title       string
	Author      string
	Filesize    string
	Extension   string
	Md5         string
	Year        string
	Language    string
	Pages       string
	Publisher   string
	Edition     string
	CoverURL    string
	DownloadURL string
	PageURL     string
}

// SearchOptions are the optional parameters available for the Search
// function.
type SearchOptions struct {
	Query         string
	SearchMirror  url.URL
	Results       int
	Print         bool
	RequireAuthor bool
	Extension     string
	Year          int
	Publisher     string
}

// GetDetailsOptions are the optional parameters available for the GetDetails
// function.
type GetDetailsOptions struct {
	Hashes        []string
	SearchMirror  url.URL
	Print         bool
	RequireAuthor bool
	Extension     string
	Year          int
	Publisher     string
}

// Search sends a query to the search.php page hosted by gen.lib.rus.ec(or any
// similar mirror) and then provides the web page's contents provided from the
// resulting http request to the parseHashes() function to extract the specific
// hashes of matches found from the search query provided.
func Search(options *SearchOptions) ([]*Book, error) {
	// libgen search only allows query Results of 25, 50 or 100.
	// We handle that here
	var res int
	switch {
	case options.Results <= 25:
		res = 25
	case options.Results <= 50:
		res = 50
	default:
		res = 100
	}

	// Define DownloadURL with required query parameters
	options.SearchMirror.Path = "search.php"
	q := options.SearchMirror.Query()
	q.Set("req", options.Query)
	q.Set("lg_topic", "libgen")
	q.Set("open", "0")
	q.Set("view", "simple")
	q.Set("res", string(res))
	q.Set("phrase", "1")
	q.Set("column", "def")
	options.SearchMirror.RawQuery = q.Encode()

	b, err := getBody(options.SearchMirror.String())
	if err != nil {
		return nil, err
	}

	// Get hashes from raw webpage and store them in hashes
	hashes := parseHashes(b, options.Results)

	books, err := GetDetails(&GetDetailsOptions{
		Hashes:        hashes,
		SearchMirror:  options.SearchMirror,
		Print:         options.Print,
		RequireAuthor: options.RequireAuthor,
		Extension:     options.Extension,
		Year:          options.Year,
		Publisher:     options.Publisher,
	})
	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetDetails retrieves more details about a specific piece of media
// based off of its unique hash/id. That information is then requested
// in JSON format and sanitized in an array of Books.
func GetDetails(options *GetDetailsOptions) ([]*Book, error) {
	var books []*Book

	// For each hash found on the page, parse it into a Book struct
	for _, hash := range options.Hashes {
		options.SearchMirror.Path = "json.php"
		q := options.SearchMirror.Query()
		q.Set("ids", hash)
		q.Set("fields", JSONQuery)
		options.SearchMirror.RawQuery = q.Encode()

		b, err := getBody(options.SearchMirror.String())
		if err != nil {
			return nil, err
		}

		book, err := parseResponse(b)
		if err != nil {
			return nil, err
		}

		// Flag filters
		if options.RequireAuthor && book.Author == "" {
			continue
		}
		if options.Extension != "" && options.Extension != book.Extension {
			continue
		}
		if options.Year != 0 {
			y, err := strconv.Atoi(book.Year)
			if err != nil {
				return nil, err
			}
			if options.Year != y {
				continue
			}
		}
		if options.Publisher != "" {
			if !strings.Contains(book.Publisher, options.Publisher) {
				continue
			}
		}
		if options.Print {
			if err := printDetails(book); err != nil {
				return nil, err
			}
		}

		// Add valid book to the []Book for the search
		books = append(books, book)
	}

	return books, nil
}

// CheckMirror returns the HTTP status code of the DownloadURL provided.
func CheckMirror(url url.URL) int {
	client := http.Client{
		Timeout: HTTPClientTimeout,
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	r, err := client.Get(url.String())
	if err != nil {
		return http.StatusBadGateway
	}
	if r.StatusCode != http.StatusOK {
		return r.StatusCode
	}
	return http.StatusOK
}

// GetWorkingMirror selects a random mirror from the []url.DownloadURL
// provided and checks the mirror for a proper HTTP status code
// for working order.
func GetWorkingMirror(urls []url.URL) url.URL {
	var mirror url.URL

	for {
		randMirror := urls[rand.Intn(len(urls))]
		if CheckMirror(randMirror) == http.StatusOK {
			mirror = randMirror
			break
		}
	}

	return mirror
}

// ParseDbdumps takes in a HTTP response and scans it for
// any string that matches a filepath and returns all results.
func ParseDbdumps(response []byte) []string {
	re := regexp.MustCompile(dbdumpReg)
	dbdumps := re.FindAllString(string(response), -1)

	for i, dbdump := range dbdumps {
		dbdumps[i] = RemoveQuotes(dbdump)
	}

	return dbdumps
}

func getBody(baseURL string) ([]byte, error) {
	client := http.Client{
		Timeout: HTTPClientTimeout,
		Transport: &http.Transport{
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	r, err := client.Get(baseURL)
	if err != nil {
		log.Printf("http.Get(%q) error: %v", baseURL, err)
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to reach to mirror %v: %v", baseURL, r.StatusCode)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	return b, nil
}

// parseHashes takes in a HTTP response and scans it for
// an MD5 hash and then returns the found hashes.
func parseHashes(response []byte, results int) []string {
	var hashes []string
	re := regexp.MustCompile(SearchHref)
	matches := re.FindAllString(string(response), -1)

	var counter int
	for _, m := range matches {
		if counter >= results {
			break
		}
		re := regexp.MustCompile(SearchMD5)
		hash := re.FindString(m)
		if len(hash) == 32 {
			hashes = append(hashes, hash)
			counter++
		}
	}

	return hashes
}

// parseResponse takes in a slice of bytes and formats it
// returns a Book object from the slice of bytes.
func parseResponse(response []byte) (*Book, error) {
	var book Book
	var formattedResp []map[string]string

	if err := json.Unmarshal(response, &formattedResp); err != nil {
		return nil, err
	}
	for _, item := range formattedResp {
		for k, v := range item {
			switch k {
			case "id":
				book.ID = v
			case "title":
				book.Title = v
			case "author":
				book.Author = v
			case "filesize":
				book.Filesize = v
			case "extension":
				book.Extension = v
			case "md5":
				book.Md5 = v
			case "year":
				book.Year = v
			case "language":
				book.Language = v
			case "pages":
				book.Pages = v
			case "publisher":
				book.Publisher = v
			case "edition":
				book.Edition = v
			case "coverurl":
				book.CoverURL = v
			}
		}
	}

	return &book, nil
}

func printDetails(book *Book) error {
	var fsize string
	size, err := strconv.Atoi(book.Filesize)
	if err != nil {
		fsize = "N/A"
	} else {
		fsize = humanize.Bytes(uint64(size))
	}

	// Print separation lines
	fmt.Println(strings.Repeat("-", 80))

	// Print ID + Title
	fTitle := fmt.Sprintf("%5s %s", color.New(color.FgHiBlue).Sprintf(book.ID), book.Title)
	fTitle = formatTitle(fTitle, TitleMaxLength)
	if runtime.GOOS == "windows" {
		_, err = fmt.Fprintf(color.Output, "%s\n    ++ ", fTitle)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("%s\n    ++ ", fTitle)
	}

	// Slice author name if it exceeds AuthorMaxLength
	var formatAuthor string
	if len(book.Author) > AuthorMaxLength {
		formatAuthor = book.Author[:AuthorMaxLength]
	} else if book.Author == "" {
		formatAuthor = "N/A"
	} else {
		formatAuthor = book.Author
	}

	err = prettify("author", formatAuthor, color.FgYellow, "-25")
	if err != nil {
		return err
	}
	err = prettify("year", book.Year, color.FgCyan, "4")
	if err != nil {
		return err
	}
	err = prettify("size", fsize, color.FgGreen, "6")
	if err != nil {
		return err
	}
	err = prettify("type", book.Extension, color.FgRed, "4")
	if err != nil {
		return err
	}
	fmt.Println()

	return nil
}

// formatTitle shortens the title of a Book down to
// the maximum allowed by TitleMaxLength.
func formatTitle(title string, maximumLength int) string {
	var fTitle []string
	var counter int

	if len(title) <= maximumLength {
		return title
	}

	title = strings.TrimSpace(title)
	for _, t := range strings.Split(title, " ") {
		counter += len(t)

		if counter > maximumLength {
			counter = 0
			t = t + "...\n"
		}
		fTitle = append(fTitle, t)
	}

	return strings.Join(fTitle, " ")
}

// prettify is a helper function that adds color and
// formats text returned to the user.
func prettify(key string, value string, col color.Attribute, align string) error {
	c := color.New(col).SprintFunc()
	a := fmt.Sprintf("%%%ss ", align)
	s := fmt.Sprintf("@%s "+a, c(key), value)
	if runtime.GOOS == "windows" {
		_, err := fmt.Fprintf(color.Output, a, s)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf(a, s)
	}
	return nil
}

// RemoveQuotes is a helper function that removes the quotes from
// dbdumps page results.
func RemoveQuotes(s string) string {
	if s == "" {
		return ""
	}
	s = s[1:]
	s = s[:len(s)-1]
	return s
}

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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

// Search sends a query to the search.php page hosted by gen.lib.rus.ec(or any
// similar mirror) and then provides the web page's contents provided from the
// resulting http request to the parseHashes() function to extract the specific
// hashes of matches found from the search query provided.
func Search(query string, results int, print bool, requireAuthor bool, extension string) ([]Book, error) {
	searchMirror := GetWorkingMirror(SearchMirrors)
	if searchMirror.Host == "" {
		return nil, errors.New("unable to reach any Library Genesis resources")
	}

	// libgen search only allows query results of 25, 50 or 100.
	// We handle that here
	var res int
	switch {
	case results <= 25:
		res = 25
	case results <= 50:
		res = 50
	default:
		res = 100
	}

	// Define URL with required query parameters
	searchMirror.Path = "search.php"
	q := searchMirror.Query()
	q.Set("req", query)
	q.Set("lg_topic", "libgen")
	q.Set("open", "0")
	q.Set("view", "simple")
	q.Set("res", string(res))
	q.Set("phrase", "1")
	q.Set("column", "def")
	searchMirror.RawQuery = q.Encode()

	// Execute GET request on search query
	r, err := http.Get(searchMirror.String())
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, err
	}

	// Read body of response to get HTML
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Close request and handle possible error.
	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	// Get hashes from raw webpage and store them in hashes
	hashes := parseHashes(string(b), results)

	books, err := GetDetails(hashes, searchMirror, print, requireAuthor, extension)
	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetDetails retrieves more details about a specific piece of media
// based off of its unique hash/id. That information is then requested
// in JSON format and sanitized in an array of Books.
func GetDetails(hashes []string, searchMirror url.URL, print bool, requireAuthor bool, extension string) ([]Book, error) {
	var (
		books        []Book
		formatAuthor string
		fsize        string
	)

	for _, hash := range hashes {
		searchMirror.Path = "json.php"
		q := searchMirror.Query()
		q.Set("ids", hash)
		q.Set("fields", JSONQuery)
		searchMirror.RawQuery = q.Encode()

		r, err := http.Get(searchMirror.String())
		if err != nil {
			log.Printf("error reaching API: %v", err)
			return nil, err
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading response from API: %v", err)
			return nil, err
		}

		book, err := parseResponse(b)
		if err != nil {
			return nil, err
		}

		// Flag filters
		if requireAuthor && book.Author == "" {
			continue
		}
		if extension != "" && extension != book.Extension {
			continue
		}

		if print {
			size, err := strconv.Atoi(book.Filesize)
			if err != nil {
				fsize = "N/A"
			} else {
				fsize = humanize.Bytes(uint64(size))
			}

			fmt.Println(strings.Repeat("-", 80))
			fTitle := fmt.Sprintf("%5s %s", color.New(color.FgHiBlue).Sprintf(book.ID), book.Title)
			fTitle = formatTitle(fTitle)
			fmt.Printf("%s\n    ++ ", fTitle)

			// Slice author name if it exceeds AuthorMaxLength
			if len(book.Author) > AuthorMaxLength {
				formatAuthor = book.Author[:AuthorMaxLength]
			} else {
				formatAuthor = book.Author
			}

			prettify("author", formatAuthor, color.FgYellow, "-25")
			prettify("year", book.Year, color.FgCyan, "4")
			prettify("size", fsize, color.FgGreen, "6")
			prettify("type", book.Extension, color.FgRed, "4")
			fmt.Println()
		}

		books = append(books, book)

		if err := r.Body.Close(); err != nil {
			return nil, err
		}
	}

	return books, nil
}

// CheckMirror returns the HTTP status code of the URL provided.
func CheckMirror(url url.URL) int {
	r, err := http.Get(url.String())
	if err != nil || r.StatusCode != http.StatusOK {
		return http.StatusBadGateway
	}

	return http.StatusOK
}

func GetWorkingMirror(urls []url.URL) url.URL {
	var mirror url.URL
	rand.Seed(time.Now().UnixNano())

	for {
		randMirror := urls[rand.Intn(len(urls))]

		if CheckMirror(randMirror) == http.StatusOK {
			mirror = randMirror
			break
		} else {
			continue
		}
	}

	return mirror
}

func parseHashes(response string, results int) []string {
	var hashes []string
	re := regexp.MustCompile(SearchHref)
	matches := re.FindAllString(response, -1)

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
func parseResponse(data []byte) (Book, error) {
	var book Book
	var formattedResp []map[string]string

	if err := json.Unmarshal(data, &formattedResp); err == nil {
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
				}
			}
		}
	} else {
		return Book{}, err
	}

	return book, nil
}

// formatTitle shortens the title of a Book down to
// the maximum allowed by TitleMaxLength.
func formatTitle(title string) string {
	var fTitle []string
	var counter int

	if len(title) < TitleMaxLength {
		return title
	}

	title = strings.TrimSpace(title)
	for _, t := range strings.Split(title, " ") {
		counter += len(t)

		if counter > TitleMaxLength {
			counter = 0
			t = t + "\n"
		}
		fTitle = append(fTitle, t)
	}

	return strings.Join(fTitle, " ")
}

func prettify(key string, value string, col color.Attribute, align string) {
	c := color.New(col).SprintFunc()
	a := fmt.Sprintf("%%%ss ", align)
	s := fmt.Sprintf("@%s "+a, c(key), value)
	fmt.Printf(a, s)
}

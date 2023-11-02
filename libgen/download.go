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
	"context"
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
	iface "github.com/ipfs/boxo/coreiface"
	ipath "github.com/ipfs/boxo/coreiface/path"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/repo/fsrepo"
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

// Heavily borrowed from https://github.com/ipfs/ipget as
// documentation on creating an IPFS client via kubo to 
// download is lacking
func DownloadBookIPFS(book *Book, outputPath string) error {
	filename := getBookFilename(book)
    ctx := context.Context(context.Background())

    // Create IPFS client
    ipfs, err := getIPFSclient(ctx)
    if err != nil {
        return err
    }

    // Parse IPFS URL
    ipfsPath, err := parseIPFSURL(book.DownloadURL)
    if err != nil {
        return err
    }

    ipfsNode, err := ipfs.Unixfs().Get(ctx, ipfsPath)
    if err != nil {
        return err
    }
    out, err := makeFile(outputPath, filename)
    if err != nil {
    	return err
    }

    var r io.Reader = ipfsNode.(io.Reader) 
    _, err = io.Copy(out, r)
    if err != nil {
    	return err
    }
    
    if err := out.Close(); err != nil {
    	return err
    }
    if err := ipfsNode.Close(); err != nil {
    	return err
    }

	return nil
}

func getIPFSclient(ctx context.Context) (iface.CoreAPI, error) {
    // Create a new IPFS client.
    defaultPath, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	// Open the repo
	repo, err := fsrepo.Open(defaultPath)
	if err != nil {
		return nil, err
	}

	// Construct the IPFS node
	node, err := core.NewNode(ctx, &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTClientOption,
		Repo:    repo,
	})
	if err != nil {
		return nil, err
	}

    // IPFS API client
    ipfs, err := coreapi.NewCoreAPI(node)
	if err == nil {
		return ipfs, nil
	}

    return nil, err
}

func parseIPFSURL(path string) (ipath.Path, error) {
    url := findMatch(ipfsReg, []byte(path))

	ipfsPath := ipath.New(string(url))
	if ipfsPath.IsValid() == nil {
		return ipfsPath, nil
	}

    return nil, ipfsPath.IsValid() 
}

// GetDownloadURL picks a random download mirror to download the specified
// resource from.
func GetDownloadURL(book *Book, useIpfs bool) error {
	chosenMirror := DownloadMirrors[rand.Intn(len(DownloadMirrors))]

	var x int
	tries := 3
	for tries >= x {
		switch chosenMirror.Hostname() {
		case "library.lol":
            if useIpfs {
			    if err := getLibraryLolURL(book, true); err != nil {
                    return err
                }
            } else {
			    if err := getLibraryLolURL(book, false); err != nil {
			       	if err := getLibgenPMURL(book); err != nil {
			       		return err
			       	}
                }
            }
		case "libgen.pm":
            if !useIpfs {
			    if err := getLibgenPMURL(book); err != nil {
			        if err := getLibraryLolURL(book, false); err != nil {
			       		return err
			       	}
                }
            } else {
                // No IPFS URLs on libgen.pm pages, fallback to 
                // library.lol
			    if err := getLibraryLolURL(book, true); err != nil {
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

func getLibraryLolURL(book *Book, useIpfs bool) error {
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "library.lol",
		Path:   "main/",
	}
	queryURL := baseURL.String() + book.Md5
	book.PageURL = queryURL

	b, err := getBody(queryURL)
	if err != nil {
		return err
	}

    downloadURL := []byte{}
    if useIpfs {
	    downloadURL = findMatch(libraryLolIPFSReg, b)
	    if downloadURL == nil {
	    	return errors.New("no valid download LibraryLol download URL found")
	    }
    } else {
	    downloadURL = findMatch(libraryLolReg, b)
	    if downloadURL == nil {
	    	return errors.New("no valid download LibraryLol download URL found")
	    }
    }

	book.DownloadURL = string(downloadURL)

	return nil
}

func getLibgenPMURL(book *Book) error {
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "libgen.pm",
        Path:   "ads",
	}
	queryURL := baseURL.String() + book.Md5
	book.PageURL = queryURL

	b, err := getBody(queryURL)
	if err != nil {
		return err
	}

	downloadURL := findMatch(libgenPMReg, b)
	if downloadURL == nil {
		return errors.New("no valid LibgenPM download URL found")
	}
	book.DownloadURL = fmt.Sprintf("https://libgen.rocks/%s", string(downloadURL))

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

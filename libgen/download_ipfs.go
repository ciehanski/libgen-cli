// Some of the code within download_ipfs.go to create an interact with IPFS was
// inspired by https://github.com/ipfs/ipget as documentation on creating an IPFS
// client via kubo to download is lacking. I've listed ipget's MIT license below.
//
// The MIT License (MIT)
//
// Copyright (c) 2015 Stephen Whitmore
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package libgen

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	iface "github.com/ipfs/boxo/coreiface"
	"github.com/ipfs/boxo/coreiface/options"
	ipath "github.com/ipfs/boxo/coreiface/path"
	ifiles "github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo"
	"github.com/ipfs/kubo/repo/fsrepo"
)

func DownloadBookIPFS(book *Book, outputPath string) error {
	filename := getBookFilename(book)
	ctx := context.Context(context.Background())

	// Create temp IPFS dir
	ipfsDir, err := os.MkdirTemp("", "libgen-cli-ipfs")
	if err != nil {
		return err
	}
	defer os.RemoveAll(ipfsDir)

	// Create IPFS client
	ipfs, err := setupIPFSclient(ctx, ipfsDir)
	if err != nil {
		return err
	}
	// Parse IPFS URL
	ipfsPath, err := parseIPFSurl(book.DownloadURL)
	if err != nil {
		return err
	}

	// Get IPFS node from URL
	ipfsNode, err := ipfs.Unixfs().Get(ctx, ipfsPath)
	if err != nil {
		return err
	}
	defer ipfsNode.Close()

	// Create progress bar
	nodeSize, err := ipfsNode.Size()
	if err != nil {
		return err
	}
	bar := pb.New64(nodeSize).Start()

	var outPath string
	if outputPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		if _, err := os.Stat(fmt.Sprintf("%s/libgen", wd)); err != nil {
			if err := os.Mkdir(fmt.Sprintf("%s/libgen", wd), 0755); err != nil {
				return err
			}
		}
		outPath = fmt.Sprintf(fmt.Sprintf("%s/libgen/%s", wd, filename))
	} else {
		outPath = filepath.Join(outputPath, filename)
	}

	// Copy IPFS node to output file
	if err := makeIPFSfile(ipfsNode, outPath, bar); err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func makeIPFSfile(ipfsNode ifiles.Node, fpath string, bar *pb.ProgressBar) error {
	switch nd := ipfsNode.(type) {
	case *ifiles.Symlink:
		return os.Symlink(nd.Target, fpath)
	case ifiles.File:
		f, err := os.Create(fpath)
		defer f.Close()
		if err != nil {
			return err
		}

		var r io.Reader = nd
		_, err = io.Copy(f, bar.NewProxyReader(r))
		if err != nil {
			return err
		}
		return nil
	case ifiles.Directory:
		err := os.Mkdir(fpath, 0755)
		if err != nil {
			return err
		}

		entries := nd.Entries()
		for entries.Next() {
			child := filepath.Join(fpath, entries.Name())
			if err := makeIPFSfile(entries.Node(), child, bar); err != nil {
				return err
			}
		}
		return entries.Err()
	default:
		return fmt.Errorf("file type %T at %q is not supported", nd, fpath)
	}
}

func setupIPFSclient(ctx context.Context, ipfsDir string) (iface.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	if err = setupIPFSplugins(defaultPath); err != nil {
		return nil, err
	}

	// Setup IPFS identity
	icfg, err := setupIPFSident()
	if err != nil {
		return nil, err
	}

	// Init IPFS repo
	localRepo, err := initIPFSrepo(ipfsDir, icfg)

	// Construct the temp IPFS node
	node, err := core.NewNode(ctx, &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTClientOption,
		Repo:    localRepo,
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

func initIPFSrepo(ipfsDir string, icfg *config.Config) (repo.Repo, error) {
	// Init IPFS local repo
	err := fsrepo.Init(ipfsDir, icfg)
	if err != nil {
		return nil, err
	}

	// Open the repo
	localRepo, err := fsrepo.Open(ipfsDir)
	if err != nil {
		return nil, err
	}

	return localRepo, nil
}

func setupIPFSident() (*config.Config, error) {
	identity, err := config.CreateIdentity(io.Discard, []options.KeyGenerateOption{
		options.Key.Type(options.Ed25519Key),
	})
	if err != nil {
		return nil, err
	}

	icfg, err := config.InitWithIdentity(identity)
	if err != nil {
		return nil, err
	}

	// Configure the temporary node
	icfg.Routing.Type = config.NewOptionalString("dhtclient")
	icfg.Datastore.NoSync = true

	return icfg, nil
}

func setupIPFSplugins(path string) error {
	// Load plugins. This will skip the repo if not available.
	plugins, err := loader.NewPluginLoader(filepath.Join(path, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func parseIPFSurl(path string) (ipath.Path, error) {
	url := findMatch(ipfsReg, []byte(path))

	ipfsPath := ipath.New(string(url))
	if ipfsPath.IsValid() == nil {
		return ipfsPath, nil
	}

	return nil, ipfsPath.IsValid()
}

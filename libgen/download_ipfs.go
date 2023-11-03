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
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo"
	"github.com/ipfs/kubo/repo/fsrepo"
)

// Heavily borrowed from https://github.com/ipfs/ipget as
// documentation on creating an IPFS client via kubo to
// download is lacking
func DownloadBookIPFS(book *Book, outputPath string) error {
	filename := getBookFilename(book)
    ctx := context.Context(context.Background())

    // Create IPFS client
    ipfs, err := setupIPFSclient(ctx)
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

    // Copy IPFS node to output file
    if err := makeIPFSfile(ipfsNode, filepath.Join(outputPath, filename), bar); err != nil {
        return err
    }

	return nil
}

func makeIPFSfile(ipfsNode files.Node, fpath string, bar *pb.ProgressBar) error {
    switch nd := ipfsNode.(type) {
	case *files.Symlink:
		return os.Symlink(nd.Target, fpath)
	case files.File:
		f, err := os.Create(fpath)
		defer f.Close()
		if err != nil {
			return err
		}

		var r io.Reader = nd
		if bar != nil {
			r = bar.NewProxyReader(r)
		}
		_, err = io.Copy(f, r)
		if err != nil {
			return err
		}
		return nil
	case files.Directory:
		err := os.Mkdir(fpath, 0777)
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

func setupIPFSclient(ctx context.Context) (iface.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	if err = setupIPFSplugins(defaultPath); err != nil {
		return nil, err
	}

    // Create temp IPFS dir
	ipfsDir, err := os.MkdirTemp("", "libgen-cli-ipfs")
	if err != nil {
		return nil, err
	}
    defer os.RemoveAll(ipfsDir)

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

package libgen

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type LibgenConf struct {
	GeneralConf GeneralConf `toml:"general"`
}

type GeneralConf struct {
	DownloadDir string `toml:"download_dir"`
}

func ReadConf() (*LibgenConf, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	confPath := confDir + "/libgen/libgen.toml"
	conf := &LibgenConf{
		GeneralConf: GeneralConf{
			DownloadDir: homeDir + "/Downloads/libgen",
		},
	}

	bytes, err := os.ReadFile(confPath)
	if err != nil {
		bytes, err = writeDefaultConf(confDir + "/libgen", confPath, conf)
		if err != nil {
			return nil, err
		}
	}

	err = toml.Unmarshal([]byte(bytes), &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func writeDefaultConf(confDir, confPath string, conf *LibgenConf) ([]byte, error) {
	err := os.MkdirAll(confDir, 0755)
	if err != nil {
		return nil, err
	}

	confBytes, err := toml.Marshal(conf)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(confPath, confBytes, 0644)
	if err != nil {
		return nil, err
	}

	return confBytes, nil
}

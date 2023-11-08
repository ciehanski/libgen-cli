## libgen-cli [![Build & Test](https://github.com/ciehanski/libgen-cli/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/ciehanski/libgen-cli/actions/workflows/build.yml) [![Coverage Status](https://coveralls.io/repos/github/ciehanski/libgen-cli/badge.svg?branch=master)](https://coveralls.io/github/ciehanski/libgen-cli?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/ciehanski/libgen-cli)](https://goreportcard.com/report/github.com/ciehanski/libgen-cli)

libgen-cli is a command line interface application which allows users to
quickly query the Library Genesis dataset and download any of its contents.

## Table of Contents
- [Installation](#installation)
- [Commands](#commands)
	- [Search](#search)
	- [Download](#download)
	- [Dbdumps](#dbdumps)
	- [Status](#status)
    - [Version](#version)
    - [Link](#link)
- [Disclaimer](#disclaimer)
- [License](#license)

![libgen-cli Example](https://github.com/ciehanski/libgen-cli/blob/master/resources/libgen-cli-example.gif)

## Installation

You can download the latest binary from the releases section of this repo
which can be found [here](https://github.com/ciehanski/libgen-cli/releases).

If you have [Golang](https://golang.org) installed on your local machine you can use the
commands belows to install it directly into your $GOPATH.

```bash
$ go install github.com/ciehanski/libgen-cli@latest
```

## Commands

### Search:

The _search_ command is the bread and butter of libgen-cli. Simply provide an
additional argument to have libgen-cli scrape the Library Genesis dataset and
provide you results available for download. See below for a few examples:

```bash
$ libgen search kubernetes
```

Force download of query results via available IPFS mirrors:

```bash
$ libgen search kubernetes -i
```

Filter the amount of results displayed:  
(Must be between 1-100).

```bash
$ libgen search kubernetes -r 5
```

Filter by file extension(s):

```bash
$ libgen search kubernetes -e pdf
```

```bash
$ libgen search kubernetes -e "pdf,epub"
```

Specify an output path:

```bash
$ libgen search kubernetes -o ~/Desktop/libgen
```

Sort the results by (id, title, author, pub, year, lang, size, ext):

```bash
$ libgen search kubernetes -s title --sort-asc=false
```

```bash
$ libgen search kubernetes --sort-by size
```

Require that the author field is listed and available for the specific search
results:
 
```bash
$ libgen search kubernetes -a
```

Filter results by year:

```bash
$ libgen search kubernetes -y 2019
```

Filter by the publisher's name:

```bash
$ libgen search kubernetes -p "Michael Joseph"
```

Filter by the file's language:

```bash
$ libgen search kubernetes -l "english"
```


### Download:

The _download_ command will allow you to download a specific book if already 
know the MD5 hash. See below for an example:

```bash
$ libgen download 2F2DBA2A621B693BB95601C16ED680F8
```

Force download of query results via available IPFS mirrors:

```bash
$ libgen download 2F2DBA2A621B693BB95601C16ED680F8 --ipfs-mirrors
```

You can bulk a list of MD5s by passing it as a command line argument: 

```bash
$ libgen download 6B4B4F0073B92248EFAB34F100CA20D4 FAA323B98939EE385BB33A1A3B88AFCA
```

Download a text list of MD5s at once: 

```bash
cat list.txt | xargs libgen download
```

Specify an output path:

```bash
$ libgen download -o ~/Desktop/ 2F2DBA2A621B693BB95601C16ED680F8
```

The _download-all_ command will allow you to download all query results. This
command uses the same flags and arguments as the _search_. See below for an example:

```bash
$ libgen download-all kubernetes
```

Specify the desired amount of results downloaded:  
(Must be between 1-100).

```bash
$ libgen download-all kubernetes -r 50
```

Specify an output path:

```bash
$ libgen download-all -o ~/Desktop/ kubernetes
```

Force download of all query results via available IPFS mirrors:

```bash
$ libgen download-all -o ~/Desktop/ kubernetes -i
```

Download all of the sorted results by (id, title, author, pub, year, lang, size, ext):

```bash
$ libgen download-all kubernetes -s year -r 100
```

```bash
$ libgen download-all kubernetes --sort-by lang --sort-asc -r 70
```


### Dbdumps:

The _dbdumps_ command will list out all of the compiled database dumps of
libgen's database and allow you to download them with ease.

```bash
$ libgen dbdumps
```

Specify an output path:

```bash
$ libgen dbdumps -o ~/Desktop
```


### Link

The _link_ command will retrieve and output the direct download link
of a specific MD5 resource.

```bash
$ libgen link 2F2DBA2A621B693BB95601C16ED680F8
```

Retrieve the available IPFS link of a specific MD5 resource:

```bash
$ libgen link 2F2DBA2A621B693BB95601C16ED680F8 -i
```


### Status:

The _status_ command simply pings the mirrors for Library Genesis and
returns the status [OK] or [FAIL] depending on if the mirror is responsive 
or not. See below for an example:

```bash
$ libgen status
```

Specify to only check the status of the download mirrors:

```bash
$ libgen status -m download
```

Specify to only check the status of the search mirrors:

```bash
$ libgen status -m search
```


### Version:

Check the version of the installed libgen-cli client:

```bash
$ libgen -v
```

## Disclaimer

This repository is for research purposes only, the use of this code is your sole responsibility.

I take NO responsibility and/or liability for how you choose to use any of the source code available 
here. By using any of the files available in this repository, you understand that you are AGREEING 
TO USE AT YOUR OWN RISK. Once again, ALL files available here are for EDUCATION and/or RESEARCH purposes ONLY.

## License
- Apache License 2.0

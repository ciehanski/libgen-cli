## libgen-cli

libgen-cli is a command line interface application which allows users to
quickly query the Library Genesis dataset and download any of its contents.


![Example](https://github.com/ciehanski/libgen-cli/blob/master/resources/libgen-cli-example.gif)

## Installation

You can download the latest binary from the releases section of this repo
which can be found [here](https://github.com/ciehanski/libgen-cli/releases).

If you have Golang installed on your local machine you can use the
commands belows to install it directly into your $GOPATH.

```
go get -u github.com/ciehanski/libgen-cli
go install github.com/ciehanski/libgen-cli
```

## Commands

#### Search:

The _search_ command is the bread and butter of libgen-cli. Simply provide an
additional argument to have libgen-cli scrap the Library Genesis dataset and
provide you results available for download. An option to control how many
results are returned is currently in progress. See below for an example:

```bash
libgen-cli search kubernetes
```

#### Download:

The _download_ command will allow you to download a specific book if already 
know the hash.

```bash
libgen-cli download 2F2DBA2A621B693BB95601C16ED680F8
```

#### Status:

The _status_ command simply pings the mirrors for Library Genesis and
returns the status [OK] or [FAIL] depending on if the mirror is responsive 
or not. See below for an example:

```bash
libgen-cli status
```

## License
- MIT
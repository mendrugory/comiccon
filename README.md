# Comiccon


[![Build Status](https://travis-ci.org/mendrugory/comiccon.svg?branch=master)](https://travis-ci.org/mendrugory/comiccon)

Comiccon is CLI toy project to download and keep updated comics from [here](https://the-eye.eu/public/Comics/).

## Installation

```bash
go get -u github.com/mendrugory/comiccon
```

## How to use it

### Help
```
$ comiccon help
```

### Download
```bash
$ comiccon download
```

If a comic is already downloaded, it will not be downloaded it again.

### Optional flags

* Base Folder: Directory where the comics will be downloaded (default: current directory)
* Extensions: Extensions of the files which will be downloaded (default: cbr, jpg and pdf)
* Link: Sub link of the route if you only want to download a part of the collection (check the [list of collections](https://the-eye.eu/public/Comics/))

```bash
$ comiccon download --basefolder /tmp/comics --extensions cbr --link "DC Chronology"
```
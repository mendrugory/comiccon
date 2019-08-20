# Comiccon

Comiccon is CLI toy project to download and keep updated comics from [here](https://the-eye.eu/public/Comics/).

#### Build

[![Build Status](https://travis-ci.org/mendrugory/comiccon.svg?branch=master)](https://travis-ci.org/mendrugory/comiccon)

#### Docker Image
[![](https://images.microbadger.com/badges/version/mendrugory/comiccon.svg)](https://microbadger.com/images/mendrugory/comiccon) [![](https://images.microbadger.com/badges/image/mendrugory/comiccon.svg)](https://microbadger.com/images/mendrugory/comiccon)


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

The configuration of the command is saved in a file called `config.json`.

### Optional flags

* Base Folder: Directory where the comics will be downloaded. It is created if it does not exist (default: current directory)
* Extensions: Extensions of the files which will be downloaded (default: cbr, jpg and pdf)
* Link: Sub link of the route if you only want to download a part of the collection (check the [list of collections](https://the-eye.eu/public/Comics/))

```bash
$ comiccon download --basefolder /tmp/comics --extensions cbr --link "DC Chronology"
```

## Docker

To run it using Docker, the only thing that we must have in mind is mapping the volumes in order to keep the comics.

```bash
$ docker run --rm -v /tmp/comics:/tmp/comics mendrugory/comiccon download --basefolder /tmp/comics --extensions cbr --link "DC Chronology"
```
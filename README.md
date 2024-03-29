# ID3go [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Command line tool for updating of media file tags.

### Table Of Content
<!-- TOC -->

- [Table Of Content](#table-of-content)
- [Features](#features)
- [Dependencies](#dependencies)
  - [Alpine](#alpine)
  - [MacOS](#macos)
  - [Ubuntu](#ubuntu)
- [Install](#install)
- [Usage](#usage)
  - [Get](#get)
  - [Set](#set)
- [Repository](#repository)
- [Contributors](#contributors)
- [Changelog](#changelog)
- [License](#license)

<!-- /TOC -->
<!--more-->
`id3go` use [go-taglib](https://github.com/wtolson/go-taglib) for easy viewing and updating of media file tags.

### Features

- Minimalistic output for easy batch processing
- Support file list, wildcard filename(shell globing)
- Support unicode through TagLib
- Display and update following tags
  - Album
  - Artist
  - Comments
  - Title
  - Track
  - Year

### Dependencies

#### Alpine

```sh
apk add taglib-dev libc-dev
```

#### MacOS

```sh
brew install taglib
```

#### Ubuntu

```sh
apt-get install libtagc0-dev
```

### Install

```sh
go get github.com/J-Siu/id3go
cd $GOPATH/src/github.com/J-Siu/id3go
go install
```

### Usage

#### Get

Display tags of files.

```sh
id3go get <files>
```

Examples:

```sh
id3go get media.mp3
id3go get *.mp3
```

#### Set

Set tags of files.

```sh
id3go set [flags] <files>
```

Flags:

short|long|usage
---|---|---
-S|--Save|save to file. Without this flag, `set` will not writing to files (dry run).
-A|--album string|set album
-a|--artist string|set artist
-c|--comment string|set comments
-h|--help|help for set
-t|--title string|set title
-T|--track string|set track
-y|--year string|set year

Examples:

```sh
# Set artist="me", title="A song title", in dry run mode (default)
id3go set -a me -t "A song title" song.mp3

# Set artist="me", title="A song title", and saving (-S) to file
id3go set -a me -t "A song title" -S song.mp3

# Set album="My Record" to all mp3 in current dir, dry run only
id3go set --album "My Record" *.mp3
```

### Repository

- [id3go](https://github.com/J-Siu/id3go)

### Contributors

- [John Sing Dao Siu](https://github.com/J-Siu)

### Changelog

- 1.0
  - Initial release
- 1.1
  - Add GitHub workflow
  - Use Go module
- 1.2
  - Use Go 1.16
- 1.3
  - Remove GitHub workflow due to taglib dependency
- v1.3.3
  - Use Go 1.18
  - Update packages
- v1.3.4
  - Update to Go 1.20 and dependency

### License

The MIT License

Copyright (c) 2023 John Siu

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
# gen-version-file

Tool to generate a version file. An humble helper to add to your release flow.

# Install

```sh
mkdir -p $GOPATH/github.com/mh-cbon
cd $GOPATH/github.com/mh-cbon
git clone https://github.com/mh-cbon/gen-version-file.git
cd gen-version-file
glide install
go install
```

# Usage

```sh
NAME:
   gen-version-file - Generate version file

USAGE:
   gen-version-file --ver=0.0.1 --lang=go

VERSION:
   0.0.1

COMMANDS:
GLOBAL OPTIONS:
   --ver value			Version number to write
   --lang value, -l value	Language to generate for. (default: "go")
   --help, -h			show help
   --version, -v		print the version
```

Using `gen-version-file --ver=2.0.1 --lang=go`, a new file is generated at `$CWD/GenVersionFile/index.go` with this content

```go
package GenVersionFile

func Version () string {
  return "2.0.1"
}

```

File that you can use in your program like this

```go
package main

import (
  "fmt"

  "your/go/program/GenVersionFile"
)

func main () {
  fmt.Println(GenVersionFile.Version())
}
```

# Shell

[![GoDoc](https://godoc.org/github.com/matthewmueller/shell?status.svg)](https://godoc.org/github.com/matthewmueller/shell)

Simple package for running shell scripts with Go. Pairs well with [gorun](https://github.com/erning/gorun).

## Example

```go
/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

import (
	"fmt"

	"github.com/matthewmueller/shell"
)

var sh = shell.Runf

func main() {
	hi := sh("echo %q", "hi")
	fmt.Printf(hi)
}
```

## Install

```
go get -u github.com/matthewmueller/shell
```

## License

MIT

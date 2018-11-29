## is-terminal

Returns true or false depending on if the `os.Stdout` is associated with a terminal.

Extracted from [Sirupsen/logrus](https://github.com/Sirupsen/logrus)

## Install

```bash
$ go get github.com/azer/is-terminal
```

## Usage

```go
package main

import (
  "github.com/azer/is-terminal"
  "syscall"
  "fmt"
)

func () {
  fmt.Println(isterminal.IsTerminal(syscall.Stdout))
  // => true or false
}
```

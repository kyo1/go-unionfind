# go-unionfind

## Installation

```sh
go get -u github.com/kyo1/go-unionfind
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/kyo1/go-unionfind"
)

func main() {
	// Initialize
	uf := unionfind.New(100)

    // Union tree
	uf.Union(0, 1)
	uf.Union(1, 2)

	if uf.Same(0, 1) {
		fmt.Println(uf.Root(0)) // 0
	}

	fmt.Println(uf.Size(2)) // 3
}
```

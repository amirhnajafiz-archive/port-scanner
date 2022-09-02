<p align="center">
    <img src="assets/logo.png" width="200" />
</p>

<h1 align="center">
Port Scanner
</h1>

<p align="center">
Fast library for scanning ports of a host
</p>

## How to use?
Get package:
```shell
go get github.com/amirhnajafiz/port-scanner
```

You need to give a host address and port range, after that the scanner will 
check the ports and returns the open ports.

```go
package main

import (
    "fmt"

    scanner "github.com/amirhnajafiz/port-scanner"
)

func main() {
    // creating scanner with URI, Port Range, Number of workers (optional, default is 100)
    openports := scanner.Scan("", 2048)

    for _, port := range openports {
        fmt.Printf("%d open\n", port)
    }
}
```

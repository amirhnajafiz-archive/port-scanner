package main

import (
	"fmt"

	scanner "github.com/amirhnajafiz/port-scanner"
)

func main() {
	openports := scanner.Scan("", 2048, 100)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

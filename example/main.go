package main

import (
	"fmt"

	scanner "github.com/amirhnajafiz/port-scanner"
)

func main() {
	{
		// scanning ports 0 to 2048
		openPorts := scanner.Scan("", []int{})

		for _, port := range openPorts {
			fmt.Printf("%d open\n", port)
		}
	}
	{
		// scanning ports 100 to 5000
		openPorts := scanner.Scan("", []int{100, 5000})

		for _, port := range openPorts {
			fmt.Printf("%d open\n", port)
		}
	}
	{
		// scanning ports 0 to 5000
		openPorts := scanner.Scan("", []int{5000})

		for _, port := range openPorts {
			fmt.Printf("%d open\n", port)
		}
	}
	{
		// scanning ports 0 to 5000 with 45 workers
		openPorts := scanner.Scan("", []int{5000}, 45)

		for _, port := range openPorts {
			fmt.Printf("%d open\n", port)
		}
	}
}

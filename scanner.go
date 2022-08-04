package portscanner

import (
	"sort"

	"github.com/amirhnajafiz/port-scanner/worker"
)

// Scan takes a uri, portRange and number of workers and will
// test the uri ports in portRange (0, portRange) with workers and then
// returns the open ports.
func Scan(uri string, portRange int, numberOfWorkers int) []int {
	ports := make(chan int, numberOfWorkers)
	response := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker.Worker(ports, response, uri)
	}

	go func() {
		for i := 0; i <= portRange; i++ {
			ports <- i
		}
	}()

	for i := 0; i <= portRange; i++ {
		port := <-response
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(response)

	sort.Ints(openports)

	return openports
}

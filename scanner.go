package portscanner

import (
	"sort"

	"github.com/amirhnajafiz/port-scanner/worker"
)

// Scan takes an uri, portRange and number of workers and will
// test the uri ports in portRange (0, portRange) with workers and then
// returns the open ports.
func Scan(uri string, portRange int, workers ...int) []int {
	numberOfWorkers := 100
	if len(workers) > 0 {
		numberOfWorkers = workers[0]
	}

	ports := make(chan int, numberOfWorkers)
	response := make(chan int)

	var openPorts []int

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
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(response)

	sort.Ints(openPorts)

	return openPorts
}

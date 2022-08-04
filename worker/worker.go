package worker

import (
	"fmt"
	"net"
)

// Worker makes a connection to the give uri and a port
// then it returns the result.
func Worker(port chan int, response chan int, uri string) {
	for p := range port {
		address := fmt.Sprintf("%s:%d", uri, p)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			response <- 0
			continue
		}

		conn.Close()
		response <- p
	}
}

package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/ayoubc/port_scanner/color"
)

var (
	// HOST = "scanme.nmap.org"
	TCP       = "tcp"
	HOST      = "127.0.0.1"
	PORT      = 80
	ChunkSize = 1000
	LIMIT     = (1 << 16)
)

var wg sync.WaitGroup

func checkPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	conn.Close()
	color.Cprint(fmt.Sprintf("Port %d open", port), "green")
}

func scanPorts(host string) {
	for i := 1; i < LIMIT; i += ChunkSize {
		groupeSize := ChunkSize
		if i+ChunkSize > LIMIT {
			groupeSize = LIMIT - i
		}

		wg.Add(groupeSize)
		for j := i; j < i+groupeSize; j++ {
			go checkPort(host, j, &wg)
		}
		wg.Wait()
	}
}

func parseArgs() string {
	host := flag.String("host", HOST, "the host to scan")
	// // Parse
	flag.Parse()

	return *host
}

func main() {
	start := time.Now()
	host := parseArgs()
	scanPorts(host)
	color.Cprint(fmt.Sprintf("Terminated in %v", time.Since(start)), "warn")
}

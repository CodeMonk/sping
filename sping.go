package main

// Sping "pings" a internet address by actually connecting to a TCP port.  It used to just do the raw SYN, but,
// since it is polite to close the connection too, this has been changed to do a full tcp connection (using dial), and
// then let the connection go.
//
//

import (
	"flag"
	"fmt"
	"log"

	"github.com/CodeMonk/sping/ping"
)

var (
	verbose   = flag.Bool("verbose", false, "Show verbose output")
	port      = flag.Int("port", 80, "Port to use for ping - default is 80")
	numPings  = flag.Int("numPings", -1, "Number of pings to send - default is unlimited")
	pingDelay = flag.Float64("pingDelay", 1.0, "Time (seconds) between pings")
)

func init() {
	flag.Parse()
}

func main() {

	if len(flag.Args()) != 1 {
		log.Fatal("Error: Must specify hostname.")
	}

	address := fmt.Sprintf("%s:%d", flag.Arg(0), *port)
	pinger := ping.New(address, *numPings, *pingDelay)

	pinger.DoPings()

}

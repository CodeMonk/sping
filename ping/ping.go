package ping

// Handling pinging a host
//

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Ping struct {
	Address  string
	NumPings int
	Delay    time.Duration
}

type PingResult struct {
	Source     *Ping
	SequenceNo int
	Error      error
	Sent       time.Time
	Received   time.Time
}

func (pr *PingResult) String() string {
	if pr.Error != nil {
		return fmt.Sprintf("%v: seq=%d *** Error: %v", pr.Source.Address, pr.SequenceNo, pr.Error)
	}
	return fmt.Sprintf("%v: seq=%d time=%v", pr.Source.Address, pr.SequenceNo, pr.Received.Sub(pr.Sent))
}
func (pr *PingResult) GoString() string {
	if pr.Error != nil {
		return fmt.Sprintf("%v: seq=%d *** Error: %v", pr.Source.Address, pr.SequenceNo, pr.Error)
	}
	return fmt.Sprintf("%v: seq=%d time=%v", pr.Source.Address, pr.SequenceNo, pr.Received.Sub(pr.Sent))
}

func New(address string, numPings int, delay float64) *Ping {
	p := &Ping{
		Address:  address,
		NumPings: numPings,
		Delay:    time.Duration(int64(delay * float64(time.Second))),
	}

	return p
}

// Ping performs a single ping, returning a ping result
func (p *Ping) Ping() *PingResult {
	pr := &PingResult{
		Source: p,
		Sent:   time.Now(),
	}

	// Set deadline so we don't wait forever.
	conn, err := net.DialTimeout("tcp", p.Address, 5*time.Second)
	pr.Received = time.Now()
	if err != nil {
		pr.Error = err
	} else {
		conn.Close()
	}

	// Send response
	return pr
}

func (p *Ping) asyncPing(seq int) {
	pr := p.Ping()
	pr.SequenceNo = seq

	log.Printf("%v", pr)
}

func (p *Ping) DoPings() {

	Finished := false
	for i := 0; !Finished; i++ {
		go p.asyncPing(i)

		if p.NumPings > 0 && i >= p.NumPings {
			Finished = true
		} else {
			time.Sleep(p.Delay)
		}
	}

}

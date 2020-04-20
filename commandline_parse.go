package lzr
import (
  "flag"
  "fmt"
  "time"
  "strings"
)

var (

    filename    *string
    workers     *int
    timeout     *int
    retransmit     *int
	cpuprofile	*string
	memprofile	*string
	handshake	*string
)

type options struct {

    Filename   string
    Workers    int
    Timeout    int
    Retransmit    int
	CPUProfile string
	MemProfile string
	Handshakes  []string
}


// Basic flag declarations are available for string, integer, and boolean options.
func init() {
  //port = flag.Int("port", 3000, "port number")
  fname := "default_"+string(time.Now().Format("20060102150405"))+".json"
  filename = flag.String("f", fname , "json file name")
  workers = flag.Int("w", 1000 , "number of worker threads for each channel")
  timeout = flag.Int("t", 1 , "number of seconds to wait in timeout queue")
  retransmit = flag.Int("r", 1 , "number of seconds until re-transmitting packet")
  cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
  memprofile = flag.String("memprofile", "", "write memory profile to this file")
  handshake = flag.String("h", "http" , "handshakes to scan with")
}

func Parse() *options {

    flag.Parse()
    opt := &options{
        Filename: *filename,
        Workers: *workers,
        Timeout: *timeout,
        Retransmit: *retransmit,
		CPUProfile: *cpuprofile,
		MemProfile: *memprofile,
		Handshakes: make([]string, strings.Count(*handshake,",")+1),
    }
	if !strings.Contains( *handshake, ",")	{
		opt.Handshakes[0] = *handshake
	} else {
		i := 0
		for _, h := range strings.Split( *handshake, "," ) {
			opt.Handshakes[i] = h
			i += 1
		}
	}
    fmt.Println("Writing results to file:", *filename)
    fmt.Println("Handshakes:", *handshake)
	if *memprofile != "" {
		fmt.Println("Writing memprofile to file:", *memprofile)
	}
	if *cpuprofile != "" {
		fmt.Println("Writing cpuprofile to file:", *cpuprofile)
	}
    fmt.Println("Worker threads:", *workers)
    fmt.Println("Timeout Interval (s):", *timeout)
    fmt.Println("Retransmit Interval (s):", *retransmit)
    //fmt.Println("port:", *port)
    return opt
}

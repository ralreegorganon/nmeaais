package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.String("port", "32778", "TCP port to bind to")
var source = flag.String("source", "nmeadata", "Text file containing NMEA data")
var interval = flag.Int64("interval", 1000000, "Interval in nanoseconds between sentences")

func main() {
	flag.Parse()

	file, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nmeaSentences []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nmeaSentences = append(nmeaSentences, scanner.Text())
	}

	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	ticker := time.NewTicker(time.Duration(*interval) * time.Nanosecond)

	log.Printf("Serving NMEA data from %v on port %v", *source, *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Accepting new connection %v", conn.RemoteAddr())

		w := bufio.NewWriter(conn)
		max := len(nmeaSentences)
		i := 0

		go func() {
			for _ = range ticker.C {
				if i == max {
					i = 0
				} else {
					_, err := w.WriteString(nmeaSentences[i] + "\n")
					if err != nil {
						log.Print(err)
						break
					}
					w.Flush()
					i++
				}
			}
		}()
	}
}

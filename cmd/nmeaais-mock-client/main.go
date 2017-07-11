package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var remote = flag.String("remote", "127.0.0.1:32780", "Remote address to send data to")
var source = flag.String("source", "nmeadata", "Text file containing NMEA data")
var interval = flag.Int64("interval", 1000, "Interval in milliseconds between sentences")

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

	ticker := time.NewTicker(time.Duration(*interval) * time.Millisecond)

	fault := make(chan bool)
	for {
		conn, err := net.Dial("tcp", *remote)
		if err != nil {
			time.Sleep(10)
			continue
		}
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
						fault <- true
						break
					}
					w.Flush()
					i++
				}
			}
		}()

		_ = <-fault
	}
}

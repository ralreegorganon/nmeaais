package main

import (
	"bufio"
	"flag"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/ralreegorganon/nmeaais"
)

var source = flag.String("source", "localhost:32779", "TCP source for AIS data")

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *source)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.WithField("err", err).Warn("Couldn't read packet")
			continue
		}

		log.WithField("packet", line).Info("Received packet")

		packet, err := nmeaais.Parse(line)
		if err != nil {
			log.WithFields(log.Fields{
				"err":    err,
				"packet": line,
			}).Warn("Couldn't parse packet")
			continue
		}

		log.WithFields(log.Fields{
			"packet": line,
			"parsed": packet,
		}).Info("Parsed packet")
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/ralreegorganon/nmeaais"
)

var source = flag.String("source", "localhost:32779", "TCP source for AIS data")

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	flag.Parse()

	pa := newPacketAccumulator()
	go func() {
		for m := range pa.messages {
			switch m.MessageType {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				x, err := m.GetAsPositionReportClassA()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				spew.Dump(x)
				break
			case 4:
				x, err := m.GetAsBaseStationReport()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				spew.Dump(x)
				break
			default:
				fmt.Printf("Unsupported message of type %v from %v\n", m.MessageType, m.MMSI)
				break
			}
		}
	}()

	conn, err := net.Dial("tcp", *source)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.WithField("err", err).Warn("Couldn't read packet")
			break
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
		}).Debug("Parsed packet")

		pa.packets <- packet
	}
}

type packetAccumulator struct {
	packets  chan *nmeaais.Packet
	messages chan *nmeaais.Message
}

func newPacketAccumulator() *packetAccumulator {
	pa := &packetAccumulator{
		packets:  make(chan *nmeaais.Packet),
		messages: make(chan *nmeaais.Message),
	}

	go pa.process()
	return pa
}

func (pa *packetAccumulator) process() {
	for p := range pa.packets {
		if p.FragmentCount == 1 {
			packets := []*nmeaais.Packet{p}
			m, err := nmeaais.Process(packets)
			if err != nil {
				log.WithFields(log.Fields{
					"err":    err,
					"packet": p.Raw,
				}).Warn("Failed to process packet into message")
			}
			pa.messages <- m
		}
	}
}

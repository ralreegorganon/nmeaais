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
var debug = flag.Bool("debug", false, "Run in debug mode")

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	flag.Parse()

	output := make(chan interface{})
	go func() {
		for m := range output {
			if *debug {
				spew.Dump(m)
			}
		}
	}()

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
				output <- x
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
				output <- x
				break
			case 5:
				x, err := m.GetAsStaticAndVoyageRelatedData()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 8:
				x, err := m.GetAsBinaryBroadcastMessage()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 15:
				x, err := m.GetAsInterrogation()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 18:
				x, err := m.GetAsPositionReportClassBStandard()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 20:
				x, err := m.GetAsDataLinkManagementMessage()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break

			case 21:
				x, err := m.GetAsAidToNavigationReport()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": m,
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 24:
				if ok, _ := m.IsStaticDataReportA(); ok {
					x, err := m.GetAsStaticDataReportA()
					if err != nil {
						log.WithFields(log.Fields{
							"err":     err,
							"message": m,
						}).Warn("Couldn't get specific message type")
						break
					}
					output <- x
				}
				if ok, _ := m.IsStaticDataReportB(); ok {
					x, err := m.GetAsStaticDataReportB()
					if err != nil {
						log.WithFields(log.Fields{
							"err":     err,
							"message": m,
						}).Warn("Couldn't get specific message type")
						break
					}
					output <- x
				}
				break
			default:
				fmt.Printf("Unsupported message of type %v from %v\n", m.MessageType, m.MMSI)
				for _, p := range m.Packets {
					fmt.Printf("%v\n", p.Raw)
				}
				fmt.Println()
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
	packetBuffer := make(map[int64][]*nmeaais.Packet)
	for p := range pa.packets {
		if p.FragmentCount == 1 {
			packets := []*nmeaais.Packet{p}
			pa.processAccumulatedPackets(packets)
		} else {
			packetBuffer[p.SequentialMessageID] = append(packetBuffer[p.SequentialMessageID], p)
			if p.FragmentCount == int64(len(packetBuffer[p.SequentialMessageID])) {
				pa.processAccumulatedPackets(packetBuffer[p.SequentialMessageID])
				delete(packetBuffer, p.SequentialMessageID)
			}
		}
	}
}

func (pa *packetAccumulator) processAccumulatedPackets(p []*nmeaais.Packet) {
	m, err := nmeaais.Process(p)
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"packets": p,
		}).Warn("Failed to process packet into message")
	}
	pa.messages <- m
}

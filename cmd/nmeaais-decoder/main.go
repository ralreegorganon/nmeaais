package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/ralreegorganon/nmeaais"
)

var source = flag.String("source", "localhost:32779", "TCP source for AIS data")
var debug = flag.Bool("debug", false, "Run in debug mode")
var debugFilter = flag.String("debugFilter", "", "Comma delimited list of message types to print when debugging")
var messageSkipFilter = flag.String("messageSkipFilter", "", "Comma delimited list of message types to skip once identified")

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	flag.Parse()

	filter := make(map[int64]bool)

	df := strings.Split(*debugFilter, ",")
	for _, s := range df {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			continue
		}
		filter[i] = true
	}

	skipFilter := make(map[int64]bool)

	sf := strings.Split(*messageSkipFilter, ",")
	for _, s := range sf {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			continue
		}
		skipFilter[i] = true
	}

	output := make(chan interface{})
	go func() {
		for m := range output {
			if *debug {
				messageType := reflect.ValueOf(m).Elem().FieldByName("MessageType").Int()
				_, ok := filter[messageType]
				if len(filter) == 0 || ok {
					spew.Dump(m)
				}
			}
		}
	}()

	pa := newPacketAccumulator()
	go func() {
		for m := range pa.messages {
			if *debug {
				_, ok := filter[m.MessageType]
				if len(filter) == 0 || ok {
					spew.Dump(m)
				}
			}

			_, ok := skipFilter[m.MessageType]
			if len(skipFilter) > 0 && ok {
				continue
			}

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
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 7:
				x, err := m.GetAsBinaryAcknowledge()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 16:
				x, err := m.GetAsAssignmentModeCommand()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
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
						"message": spew.Sdump(m),
					}).Warn("Couldn't get specific message type")
					break
				}
				output <- x
				break
			case 22:
				x, err := m.GetAsChannelManagement()
				if err != nil {
					log.WithFields(log.Fields{
						"err":     err,
						"message": spew.Sdump(m),
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
							"message": spew.Sdump(m),
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
							"message": spew.Sdump(m),
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

		log.WithField("packet", line).Debug("Received packet")

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

type packets []*nmeaais.Packet

func (slice packets) Len() int {
	return len(slice)
}

func (slice packets) Less(i, j int) bool {
	return slice[i].FragmentNumber < slice[j].FragmentNumber
}

func (slice packets) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (pa *packetAccumulator) process() {
	packetBuffer := make(map[int64][]*nmeaais.Packet)
	for p := range pa.packets {
		if p.FragmentCount == 1 {
			packets := []*nmeaais.Packet{p}
			pa.processAccumulatedPackets(packets)
		} else {
			packetBuffer[p.SequentialMessageID] = append(packetBuffer[p.SequentialMessageID], p)
			bufferedFragmentCount := int64(len(packetBuffer[p.SequentialMessageID]))
			if p.FragmentCount == bufferedFragmentCount {
				sort.Sort(packets(packetBuffer[p.SequentialMessageID]))
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
			"packets": spew.Sdump(p),
		}).Warn("Failed to process packet into message")
		return
	}
	pa.messages <- m
}

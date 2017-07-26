package main

import (
	"bufio"
	"flag"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ralreegorganon/nmeaais"
	log "github.com/sirupsen/logrus"
)

var source = flag.String("source", "localhost:32779", "TCP source for AIS data")
var debug = flag.Bool("debug", false, "Run in debug mode")
var debugFilter = flag.String("debugFilter", "", "Comma delimited list of message types to print when debugging")

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	filter := make(map[int64]bool)

	df := strings.Split(*debugFilter, ",")
	for _, s := range df {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			continue
		}
		filter[i] = true
	}

	decoder := nmeaais.NewDecoder()

	go func() {
		for o := range decoder.Output {
			if o.Error != nil {
				log.WithFields(log.Fields{
					"err":     o.Error,
					"packets": spew.Sdump(o.SourcePackets),
					"message": spew.Sdump(o.SourceMessage),
				}).Debug("Failed to process packets into message")
				continue
			}

			if *debug {
				messageType := reflect.ValueOf(o.DecodedMessage).Elem().FieldByName("MessageType").Int()
				_, ok := filter[messageType]
				if len(filter) == 0 || ok {
					spew.Dump(o.SourcePackets)
					spew.Dump(o.DecodedMessage)
				}
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
			log.WithField("err", err).Error("Couldn't read packet")
			close(decoder.Input)
			break
		}
		decoder.Input <- nmeaais.DecoderInput{
			Input:     line,
			Timestamp: time.Now(),
		}
	}
}

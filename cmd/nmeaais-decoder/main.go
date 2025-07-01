package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ralreegorganon/nmeaais"
)

var source = flag.String("source", "localhost:32779", "TCP source for AIS data")
var debug = flag.Bool("debug", false, "Run in debug mode")
var debugFilter = flag.String("debugFilter", "", "Comma delimited list of message types to print when debugging")

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
}

func main() {
	flag.Parse()

	if *debug {
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		slog.SetDefault(logger)
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
				packetsJSON, _ := json.MarshalIndent(o.SourcePackets, "", "  ")
				messageJSON, _ := json.MarshalIndent(o.SourceMessage, "", "  ")
				slog.Warn("Failed to process packets into message",
					"err", o.Error,
					"packets", string(packetsJSON),
					"message", string(messageJSON))
				continue
			}

			if *debug {
				messageType := reflect.ValueOf(o.DecodedMessage).Elem().FieldByName("MessageType").Int()
				_, ok := filter[messageType]
				if len(filter) == 0 || ok {
					packetsJSON, _ := json.MarshalIndent(o.SourcePackets, "", "  ")
					messageJSON, _ := json.MarshalIndent(o.DecodedMessage, "", "  ")
					fmt.Println(string(packetsJSON))
					fmt.Println(string(messageJSON))
				}
			}
		}
	}()

	conn, err := net.Dial("tcp", *source)
	if err != nil {
		slog.Error("Failed to connect to TCP source", "err", err)
		os.Exit(1)
	}

	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			slog.Error("Couldn't read packet", "err", err)
			close(decoder.Input)
			break
		}
		decoder.Input <- nmeaais.DecoderInput{
			Input:     line,
			Timestamp: time.Now(),
		}
	}
}

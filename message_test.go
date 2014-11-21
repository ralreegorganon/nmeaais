package nmeaais

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func buildPackets(raws []string) []*Packet {
	packets := make([]*Packet, 0)
	for _, raw := range raws {
		packet, err := Parse(raw)
		if err != nil {
			fmt.Println(err)
		}
		packets = append(packets, packet)
	}
	return packets
}

func TestNmeaMessageProcessing(t *testing.T) {
	Convey("When processing a multi-part message", t, func() {
		Convey("That does not contain a matching number of packets", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			Convey("The processor should return nil for the message", func() {
				So(message, ShouldBeNil)
			})
			Convey("The processor should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has packets out of sequence", func() {
			raws := []string{
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			Convey("The processor should return nil for the message", func() {
				So(message, ShouldBeNil)
			})
			Convey("The processor should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has packets from multiple messages ", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,,B,1@0000000000000,2*66",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			Convey("The processor should return nil for the message", func() {
				So(message, ShouldBeNil)
			})
			Convey("The processor should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)

			Convey("The processor should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(message, ShouldNotBeNil)
				})
			})
			Convey("The processor should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("When processing a single-part message", t, func() {
		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,1,1,,A,133m@ogP00PD;88MD5MTDww@2D7k,0*46",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)

			Convey("The processor should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(message, ShouldNotBeNil)
				})
			})
			Convey("The processor should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("When processing a single-part message", t, func() {
		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,1,1,,A,133m@ogP00PD;88MD5MTDww@2D7k,0*46",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)

			Convey("The processor should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(message, ShouldNotBeNil)
				})
			})
			Convey("The processor should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})

}

package nmeaais

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func accumulatePackets(raws []string, pa *PacketAccumulator) {
	for _, raw := range raws {
		packet, err := Parse(raw)
		if err != nil {
			fmt.Println(err)
		}
		pa.Packets <- packet
	}
	close(pa.Packets)
}

func TestPackageAccumulator(t *testing.T) {
	Convey("When processing a multi-part message", t, func() {
		Convey("That does not contain a matching number of packets", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator shouldn't return a result", func() {
				So(result, ShouldBeNil)
			})
		})

		Convey("That has packets out of sequence", func() {
			raws := []string{
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(result.Message, ShouldNotBeNil)
				})
				Convey("Where the packets have been sorted", func() {
					So(result.Packets[0].FragmentNumber, ShouldEqual, 1)
					So(result.Packets[1].FragmentNumber, ShouldEqual, 2)
				})
			})
		})

		Convey("That has packets from multiple incomplete messages ", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,,B,1@0000000000000,2*66",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator shouldn't return a result", func() {
				So(result, ShouldBeNil)
			})
		})

		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(result.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error", func() {
				So(result.Error, ShouldBeNil)
			})
		})
	})

	Convey("When processing a single-part message", t, func() {
		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,1,1,,A,133m@ogP00PD;88MD5MTDww@2D7k,0*46",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator should return a message", func() {
				Convey("Where the message is not nil", func() {
					So(result.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error", func() {
				So(result.Error, ShouldBeNil)
			})
		})
	})

	Convey("When processing a message", t, func() {
		Convey("That has an empty payload", func() {
			raws := []string{
				"!AIVDM,1,1,,A,,0*26",
			}

			pa := NewPacketAccumulator()
			accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator should return nil for the message", func() {
				So(result.Message, ShouldBeNil)
			})
			Convey("The accumulator should return an error", func() {
				So(result.Error, ShouldNotBeNil)
			})
		})
	})

}

package nmeaais

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func accumulatePackets(raws []string, pa *PacketAccumulator) {
	for _, raw := range raws {
		packet, err := Parse(raw)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pa.Packets <- packet
	}
	close(pa.Packets)
}

func accumulatePacketsWithDelay(raws []string, delay time.Duration, pa *PacketAccumulator) {
	t := time.Now()
	for _, raw := range raws {
		packet, err := ParseAtTime(raw, t)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pa.Packets <- packet
		t = t.Add(delay)
	}
	close(pa.Packets)
}

func TestPacketAccumulator(t *testing.T) {
	Convey("When processing a multi-part message", t, func() {
		Convey("That does not contain a matching number of packets", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
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
			go accumulatePackets(raws, pa)
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

		Convey("That has packets from multiple incomplete messages", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,,B,1@0000000000000,2*66",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results

			Convey("The accumulator shouldn't return a result", func() {
				So(result, ShouldBeNil)
			})
		})

		Convey("That has packets too far apart in time", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			pa := NewPacketAccumulator()
			go accumulatePacketsWithDelay(raws, time.Duration(3)*time.Second, pa)
			result := <-pa.Results

			Convey("The accumulator shouldn't return a result", func() {
				So(result, ShouldBeNil)
			})
		})

		Convey("That has interwoven packets with colliding sequential message identifier", func() {
			raws := []string{
				"!AIVDM,2,1,5,A,55MuQO000001L@;SGO8dDhiV0l4F22222222221J0000000004430E2CUCH0,0*28",
				"!AIVDM,2,1,5,A,55NHRFP2@pvmL@GS;ODPu>1<TiHE:0598uN2221620s8:4V@07li@E531H5h,0*01",
				"!AIVDM,2,2,5,A,AD`0cPs`880,2*54",
				"!AIVDM,2,2,5,A,88888888880,2*21",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)

			result1 := <-pa.Results
			Convey("The accumulator should return the first message", func() {
				Convey("Where the message is not nil", func() {
					So(result1.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error for the first message", func() {
				So(result1.Error, ShouldBeNil)
			})

			result2 := <-pa.Results
			Convey("The accumulator should return the second message", func() {
				Convey("Where the message is not nil", func() {
					So(result2.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error for the second message", func() {
				So(result2.Error, ShouldBeNil)
			})
		})

		Convey("That has interwoven packets on channels A and B", func() {
			raws := []string{
				"!AIVDM,2,1,6,B,542M92h00001@<7;?G0PD4i@R0<tqA8tj37>220o0h:2240Ht50000000000,0*3B",
				"!AIVDM,2,1,2,A,542M92h00001@<7;?G0PD4i@R0<tqA8tj37>220o0h:2240Ht500000000000000,0*3C",
				"!AIVDM,2,2,2,A,0000002,2*24",
				"!AIVDM,2,2,6,B,00000000000,2*21",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)

			result1 := <-pa.Results
			Convey("The accumulator should return the first message", func() {
				Convey("Where the message is not nil", func() {
					So(result1.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error for the first message", func() {
				So(result1.Error, ShouldBeNil)
			})

			result2 := <-pa.Results
			Convey("The accumulator should return the second message", func() {
				Convey("Where the message is not nil", func() {
					So(result2.Message, ShouldNotBeNil)
				})
			})
			Convey("The accumulator should not return an error for the second message", func() {
				So(result2.Error, ShouldBeNil)
			})
		})

		Convey("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
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
			go accumulatePackets(raws, pa)
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

		Convey("That starts with !BSVDM", func() {
			raws := []string{
				"!BSVDM,1,1,,A,13mJDd040=0Fr:TRk7wv0JwT2@Mu,0*45",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
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
			go accumulatePackets(raws, pa)
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

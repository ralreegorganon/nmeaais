package nmeaais

import (
	"fmt"

	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

var _ = Describe("PacketAccumulator", func() {
	Describe("When processing a multi-part message", func() {
		Context("That does not contain a matching number of packets", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			It("The accumulator shouldn't return a result", func() {
				Expect(result).To(BeNil())
			})
		})
		Context("That has packets out of sequence", func() {
			raws := []string{
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			Context("The accumulator should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(result.Message).To(Not(BeNil()))
				})
				It("Where the packets have been sorted", func() {
					Expect(result.Packets[0].FragmentNumber).To(Equal(int64(1)))
					Expect(result.Packets[1].FragmentNumber).To(Equal(int64(2)))
				})
			})
		})
		Context("That has packets from multiple incomplete messages", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,,B,1@0000000000000,2*66",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			It("The accumulator shouldn't return a result", func() {
				Expect(result).To(BeNil())
			})
		})
		Context("That has packets too far apart in time", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			pa := NewPacketAccumulator()
			go accumulatePacketsWithDelay(raws, time.Duration(3)*time.Second, pa)
			result := <-pa.Results
			It("The accumulator shouldn't return a result", func() {
				Expect(result).To(BeNil())
			})
		})
		Context("That has interwoven packets with colliding sequential message identifier", func() {
			raws := []string{
				"!AIVDM,2,1,5,A,55MuQO000001L@;SGO8dDhiV0l4F22222222221J0000000004430E2CUCH0,0*28",
				"!AIVDM,2,1,5,A,55NHRFP2@pvmL@GS;ODPu>1<TiHE:0598uN2221620s8:4V@07li@E531H5h,0*01",
				"!AIVDM,2,2,5,A,AD`0cPs`880,2*54",
				"!AIVDM,2,2,5,A,88888888880,2*21",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)

			result1 := <-pa.Results
			Context("The accumulator should return the first message", func() {
				It("Where the message is not nil", func() {
					Expect(result1.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error for the first message", func() {
				Expect(result1.Error).To(BeNil())
			})

			result2 := <-pa.Results
			Context("The accumulator should return the second message", func() {
				It("Where the message is not nil", func() {
					Expect(result2.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error for the second message", func() {
				Expect(result2.Error).To(BeNil())
			})
		})
		Context("That has interwoven packets on channels A and B", func() {
			raws := []string{
				"!AIVDM,2,1,6,B,542M92h00001@<7;?G0PD4i@R0<tqA8tj37>220o0h:2240Ht50000000000,0*3B",
				"!AIVDM,2,1,2,A,542M92h00001@<7;?G0PD4i@R0<tqA8tj37>220o0h:2240Ht500000000000000,0*3C",
				"!AIVDM,2,2,2,A,0000002,2*24",
				"!AIVDM,2,2,6,B,00000000000,2*21",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)

			result1 := <-pa.Results
			Context("The accumulator should return the first message", func() {
				It("Where the message is not nil", func() {
					Expect(result1.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error for the first message", func() {
				Expect(result1.Error).To(BeNil())
			})

			result2 := <-pa.Results
			Context("The accumulator should return the second message", func() {
				It("Where the message is not nil", func() {
					Expect(result2.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error for the second message", func() {
				Expect(result2.Error).To(BeNil())
			})
		})
		Context("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			Context("The accumulator should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(result.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error", func() {
				Expect(result.Error).To(BeNil())
			})
		})
	})
	Describe("When processing a single-part message", func() {
		Context("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,1,1,,A,133m@ogP00PD;88MD5MTDww@2D7k,0*46",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			Context("The accumulator should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(result.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error", func() {
				Expect(result.Error).To(BeNil())
			})
		})
		Context("That starts with !BSVDM", func() {
			raws := []string{
				"!BSVDM,1,1,,A,13mJDd040=0Fr:TRk7wv0JwT2@Mu,0*45",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			Context("The accumulator should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(result.Message).To(Not(BeNil()))
				})
			})
			It("The accumulator should not return an error", func() {
				Expect(result.Error).To(BeNil())
			})
		})
	})
	Describe("When processing a message", func() {
		Context("That has an empty payload", func() {
			raws := []string{
				"!AIVDM,1,1,,A,,0*26",
			}

			pa := NewPacketAccumulator()
			go accumulatePackets(raws, pa)
			result := <-pa.Results
			It("The accumulator should return nil for the message", func() {
				Expect(result.Message).To(BeNil())
			})
			It("The accumulator should return an error", func() {
				Expect(result.Error).To(Not(BeNil()))
			})
		})
	})

})

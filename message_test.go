package nmeaais

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func buildPackets(raws []string) []*Packet {
	var packets []*Packet
	for _, raw := range raws {
		packet, err := Parse(raw)
		if err != nil {
			fmt.Println(err)
		}
		packets = append(packets, packet)
	}
	return packets
}

var _ = Describe("NmeaMessageProcessing", func() {
	Describe("When processing a multi-part message", func() {
		Context("That does not contain a matching number of packets", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			It("The processor should return nil for the message", func() {
				Expect(message).To(BeNil())
			})
			It("The processor should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has packets out of sequence", func() {
			raws := []string{
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			It("The processor should return nil for the message", func() {
				Expect(message).To(BeNil())
			})
			It("The processor should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has packets from multiple messages ", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,,B,1@0000000000000,2*66",
			}
			packets := buildPackets(raws)
			message, err := Process(packets)
			It("The processor should return nil for the message", func() {
				Expect(message).To(BeNil())
			})
			It("The processor should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,2,1,3,B,55P5TL01VIaAL@7WKO@mBplU@<PDhh000000001S;AJ::4A80?4i@E53,0*3E",
				"!AIVDM,2,2,3,B,1@0000000000000,2*55",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			Context("The processor should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(message).To(Not(BeNil()))
				})
			})
			It("The processor should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("When processing a single-part message", func() {
		Context("That is a valid NMEA 0183 format", func() {
			raws := []string{
				"!AIVDM,1,1,,A,133m@ogP00PD;88MD5MTDww@2D7k,0*46",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			Context("The processor should return a message", func() {
				It("Where the message is not nil", func() {
					Expect(message).To(Not(BeNil()))
				})
			})
			It("The processor should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("When processing a message", func() {
		Context("That has an empty payload", func() {
			raws := []string{
				"!AIVDM,1,1,,A,,0*26",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			It("The processor should return nil for the message", func() {
				Expect(message).To(BeNil())
			})
			It("The processor should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
	})

})

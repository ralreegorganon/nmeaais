package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NmeaPacketParsing", func() {
	Describe("When parsing a raw packet", func() {
		Context("That does not start with !", func() {
			raw := "$GPAAM,A,A,0.10,N,WPTNME*32"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That does not have 7 parts", func() {
			raw := "!AIVDM,1,1,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has an invalid fragment count", func() {
			raw := "!AIVDM,A,1,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has an invalid fragment number", func() {
			raw := "!AIVDM,1,Z,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has an invalid sequential message id", func() {
			raw := "!AIVDM,1,1,Z,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has an invalid radio channel", func() {
			raw := "!AIVDM,1,1,,Z,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has an invalid fill bit value", func() {
			raw := "!AIVDM,1,1,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,6*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That has a non-matching checksum", func() {
			raw := "!AIVDM,1,1,,B,177KQJ5000G?tO`K>RA1aUbN0TKH,0*5C"
			packet, err := Parse(raw)
			It("The parser should return nil for the packet", func() {
				Expect(packet).To(BeNil())
			})
			It("The parser should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
		Context("That is a valid NMEA 0183 format", func() {
			Context("That starts with !AIVDM", func() {
				raw := "!AIVDM,1,1,,B,176u=;?000`:RhH<h?IP0CBT08;5,0*50"
				packet, err := Parse(raw)
				Context("The parser should return a packet", func() {
					It("Where the start delimiter is correct", func() {
						Expect(packet.StartDelimiter).To(Equal("!"))
					})
					It("Where the tag is correct", func() {
						Expect(packet.Tag).To(Equal("AIVDM"))
					})
					It("Where the fragment count is correct", func() {
						Expect(packet.FragmentCount).To(Equal(int64(1)))
					})
					It("Where the fragment number is correct", func() {
						Expect(packet.FragmentNumber).To(Equal(int64(1)))
					})
					It("Where the sequential message number is correct", func() {
						Expect(packet.SequentialMessageID).To(Equal(int64(0)))
					})
					It("Where the radio channel is correct", func() {
						Expect(packet.RadioChannel).To(Equal("B"))
					})
					It("Where the payload is correct", func() {
						Expect(packet.Payload).To(Equal("176u=;?000`:RhH<h?IP0CBT08;5"))
					})
					It("Where the fill bits are correct", func() {
						Expect(packet.FillBits).To(Equal(int64(0)))
					})
					It("Where the checksum is correct", func() {
						Expect(packet.Checksum).To(Equal("50"))
					})
				})
				It("The parser should not return an error", func() {
					Expect(err).To(BeNil())
				})
			})
			Context("That starts with !BSVDM", func() {
				raw := "!BSVDM,1,1,,A,13mJDd040=0Fr:TRk7wv0JwT2@Mu,0*45"
				packet, err := Parse(raw)
				Context("The parser should return a packet", func() {
					It("Where the start delimiter is correct", func() {
						Expect(packet.StartDelimiter).To(Equal("!"))
					})
					It("Where the tag is correct", func() {
						Expect(packet.Tag).To(Equal("BSVDM"))
					})
					It("Where the fragment count is correct", func() {
						Expect(packet.FragmentCount).To(Equal(int64(1)))
					})
					It("Where the fragment number is correct", func() {
						Expect(packet.FragmentNumber).To(Equal(int64(1)))
					})
					It("Where the sequential message number is correct", func() {
						Expect(packet.SequentialMessageID).To(Equal(int64(0)))
					})
					It("Where the radio channel is correct", func() {
						Expect(packet.RadioChannel).To(Equal("A"))
					})
					It("Where the payload is correct", func() {
						Expect(packet.Payload).To(Equal("13mJDd040=0Fr:TRk7wv0JwT2@Mu"))
					})
					It("Where the fill bits are correct", func() {
						Expect(packet.FillBits).To(Equal(int64(0)))
					})
					It("Where the checksum is correct", func() {
						Expect(packet.Checksum).To(Equal("45"))
					})
				})
				It("The parser should not return an error", func() {
					Expect(err).To(BeNil())
				})
			})
		})
	})
})

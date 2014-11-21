package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNmeaPacketParsing(t *testing.T) {
	Convey("When parsing a raw packet", t, func() {
		Convey("That that does not start with !", func() {
			raw := "$GPAAM,A,A,0.10,N,WPTNME*32"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That does not have 7 parts", func() {
			raw := "!AIVDM,1,1,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has an invalid fragment count", func() {
			raw := "!AIVDM,A,1,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has an invalid fragment number", func() {
			raw := "!AIVDM,1,Z,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has an invalid sequential message id", func() {
			raw := "!AIVDM,1,1,Z,B,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has an invalid radio channel", func() {
			raw := "!AIVDM,1,1,,Z,177KQJ5000G?tO`K>RA1wUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has an invalid fill bit value", func() {
			raw := "!AIVDM,1,1,,B,177KQJ5000G?tO`K>RA1wUbN0TKH,6*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That has a non-matching checksum", func() {
			raw := "!AIVDM,1,1,,B,177KQJ5000G?tO`K>RA1aUbN0TKH,0*5C"
			packet, err := Parse(raw)
			Convey("The parser should return nil for the packet", func() {
				So(packet, ShouldBeNil)
			})
			Convey("The parser should return an error", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("That is a valid NMEA 0183 format", func() {
			raw := "!AIVDM,1,1,,B,176u=;?000`:RhH<h?IP0CBT08;5,0*50"
			packet, err := Parse(raw)
			Convey("The parser should return a packet", func() {
				Convey("Where the start delimiter is correct", func() {
					So(packet.StartDelimiter, ShouldEqual, "!")
				})
				Convey("Where the tag is correct", func() {
					So(packet.Tag, ShouldEqual, "AIVDM")
				})
				Convey("Where the fragment count is correct", func() {
					So(packet.FragmentCount, ShouldEqual, 1)
				})
				Convey("Where the fragment number is correct", func() {
					So(packet.FragmentNumber, ShouldEqual, 1)
				})
				Convey("Where the sequential message number is correct", func() {
					So(packet.SequentialMessageId, ShouldEqual, 0)
				})
				Convey("Where the radio channel is correct", func() {
					So(packet.RadioChannel, ShouldEqual, "B")
				})
				Convey("Where the payload is correct", func() {
					So(packet.Payload, ShouldEqual, "176u=;?000`:RhH<h?IP0CBT08;5")
				})
				Convey("Where the fill bits are correct", func() {
					So(packet.FillBits, ShouldEqual, 0)
				})
				Convey("Where the checksum is correct", func() {
					So(packet.Checksum, ShouldEqual, "50")
				})
			})
			Convey("The parser should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

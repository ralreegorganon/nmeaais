package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType25MessageProcessing(t *testing.T) {
	Convey("When processing a type 25 message", t, func() {
		/* NEED AN EXAMPLE
		Convey("That is addressed and structured", func() {
			raws := []string{
				"",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type25, err := message.GetAsSingleSlotBinaryMessage()

			expected := &SingleSlotBinaryMessage{}

			Convey("The get should return a type 25 message", func() {
				Convey("Where the message is not nil", func() {
					So(type25, ShouldNotBeNil)
				})
			})

			Convey("The get should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The fields should be populated correctly", func() {
				So(type25, ShouldResemble, expected)
			})
		})
		*/
		Convey("That is addressed and unstructured", func() {
			raws := []string{
				"!AIVDM,1,1,,A,I6SWo?8P00a3PKpEKEVj0?vNP<65,0*73",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type25, err := message.GetAsSingleSlotBinaryMessage()

			expected := &SingleSlotBinaryMessage{
				MessageType:        25,
				RepeatIndicator:    0,
				MMSI:               440006460,
				Addressed:          true,
				Structured:         false,
				DestinationMMSI:    134218384,
				DesignatedAreaCode: 0,
				FunctionalID:       0,
				Data:               []uint8{224, 111, 133, 91, 86, 108, 128, 63, 231, 160, 48, 97},
			}

			Convey("The get should return a type 25 message", func() {
				Convey("Where the message is not nil", func() {
					So(type25, ShouldNotBeNil)
				})
			})

			Convey("The get should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The fields should be populated correctly", func() {
				So(type25, ShouldResemble, expected)
			})
		})

		Convey("That is unaddressed and structured", func() {
			raws := []string{
				"!AIVDM,1,1,,A,I8IRGB40QPPa0:<HP::V=gwv0l48,0*0E",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type25, err := message.GetAsSingleSlotBinaryMessage()

			expected := &SingleSlotBinaryMessage{
				MessageType:        25,
				RepeatIndicator:    0,
				MMSI:               563648328,
				Addressed:          false,
				Structured:         true,
				DestinationMMSI:    0,
				DesignatedAreaCode: 2,
				FunctionalID:       6,
				Data:               []uint8{8, 41, 0, 163, 24, 128, 162, 166, 54, 255, 254, 3, 65, 8},
			}

			Convey("The get should return a type 25 message", func() {
				Convey("Where the message is not nil", func() {
					So(type25, ShouldNotBeNil)
				})
			})

			Convey("The get should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The fields should be populated correctly", func() {
				So(type25, ShouldResemble, expected)
			})
		})

		Convey("That is unaddressed and unstructured", func() {
			raws := []string{
				"!AIVDM,1,1,,A,I6SWVNP001a3P8FEKNf=Qb0@00S8,0*6B",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type25, err := message.GetAsSingleSlotBinaryMessage()

			expected := &SingleSlotBinaryMessage{
				MessageType:        25,
				RepeatIndicator:    0,
				MMSI:               440002170,
				Addressed:          false,
				Structured:         false,
				DestinationMMSI:    0,
				DesignatedAreaCode: 0,
				FunctionalID:       0,
				Data:               []uint8{0, 0, 26, 67, 128, 133, 149, 109, 235, 141, 134, 160, 16, 0, 8, 200},
			}

			Convey("The get should return a type 25 message", func() {
				Convey("Where the message is not nil", func() {
					So(type25, ShouldNotBeNil)
				})
			})

			Convey("The get should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The fields should be populated correctly", func() {
				So(type25, ShouldResemble, expected)
			})
		})
	})
}

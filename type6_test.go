package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType6MessageProcessing(t *testing.T) {
	Convey("When processing a type 6 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,65D7EH5DoW300400A@E=B04<d0,4*46",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type6, err := message.GetAsBinaryAddressedMessage()

		expected := &BinaryAddressedMessage{
			MessageType:        6,
			RepeatIndicator:    0,
			MMSI:               356636000,
			SequenceNumber:     1,
			DestinationMMSI:    355966000,
			RetransmitFlag:     false,
			DesignatedAreaCode: 1,
			FunctionalID:       0,
			Data:               []uint8{0, 69, 5, 77, 72, 1, 12, 176},
		}

		Convey("The get should return a type 6 message", func() {
			Convey("Where the message is not nil", func() {
				So(type6, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type6, ShouldResemble, expected)
		})
	})
}

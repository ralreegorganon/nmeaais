package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType7MessageProcessing(t *testing.T) {
	Convey("When processing a type 7 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,75MwQW2G`lEH,0*6C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type7, err := message.GetAsBinaryAcknowledge()

		expected := &BinaryAcknowledge{
			MessageType:     7,
			RepeatIndicator: 0,
			MMSI:            366993820,
			MMSI1:           636014934,
			MMSI1Sequence:   0,
		}

		Convey("The get should return a type 7 message", func() {
			Convey("Where the message is not nil", func() {
				So(type7, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type7, ShouldResemble, expected)
		})
	})
}

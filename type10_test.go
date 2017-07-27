package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType10MessageProcessing(t *testing.T) {
	Convey("When processing a type 10 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,:5MwvSQGRlc8,0*45",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type10, err := message.GetAsUTCDateInquiry()

		expected := &UTCDateInquiry{
			MessageType:     10,
			RepeatIndicator: 0,
			MMSI:            367001230,
			DestinationMMSI: 367186610,
		}

		Convey("The get should return a type 10 message", func() {
			Convey("Where the message is not nil", func() {
				So(type10, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type10, ShouldResemble, expected)
		})
	})
}

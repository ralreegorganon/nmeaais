package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType14MessageProcessing(t *testing.T) {
	Convey("When processing a type 14 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,>>M;1IM<59B1@E=@,0*5E",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type14, err := message.GetAsSafetyRelatedBroadcast()

		expected := &SafetyRelatedBroadcast{
			MessageType:     14,
			RepeatIndicator: 0,
			MMSI:            970113381,
			SequenceNumber:  0,
			DestinationMMSI: 0,
			RetransmitFlag:  false,
			Text:            "SART TEST",
		}

		Convey("The get should return a type 14 message", func() {
			Convey("Where the message is not nil", func() {
				So(type14, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type14, ShouldResemble, expected)
		})
	})
}

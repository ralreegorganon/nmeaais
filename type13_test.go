package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType13MessageProcessing(t *testing.T) {
	Convey("When processing a type 13 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,=4WCf22Gaw0`,0*5C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type13, err := message.GetAsSafetyRelatedAcknowledge()

		expected := &SafetyRelatedAcknowledge{
			MessageType:     13,
			RepeatIndicator: 0,
			MMSI:            309653000,
			MMSI1:           636091402,
			MMSI1Sequence:   0,
		}

		Convey("The get should return a type 13 message", func() {
			Convey("Where the message is not nil", func() {
				So(type13, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type13, ShouldResemble, expected)
		})
	})
}

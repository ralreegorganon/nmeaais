package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType27MessageProcessing(t *testing.T) {
	Convey("When processing a type 27 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,Km31e<1KQ?SO4P5d,0*66",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type27, err := message.GetAsLongRangeAISBroadcast()

		expected := &LongRangeAISBroadcast{
			MessageType:        27,
			RepeatIndicator:    3,
			MMSI:               338718000,
			PositionAccuracy:   false,
			RAIM:               false,
			NavigationStatus:   "Moored",
			Longitude:          -122.35,
			Latitude:           47.58833333333333,
			SpeedOverGround:    0,
			CourseOverGround:   91,
			GNSSPositionStatus: 0,
		}

		Convey("The get should return a type 27 message", func() {
			Convey("Where the message is not nil", func() {
				So(type27, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type27, ShouldResemble, expected)
		})
	})
}

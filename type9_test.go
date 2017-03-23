package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType9MessageProcessing(t *testing.T) {
	Convey("When processing a type 9 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,91b76w001L163a8QIdP8O<h00PS6,0*10",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type9, err := message.GetAsStandardSARAircraftPositionReport()

		expected := &StandardSARAircraftPositionReport{
			MessageType:      9,
			RepeatIndicator:  0,
			MMSI:             111265532,
			Altitude:         0,
			SpeedOverGround:  92,
			PositionAccuracy: false,
			Longitude:        15.304166666666667,
			Latitude:         58.373333333333335,
			CourseOverGround: 217.2,
			TimeStamp:        51,
			DTE:              false,
			RAIM:             false,
			RadioStatus:      133318,
		}

		Convey("The get should return a type 9 message", func() {
			Convey("Where the message is not nil", func() {
				So(type9, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type9, ShouldResemble, expected)
		})
	})
}

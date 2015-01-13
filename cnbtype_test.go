package nmeaais

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType1MessageProcessing(t *testing.T) {
	Convey("When processing a cnb type (1,2,3) message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,15RTgt0PAso;90TKcjM8h6g208CQ,0*4A",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type1, err := message.GetAsPositionReportClassA()

		expected := &PositionReportClassA{
			MessageType:       1,
			RepeatIndicator:   0,
			MMSI:              371798000,
			NavigationStatus:  "Under way using engine",
			RateOfTurn:        math.Inf(-1),
			SpeedOverGround:   12.3,
			PositionAccuracy:  true,
			Longitude:         -123.39538333333333,
			Latitude:          48.38163333333333,
			CourseOverGround:  224,
			TrueHeading:       215,
			TimeStamp:         33,
			ManeuverIndicator: "Not available",
			RAIM:              false,
			RadioStatus:       34017,
		}

		Convey("The get should return a cnb type (1,2,3) message", func() {
			Convey("Where the message is not nil", func() {
				So(type1, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type1, ShouldResemble, expected)
		})
	})

	Convey("When processing an invalid cnb type (1,2,3) message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,13n;;V001`0q,0*0C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type1, err := message.GetAsPositionReportClassA()

		Convey("The get should return a cnb type (1,2,3) message", func() {
			Convey("Where the message is nil", func() {
				So(type1, ShouldBeNil)
			})
		})

		Convey("The get should return an error", func() {
			So(err, ShouldNotBeNil)
		})

	})
}

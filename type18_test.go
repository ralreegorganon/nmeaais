package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType18MessageProcessing(t *testing.T) {
	Convey("When processing a type 18 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,B52Mu0@00El8HO6oJS<Igwk5kP06,0*7B",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type18, err := message.GetAsPositionReportClassBStandard()

		expected := &PositionReportClassBStandard{
			MessageType:      18,
			RepeatIndicator:  0,
			MMSI:             338132225,
			SpeedOverGround:  0.1,
			PositionAccuracy: false,
			Longitude:        -122.21941666666666,
			Latitude:         47.99581833333333,
			CourseOverGround: 41.1,
			TrueHeading:      511,
			TimeStamp:        38,
			CSUnit:           true,
			Display:          false,
			DSC:              true,
			Band:             true,
			Message22:        true,
			Assigned:         false,
			RAIM:             false,
			RadioStatus:      917510,
		}

		Convey("The get should return a type 18 message", func() {
			Convey("Where the message is not nil", func() {
				So(type18, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type18, ShouldResemble, expected)
		})
	})

	Convey("When processing an invalid type 18 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,13n;;V001`0q,0*0C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type18, err := message.GetAsPositionReportClassBStandard()

		Convey("The get should return a type 18 message", func() {
			Convey("Where the message is nil", func() {
				So(type18, ShouldBeNil)
			})
		})

		Convey("The get should return an error", func() {
			So(err, ShouldNotBeNil)
		})

	})
}

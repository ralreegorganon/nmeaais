package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType19MessageProcessing(t *testing.T) {
	Convey("When processing a type 19 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,C5N3SRgPEnJGEBT>NhWAwwo862PaLELTBJ:V00000000S0D:R220,0*0B",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type19, err := message.GetAsPositionReportClassBExtended()

		expected := &PositionReportClassBExtended{
			MessageType:          19,
			RepeatIndicator:      0,
			MMSI:                 367059850,
			SpeedOverGround:      8.7,
			PositionAccuracy:     false,
			Longitude:            -88.81039166666666,
			Latitude:             29.543695,
			CourseOverGround:     335.9,
			TrueHeading:          511,
			TimeStamp:            46,
			VesselName:           "CAPT.J.RIMES",
			ShipType:             "Cargo, all ships of this type",
			DimensionToBow:       5,
			DimensionToStern:     21,
			DimensionToPort:      4,
			DimensionToStarboard: 4,
			EPFDType:             "GPS",
			RAIM:                 false,
			DTE:                  false,
			Assigned:             false,
		}

		Convey("The get should return a type 19 message", func() {
			Convey("Where the message is not nil", func() {
				So(type19, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type19, ShouldResemble, expected)
		})
	})
}

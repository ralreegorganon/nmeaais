package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType21MessageProcessing(t *testing.T) {
	Convey("When processing a type 21 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,E>k1kFd1WWUh4W62b@1:WdhHpP0J`lV<AQ@:000003vP10,4*69",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type21, err := message.GetAsAidToNavigationReport()

		expected := &AidToNavigationReport{
			MessageType:          21,
			RepeatIndicator:      0,
			MMSI:                 993031002,
			AidType:              "Port hand Mark",
			Name:                 "COOK INLET BUOY 11",
			PositionAccuracy:     true,
			Longitude:            99.18689333333333,
			Latitude:             2.0376316666666665,
			DimensionToBow:       0,
			DimensionToStern:     0,
			DimensionToPort:      0,
			DimensionToStarboard: 0,
			EPFDType:             "Surveyed",
			UTCSecond:            61,
			OffPositionIndicator: false,
			RAIM:                 false,
			VirtualAid:           true,
			AssignedMode:         false,
			NameExtension:        "",
		}

		Convey("The get should return a type 21 message", func() {
			Convey("Where the message is not nil", func() {
				So(type21, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type21, ShouldResemble, expected)
		})
	})
}

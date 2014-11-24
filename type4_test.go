package nmeaais

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType4MessageProcessing(t *testing.T) {
	Convey("When processing a type 4 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,402R3dAurtDNn0n7C@QIev100@PE,0*45",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type4, err := message.GetAsBaseStationReport()

		expected := &BaseStationReport{
			MessageType:      4,
			RepeatIndicator:  0,
			MMSI:             2655153,
			PositionAccuracy: false,
			TimeStamp:        time.Date(2014, time.November, 24, 20, 30, 54, 0, time.UTC),
			Longitude:        11.8214,
			Latitude:         58.37396,
			EPFDType:         "GPS",
			RAIM:             false,
			RadioStatus:      67605,
		}

		Convey("The get should return a type 4 message", func() {
			Convey("Where the message is not nil", func() {
				So(type4, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type4, ShouldResemble, expected)
		})
	})
}

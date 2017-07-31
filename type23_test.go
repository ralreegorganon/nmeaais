package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType23MessageProcessing(t *testing.T) {
	Convey("When processing a type 23 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,G02:Kn01R`sn@291nj600000900,2*12",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type23, err := message.GetAsGroupAssignmentCommand()

		expected := &GroupAssignmentCommand{
			MessageType:     23,
			RepeatIndicator: 0,
			MMSI:            2268120,
			NELongitude:     2.63,
			NELatitutde:     51.07,
			SWLongitude:     1.8266666666666667,
			SWLatitude:      50.68,
			StationType:     "Regional use and inland waterways",
			ShipType:        "Not available",
			TxRxMode:        0,
			ReportInterval:  "Next Shorter Reporting Interval",
			QuietTime:       0,
		}

		Convey("The get should return a type 23 message", func() {
			Convey("Where the message is not nil", func() {
				So(type23, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type23, ShouldResemble, expected)
		})
	})
}

package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType5MessageProcessing(t *testing.T) {
	Convey("When processing a type 5 message", t, func() {
		raws := []string{
			"!AIVDM,2,1,9,B,55OER>01sWpeL@GS?CM0th5:1=@u8n222222220P1PJ354AB0;PCPj3lPAiH,0*1B",
			"!AIVDM,2,2,9,B,88888888880,2*2E",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type5, err := message.GetAsStaticAndVoyageRelatedData()

		expected := &StaticAndVoyageRelatedData{
			MessageType:          5,
			RepeatIndicator:      0,
			MMSI:                 368403000,
			AISVersion:           0,
			IMONumber:            8101771,
			CallSign:             "WDE8347",
			VesselName:           "POLAR STORM",
			ShipType:             "Towing: length exceeds 200m or breadth exceeds 25m",
			DimensionToBow:       12,
			DimensionToStern:     26,
			DimensionToPort:      3,
			DimensionToStarboard: 5,
			EPFDType:             "GPS",
			ETAMonth:             1,
			ETADay:               2,
			ETAHour:              18,
			ETAMinute:            0,
			Draught:              4.6,
			Destination:          "ANCHORAGE",
			DTE:                  false,
		}

		Convey("The get should return a type 5 message", func() {
			Convey("Where the message is not nil", func() {
				So(type5, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type5, ShouldResemble, expected)
		})
	})

	Convey("When processing a type 5 message with a short payload", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,50000010000000000000000000000000,0*22",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type5, err := message.GetAsStaticAndVoyageRelatedData()

		Convey("The get should return an error", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("The get should return nil for the message", func() {
			So(type5, ShouldBeNil)
		})
	})
}

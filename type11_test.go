package nmeaais

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType11MessageProcessing(t *testing.T) {
	Convey("When processing a type 11 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,;5N;BdQuw;:i5mAi:nS27jQ02000,0*3B",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type11, err := message.GetAsUTCDateResponse()

		expected := &UTCDateResponse{
			MessageType:      11,
			RepeatIndicator:  0,
			MMSI:             367186610,
			PositionAccuracy: true,
			TimeStamp:        time.Date(2015, time.December, 22, 10, 49, 5, 0, time.UTC),
			Longitude:        -149.90960833333332,
			Latitude:         61.22487,
			EPFDType:         "GPS",
			RAIM:             true,
			RadioStatus:      0,
		}

		Convey("The get should return a type 11 message", func() {
			Convey("Where the message is not nil", func() {
				So(type11, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type11, ShouldResemble, expected)
		})
	})
}

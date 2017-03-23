package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType16MessageProcessing(t *testing.T) {
	Convey("When processing a type 16 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,@h3OvjBGaw3h3h0000000000,0*7E",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type16, err := message.GetAsAssignmentModeCommand()

		expected := &AssignmentModeCommand{
			MessageType:      16,
			RepeatIndicator:  3,
			MMSI:             3669705,
			DestinationMMSI1: 636091452,
			Offset1:          60,
			Increment1:       0,
			DestinationMMSI2: 0,
			Offset2:          0,
			Increment2:       0,
		}

		Convey("The get should return a type 16 message", func() {
			Convey("Where the message is not nil", func() {
				So(type16, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type16, ShouldResemble, expected)
		})
	})
}

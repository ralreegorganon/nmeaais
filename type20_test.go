package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType20MessageProcessing(t *testing.T) {
	Convey("When processing a type 20 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,Dh3Ovj@11N>6;HfGL00Nfp0,2*1B",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type20, err := message.GetAsDataLinkManagementMessage()

		expected := &DataLinkManagementMessage{
			MessageType:     20,
			RepeatIndicator: 3,
			MMSI:            3669705,
			Offset1:         16,
			ReservedSlots1:  5,
			Timeout1:        7,
			Increment1:      225,
			Offset2:         2230,
			ReservedSlots2:  2,
			Timeout2:        7,
			Increment2:      375,
			Offset3:         0,
			ReservedSlots3:  1,
			Timeout3:        7,
			Increment3:      750,
			Offset4:         0,
			ReservedSlots4:  0,
			Timeout4:        0,
			Increment4:      0,
		}

		Convey("The get should return a type 20 message", func() {
			Convey("Where the message is not nil", func() {
				So(type20, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type20, ShouldResemble, expected)
		})
	})
}

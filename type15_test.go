package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType15MessageProcessing(t *testing.T) {
	Convey("When processing a type 15 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,?h3Ovj@p>iBPD00,2*21",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type15, err := message.GetAsInterrogation()

		expected := &Interrogation{
			MessageType:        15,
			RepeatIndicator:    3,
			MMSI:               3669705,
			InterrogatedMMSI1:  235849000,
			FirstMessageType1:  5,
			FirstSlotOffset1:   0,
			SecondMessageType1: 0,
			SecondSlotOffset1:  0,
			InterrogatedMMSI2:  0,
			FirstMessageType2:  0,
			FirstSlotOffset2:   0,
		}

		Convey("The get should return a type 15 message", func() {
			Convey("Where the message is not nil", func() {
				So(type15, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type15, ShouldResemble, expected)
		})
	})
}

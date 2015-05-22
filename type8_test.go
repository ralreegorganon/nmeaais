package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType8MessageProcessing(t *testing.T) {
	Convey("When processing a type 8 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,85Mwqd1Kf4dldnKQ<>bW6RGmDu<6U5f1>W<LMGV85qe;dkv@rN5h,0*7D",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type8, err := message.GetAsBinaryBroadcastMessage()

		expected := &BinaryBroadcastMessage{
			MessageType:        8,
			RepeatIndicator:    0,
			MMSI:               366999984,
			DesignatedAreaCode: 366,
			FunctionalID:       56,
			Data:               []byte{75, 52, 179, 102, 225, 48, 234, 167, 26, 37, 245, 83, 211, 6, 148, 91, 129, 58, 115, 28, 117, 121, 136, 23, 155, 75, 179, 63, 144, 233, 225, 112},
		}

		Convey("The get should return a type 8 message", func() {
			Convey("Where the message is not nil", func() {
				So(type8, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type8, ShouldResemble, expected)
		})
	})
}

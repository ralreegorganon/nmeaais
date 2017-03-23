package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType17MessageProcessing(t *testing.T) {
	Convey("When processing a type 17 message", t, func() {
		raws := []string{
			"!AIVDM,2,1,5,A,A02VqLPA4I6C07h5Ed1h<OrsuBTTwS?r:C?w`?la<gno1RTRwSP9:BcurA8a,0*3A",
			"!AIVDM,2,2,5,A,:Oko02TSwu8<:Jbb,0*11",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type17, err := message.GetAsDGNSSBroadcastBinaryMessage()

		expected := &DGNSSBroadcastBinaryMessage{
			MessageType:     17,
			RepeatIndicator: 0,
			MMSI:            2734450,
			Longitude:       29.13,
			Latitude:        59.986666666666665,
			Payload:         []byte{124, 5, 86, 192, 112, 49, 254, 187, 245, 41, 36, 254, 51, 250, 41, 51, 255, 160, 253, 41, 50, 253, 183, 6, 41, 34, 254, 56, 9, 41, 42, 253, 233, 18, 41, 41, 252, 247, 0, 41, 35, 255, 210, 12, 41, 170, 170},
		}

		Convey("The get should return a type 17 message", func() {
			Convey("Where the message is not nil", func() {
				So(type17, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type17, ShouldResemble, expected)
		})
	})
}

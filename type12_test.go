package nmeaais

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestType12MessageProcessing(t *testing.T) {
	Convey("When processing a type 12 message", t, func() {
		raws := []string{
			"!AIVDM,2,1,1,A,<D62222222208:5vmEEEOPAGEso0009m5@CFb;vNnQIsW008t>AOOfbbbWp4,0*40",
			"!AIVDM,2,2,1,A,=fD:8=w0?A@,2*4C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type12, err := message.GetAsAddressedSafetyRelated()

		expected := &AddressedSafetyRelated{
			MessageType:     12,
			RepeatIndicator: 1,
			MMSI:            274760200,
			SequenceNumber:  0,
			DestinationMMSI: 545392672,
			RetransmitFlag:  false,
			Text:            "HJE>5UUU- QWU;7",
		}

		Convey("The get should return a type 12 message", func() {
			Convey("Where the message is not nil", func() {
				So(type12, ShouldNotBeNil)
			})
		})

		Convey("The get should not return an error", func() {
			So(err, ShouldBeNil)
		})

		Convey("The fields should be populated correctly", func() {
			So(type12, ShouldResemble, expected)
		})
	})
}

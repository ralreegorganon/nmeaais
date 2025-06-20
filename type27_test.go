package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type27MessageProcessing", func() {
	Describe("When processing a type 27 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,Km31e<1KQ?SO4P5d,0*66",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type27, err := message.GetAsLongRangeAISBroadcast()

		expected := &LongRangeAISBroadcast{
			MessageType:        27,
			RepeatIndicator:    3,
			MMSI:               338718000,
			PositionAccuracy:   false,
			RAIM:               false,
			NavigationStatus:   "Moored",
			Longitude:          -122.35,
			Latitude:           47.58833333333333,
			SpeedOverGround:    0,
			CourseOverGround:   91,
			GNSSPositionStatus: 0,
		}
		Context("The get should return a type 27 message", func() {
			It("Where the message is not nil", func() {
				Expect(type27).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type27).To(Equal(expected))
		})
	})
})

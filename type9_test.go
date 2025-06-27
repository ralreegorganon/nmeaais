package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type9MessageProcessing", func() {
	Describe("When processing a type 9 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,91b76w001L163a8QIdP8O<h00PS6,0*10",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type9, err := message.GetAsStandardSARAircraftPositionReport()

		expected := &StandardSARAircraftPositionReport{
			MessageType:      9,
			RepeatIndicator:  0,
			MMSI:             111265532,
			Altitude:         0,
			SpeedOverGround:  92,
			PositionAccuracy: false,
			Longitude:        15.304166666666667,
			Latitude:         58.373333333333335,
			CourseOverGround: 217.2,
			TimeStamp:        51,
			DTE:              false,
			RAIM:             false,
			RadioStatus:      133318,
		}
		Context("The get should return a type 9 message", func() {
			It("Where the message is not nil", func() {
				Expect(type9).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type9).To(Equal(expected))
		})
	})
})

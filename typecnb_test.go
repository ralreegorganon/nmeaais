package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type1MessageProcessing", func() {
	Describe("When processing a cnb type (1,2,3) message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,15RTgt0PAso;90TKcjM8h6g208CQ,0*4A",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type1, err := message.GetAsPositionReportClassA()

		expected := &PositionReportClassA{
			MessageType:       1,
			RepeatIndicator:   0,
			MMSI:              371798000,
			NavigationStatus:  "Under way using engine",
			RateOfTurn:        -127,
			SpeedOverGround:   12.3,
			PositionAccuracy:  true,
			Longitude:         -123.39538333333333,
			Latitude:          48.38163333333333,
			CourseOverGround:  224,
			TrueHeading:       215,
			TimeStamp:         33,
			ManeuverIndicator: "Not available",
			RAIM:              false,
			RadioStatus:       34017,
		}
		Context("The get should return a cnb type (1,2,3) message", func() {
			It("Where the message is not nil", func() {
				Expect(type1).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type1).To(Equal(expected))
		})
	})
	Describe("When processing an invalid cnb type (1,2,3) message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,13n;;V001`0q,0*0C",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type1, err := message.GetAsPositionReportClassA()
		Context("The get should return a cnb type (1,2,3) message", func() {
			It("Where the message is nil", func() {
				Expect(type1).To(BeNil())
			})
		})
		It("The get should return an error", func() {
			Expect(err).To(Not(BeNil()))
		})
	})
})

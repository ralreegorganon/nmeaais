package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type19MessageProcessing", func() {
	Describe("When processing a type 19 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,C5N3SRgPEnJGEBT>NhWAwwo862PaLELTBJ:V00000000S0D:R220,0*0B",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type19, err := message.GetAsPositionReportClassBExtended()

		expected := &PositionReportClassBExtended{
			MessageType:          19,
			RepeatIndicator:      0,
			MMSI:                 367059850,
			SpeedOverGround:      8.7,
			PositionAccuracy:     false,
			Longitude:            -88.81039166666666,
			Latitude:             29.543695,
			CourseOverGround:     335.9,
			TrueHeading:          511,
			TimeStamp:            46,
			VesselName:           "CAPT.J.RIMES",
			ShipType:             "Cargo, all ships of this type",
			DimensionToBow:       5,
			DimensionToStern:     21,
			DimensionToPort:      4,
			DimensionToStarboard: 4,
			EPFDType:             "GPS",
			RAIM:                 false,
			DTE:                  false,
			Assigned:             false,
		}
		Context("The get should return a type 19 message", func() {
			It("Where the message is not nil", func() {
				Expect(type19).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type19).To(Equal(expected))
		})
	})
})

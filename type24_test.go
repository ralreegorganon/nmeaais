package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type24AMessageProcessing", func() {
	Describe("When processing a type 24 A message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,H7P<1>1LPU@D8U8A<0000000000,2*6C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type24, err := message.GetAsStaticDataReportA()

		expected := &StaticDataReportA{
			MessageType:     24,
			RepeatIndicator: 0,
			MMSI:            503513400,
			PartNumber:      0,
			VesselName:      "WHITEBIRDS",
		}
		Context("The get should return a type 24 A message", func() {
			It("Where the message is not nil", func() {
				Expect(type24).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type24).To(Equal(expected))
		})
	})
})

var _ = Describe("Type24BMessageProcessing", func() {
	Describe("When processing a type 24 B message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,H3`fKe4T>1F93?0@3pipp01@4320,0*77",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type24, err := message.GetAsStaticDataReportB()

		expected := &StaticDataReportB{
			MessageType:          24,
			RepeatIndicator:      0,
			MMSI:                 244030388,
			PartNumber:           1,
			ShipType:             "Sailing",
			VendorID:             "NAV",
			UnitModelCode:        2,
			SerialNumber:         275392,
			CallSign:             "PC8188",
			DimensionToBow:       10,
			DimensionToStern:     4,
			DimensionToPort:      3,
			DimensionToStarboard: 2,
			MothershipMMSI:       0,
		}
		Context("The get should return a type 24 B message", func() {
			It("Where the message is not nil", func() {
				Expect(type24).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type24).To(Equal(expected))
		})
	})
})

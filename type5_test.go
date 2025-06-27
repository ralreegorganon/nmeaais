package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type5MessageProcessing", func() {
	Describe("When processing a type 5 message", func() {
		raws := []string{
			"!AIVDM,2,1,9,B,55OER>01sWpeL@GS?CM0th5:1=@u8n222222220P1PJ354AB0;PCPj3lPAiH,0*1B",
			"!AIVDM,2,2,9,B,88888888880,2*2E",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type5, err := message.GetAsStaticAndVoyageRelatedData()

		expected := &StaticAndVoyageRelatedData{
			MessageType:          5,
			RepeatIndicator:      0,
			MMSI:                 368403000,
			AISVersion:           0,
			IMONumber:            8101771,
			CallSign:             "WDE8347",
			VesselName:           "POLAR STORM",
			ShipType:             "Towing: length exceeds 200m or breadth exceeds 25m",
			DimensionToBow:       12,
			DimensionToStern:     26,
			DimensionToPort:      3,
			DimensionToStarboard: 5,
			EPFDType:             "GPS",
			ETAMonth:             1,
			ETADay:               2,
			ETAHour:              18,
			ETAMinute:            0,
			Draught:              4.6,
			Destination:          "ANCHORAGE",
			DTE:                  false,
		}
		Context("The get should return a type 5 message", func() {
			It("Where the message is not nil", func() {
				Expect(type5).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type5).To(Equal(expected))
		})
	})
	Describe("When processing a type 5 message with a short payload", func() {
		raws := []string{
			"!AIVDM,1,1,,A,50000010000000000000000000000000,0*22",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type5, err := message.GetAsStaticAndVoyageRelatedData()
		It("The get should return an error", func() {
			Expect(err).To(Not(BeNil()))
		})
		It("The get should return nil for the message", func() {
			Expect(type5).To(BeNil())
		})
	})
})

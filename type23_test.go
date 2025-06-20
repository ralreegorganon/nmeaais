package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type23MessageProcessing", func() {
	Describe("When processing a type 23 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,G02:Kn01R`sn@291nj600000900,2*12",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type23, err := message.GetAsGroupAssignmentCommand()

		expected := &GroupAssignmentCommand{
			MessageType:     23,
			RepeatIndicator: 0,
			MMSI:            2268120,
			NELongitude:     2.63,
			NELatitutde:     51.07,
			SWLongitude:     1.8266666666666667,
			SWLatitude:      50.68,
			StationType:     "Regional use and inland waterways",
			ShipType:        "Not available",
			TxRxMode:        0,
			ReportInterval:  "Next Shorter Reporting Interval",
			QuietTime:       0,
		}
		Context("The get should return a type 23 message", func() {
			It("Where the message is not nil", func() {
				Expect(type23).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type23).To(Equal(expected))
		})
	})
})

package nmeaais

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type4MessageProcessing", func() {
	Describe("When processing a type 4 message", func() {
		var (
			raws     []string
			packets  []*Packet
			message  *Message
			type4    *BaseStationReport
			err      error
			expected *BaseStationReport
		)

		BeforeEach(func() {
			raws = []string{
				"!AIVDM,1,1,,B,402R3dAurtDNn0n7C@QIev100@PE,0*45",
			}

			packets = buildPackets(raws)
			message, err = Process(packets)
			type4, err = message.GetAsBaseStationReport()

			expected = &BaseStationReport{
				MessageType:      4,
				RepeatIndicator:  0,
				MMSI:             2655153,
				PositionAccuracy: false,
				TimeStamp:        time.Date(2014, time.November, 24, 20, 30, 54, 0, time.UTC),
				Longitude:        11.8214,
				Latitude:         58.37396,
				EPFDType:         "GPS",
				RAIM:             false,
				RadioStatus:      67605,
			}
		})

		Context("The get should return a type 4 message", func() {
			It("Where the message is not nil", func() {
				Expect(type4).To(Not(BeNil()))
			})
		})

		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("The fields should be populated correctly", func() {
			Expect(type4).To(Equal(expected))
		})
	})
})

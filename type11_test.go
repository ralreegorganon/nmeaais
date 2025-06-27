package nmeaais

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type11MessageProcessing", func() {
	Describe("When processing a type 11 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,;5N;BdQuw;:i5mAi:nS27jQ02000,0*3B",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type11, err := message.GetAsUTCDateResponse()

		expected := &UTCDateResponse{
			MessageType:      11,
			RepeatIndicator:  0,
			MMSI:             367186610,
			PositionAccuracy: true,
			TimeStamp:        time.Date(2015, time.December, 22, 10, 49, 5, 0, time.UTC),
			Longitude:        -149.90960833333332,
			Latitude:         61.22487,
			EPFDType:         "GPS",
			RAIM:             true,
			RadioStatus:      0,
		}
		Context("The get should return a type 11 message", func() {
			It("Where the message is not nil", func() {
				Expect(type11).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type11).To(Equal(expected))
		})
	})
})

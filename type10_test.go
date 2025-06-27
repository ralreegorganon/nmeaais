package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type10MessageProcessing", func() {
	Describe("When processing a type 10 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,:5MwvSQGRlc8,0*45",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type10, err := message.GetAsUTCDateInquiry()

		expected := &UTCDateInquiry{
			MessageType:     10,
			RepeatIndicator: 0,
			MMSI:            367001230,
			DestinationMMSI: 367186610,
		}
		Context("The get should return a type 10 message", func() {
			It("Where the message is not nil", func() {
				Expect(type10).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type10).To(Equal(expected))
		})
	})
})

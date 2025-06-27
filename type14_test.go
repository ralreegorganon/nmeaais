package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type14MessageProcessing", func() {
	Describe("When processing a type 14 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,>>M;1IM<59B1@E=@,0*5E",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type14, err := message.GetAsSafetyRelatedBroadcast()

		expected := &SafetyRelatedBroadcast{
			MessageType:     14,
			RepeatIndicator: 0,
			MMSI:            970113381,
			SequenceNumber:  0,
			DestinationMMSI: 0,
			RetransmitFlag:  false,
			Text:            "SART TEST",
		}
		Context("The get should return a type 14 message", func() {
			It("Where the message is not nil", func() {
				Expect(type14).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type14).To(Equal(expected))
		})
	})
})

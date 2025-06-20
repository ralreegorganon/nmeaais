package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type13MessageProcessing", func() {
	Describe("When processing a type 13 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,=4WCf22Gaw0`,0*5C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type13, err := message.GetAsSafetyRelatedAcknowledge()

		expected := &SafetyRelatedAcknowledge{
			MessageType:     13,
			RepeatIndicator: 0,
			MMSI:            309653000,
			MMSI1:           636091402,
			MMSI1Sequence:   0,
		}
		Context("The get should return a type 13 message", func() {
			It("Where the message is not nil", func() {
				Expect(type13).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type13).To(Equal(expected))
		})
	})
})

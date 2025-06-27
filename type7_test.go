package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type7MessageProcessing", func() {
	Describe("When processing a type 7 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,75MwQW2G`lEH,0*6C",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type7, err := message.GetAsBinaryAcknowledge()

		expected := &BinaryAcknowledge{
			MessageType:     7,
			RepeatIndicator: 0,
			MMSI:            366993820,
			MMSI1:           636014934,
			MMSI1Sequence:   0,
		}
		Context("The get should return a type 7 message", func() {
			It("Where the message is not nil", func() {
				Expect(type7).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type7).To(Equal(expected))
		})
	})
})

package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type6MessageProcessing", func() {
	Describe("When processing a type 6 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,65D7EH5DoW300400A@E=B04<d0,4*46",
		}

		packets := buildPackets(raws)
		message, _ := Process(packets)
		type6, err := message.GetAsBinaryAddressedMessage()

		expected := &BinaryAddressedMessage{
			MessageType:        6,
			RepeatIndicator:    0,
			MMSI:               356636000,
			SequenceNumber:     1,
			DestinationMMSI:    355966000,
			RetransmitFlag:     false,
			DesignatedAreaCode: 1,
			FunctionalID:       0,
			Data:               []uint8{0, 69, 5, 77, 72, 1, 12, 176},
		}
		Context("The get should return a type 6 message", func() {
			It("Where the message is not nil", func() {
				Expect(type6).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type6).To(Equal(expected))
		})
	})
})

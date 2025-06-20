package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type15MessageProcessing", func() {
	Describe("When processing a type 15 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,?h3Ovj@p>iBPD00,2*21",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type15, err := message.GetAsInterrogation()

		expected := &Interrogation{
			MessageType:        15,
			RepeatIndicator:    3,
			MMSI:               3669705,
			InterrogatedMMSI1:  235849000,
			FirstMessageType1:  5,
			FirstSlotOffset1:   0,
			SecondMessageType1: 0,
			SecondSlotOffset1:  0,
			InterrogatedMMSI2:  0,
			FirstMessageType2:  0,
			FirstSlotOffset2:   0,
		}
		Context("The get should return a type 15 message", func() {
			It("Where the message is not nil", func() {
				Expect(type15).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type15).To(Equal(expected))
		})
	})
})

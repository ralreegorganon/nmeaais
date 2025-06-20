package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type12MessageProcessing", func() {
	Describe("When processing a type 12 message", func() {
		raws := []string{
			"!AIVDM,2,1,1,A,<D62222222208:5vmEEEOPAGEso0009m5@CFb;vNnQIsW008t>AOOfbbbWp4,0*40",
			"!AIVDM,2,2,1,A,=fD:8=w0?A@,2*4C",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type12, err := message.GetAsAddressedSafetyRelated()

		expected := &AddressedSafetyRelated{
			MessageType:     12,
			RepeatIndicator: 1,
			MMSI:            274760200,
			SequenceNumber:  0,
			DestinationMMSI: 545392672,
			RetransmitFlag:  false,
			Text:            "HJE>5UUU- QWU;7",
		}
		Context("The get should return a type 12 message", func() {
			It("Where the message is not nil", func() {
				Expect(type12).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type12).To(Equal(expected))
		})
	})
})

package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type16MessageProcessing", func() {
	Describe("When processing a type 16 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,@h3OvjBGaw3h3h0000000000,0*7E",
		}

		packets := buildPackets(raws)
		message, err := Process(packets)
		type16, err := message.GetAsAssignmentModeCommand()

		expected := &AssignmentModeCommand{
			MessageType:      16,
			RepeatIndicator:  3,
			MMSI:             3669705,
			DestinationMMSI1: 636091452,
			Offset1:          60,
			Increment1:       0,
			DestinationMMSI2: 0,
			Offset2:          0,
			Increment2:       0,
		}
		Context("The get should return a type 16 message", func() {
			It("Where the message is not nil", func() {
				Expect(type16).To(Not(BeNil()))
			})
		})
		It("The get should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("The fields should be populated correctly", func() {
			Expect(type16).To(Equal(expected))
		})
	})
})

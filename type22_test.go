package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type22MessageProcessing", func() {
	Describe("When processing a type 22 message", func() {
		Context("Addressed to individual stations", func() {
			raws := []string{
				"!AIVDM,1,1,,B,FGsEEEEv15NU47?000G@8JnKKuwGFT<V0<1gg6QvmEEEOP@,2*4D",
			}

			packets := buildPackets(raws)
			message, _ := Process(packets)
			type22, err := message.GetAsChannelManagement()

			expected := &ChannelManagement{
				MessageType:     22,
				RepeatIndicator: 1,
				MMSI:            531977557,
				ChannelA:        2016,
				ChannelB:        1111,
				TxRxMode:        10,
				Power:           false,
				Addressed:       true,
				ChannelABand:    true,
				ChannelBBand:    false,
				ZoneSize:        6,
				MMSI1:           679714304,
				MMSI2:           24381547,
				NELongitude:     0,
				NELatitutde:     0,
				SWLongitude:     0,
				SWLatitude:      0,
			}
			Context("The get should return a type 22 message", func() {
				It("Where the message is not nil", func() {
					Expect(type22).To(Not(BeNil()))
				})
			})
			It("The get should not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("The fields should be populated correctly", func() {
				Expect(type22).To(Equal(expected))
			})
		})
		Context("Broadcast to an area", func() {
			raws := []string{
				"!AIVDM,1,1,,B,FGrEEEEv4h3OvjAuuQDnDo=cc8Knew?02<8iQt1vUEEEOTP,2*54",
			}

			packets := buildPackets(raws)
			message, _ := Process(packets)
			type22, err := message.GetAsChannelManagement()

			expected := &ChannelManagement{
				MessageType:     22,
				RepeatIndicator: 1,
				MMSI:            530928981,
				ChannelA:        2017,
				ChannelB:        768,
				TxRxMode:        13,
				Power:           true,
				Addressed:       false,
				ChannelABand:    true,
				ChannelBBand:    false,
				ZoneSize:        7,
				MMSI1:           0,
				MMSI2:           0,
				NELongitude:     -1.0366666666666666,
				NELatitutde:     53.74333333333333,
				SWLongitude:     36.20166666666667,
				SWLatitude:      66.39166666666667,
			}
			Context("The get should return a type 22 message", func() {
				It("Where the message is not nil", func() {
					Expect(type22).To(Not(BeNil()))
				})
			})
			It("The get should not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("The fields should be populated correctly", func() {
				Expect(type22).To(Equal(expected))
			})
		})
	})
})

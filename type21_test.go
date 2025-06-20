package nmeaais

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type21MessageProcessing", func() {
	Describe("When processing a type 21 message", func() {
		Context("Without a name extension", func() {
			raws := []string{
				"!AIVDM,1,1,,A,E>k1kFd1WWUh4W62b@1:WdhHpP0J`lV<AQ@:000003vP10,4*69",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type21, err := message.GetAsAidToNavigationReport()

			expected := &AidToNavigationReport{
				MessageType:          21,
				RepeatIndicator:      0,
				MMSI:                 993031002,
				AidType:              "Port hand Mark",
				Name:                 "COOK INLET BUOY 11",
				PositionAccuracy:     true,
				Longitude:            -149.93683333333334,
				Latitude:             61.23533333333334,
				DimensionToBow:       0,
				DimensionToStern:     0,
				DimensionToPort:      0,
				DimensionToStarboard: 0,
				EPFDType:             "Surveyed",
				UTCSecond:            61,
				OffPositionIndicator: false,
				RAIM:                 false,
				VirtualAid:           true,
				AssignedMode:         false,
				NameExtension:        "",
			}
			Context("The get should return a type 21 message", func() {
				It("Where the message is not nil", func() {
					Expect(type21).To(Not(BeNil()))
				})
			})
			It("The get should not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("The fields should be populated correctly", func() {
				Expect(type21).To(Equal(expected))
			})
		})
		Context("With a name extension", func() {
			raws := []string{
				"!AIVDM,2,1,5,B,E1mg=5J1T4W0h97aRh6ba84<h2d;W:Te=eLvH50```q,0*46",
				"!AIVDM,2,2,5,B,:D44QDlp0C1DU00,2*36",
			}

			packets := buildPackets(raws)
			message, err := Process(packets)
			type21, err := message.GetAsAidToNavigationReport()

			expected := &AidToNavigationReport{
				MessageType:          21,
				RepeatIndicator:      0,
				MMSI:                 123456789,
				AidType:              "Cardinal Mark N",
				Name:                 "CHINA ROSE MURPHY EX",
				PositionAccuracy:     false,
				Longitude:            -122.69859166666667,
				Latitude:             47.92061833333333,
				DimensionToBow:       5,
				DimensionToStern:     5,
				DimensionToPort:      5,
				DimensionToStarboard: 5,
				EPFDType:             "GPS",
				UTCSecond:            50,
				OffPositionIndicator: false,
				RAIM:                 false,
				VirtualAid:           false,
				AssignedMode:         false,
				NameExtension:        "PRESS ALERT",
			}
			Context("The get should return a type 21 message", func() {
				It("Where the message is not nil", func() {
					Expect(type21).To(Not(BeNil()))
				})
			})
			It("The get should not return an error", func() {
				Expect(err).To(BeNil())
			})
			It("The fields should be populated correctly", func() {
				Expect(type21).To(Equal(expected))
			})
		})
	})
})

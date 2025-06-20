package nmeaais

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func accumulateInput(raws []string, d *Decoder) {
	for _, raw := range raws {
		d.Input <- DecoderInput{
			Input:     raw,
			Timestamp: time.Now(),
		}
	}
	close(d.Input)
}

var _ = Describe("Decoder", func() {
	Describe("When decoding a type 1 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,15RTgt0PAso;90TKcjM8h6g208CQ,0*4A",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 1 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&PositionReportClassA{}))
		})
	})
	Describe("When decoding a type 2 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,25Cjtd0Oj;Jp7ilG7=UkKBoB0<06,0*60",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 2 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&PositionReportClassA{}))
		})
	})
	Describe("When decoding a type 3 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,33uIKN000011dcNQ==j<5`Qj059S,0*50",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 3 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&PositionReportClassA{}))
		})
	})
	Describe("When decoding a type 4 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,402R3dAurtDNn0n7C@QIev100@PE,0*45",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 4 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&BaseStationReport{}))
		})
	})
	Describe("When decoding a type 5 message", func() {
		raws := []string{
			"!AIVDM,2,1,9,B,55OER>01sWpeL@GS?CM0th5:1=@u8n222222220P1PJ354AB0;PCPj3lPAiH,0*1B",
			"!AIVDM,2,2,9,B,88888888880,2*2E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 5 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&StaticAndVoyageRelatedData{}))
		})
	})
	Describe("When decoding a type 6 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,65D7EH5DoW300400A@E=B04<d0,4*46",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 6 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&BinaryAddressedMessage{}))
		})
	})
	Describe("When decoding a type 7 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,75MwQW2G`lEH,0*6C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 7 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&BinaryAcknowledge{}))
		})
	})
	Describe("When decoding a type 8 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,85Mwqd1Kf4dldnKQ<>bW6RGmDu<6U5f1>W<LMGV85qe;dkv@rN5h,0*7D",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 8 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&BinaryBroadcastMessage{}))
		})
	})
	Describe("When decoding a type 9 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,91b76w001L163a8QIdP8O<h00PS6,0*10",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 9 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&StandardSARAircraftPositionReport{}))
		})
	})
	Describe("When decoding a type 10 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,:5MwvSQGRlc8,0*45",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 10 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&UTCDateInquiry{}))
		})
	})
	Describe("When decoding a type 11 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,;5N;BdQuw;:i5mAi:nS27jQ02000,0*3B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 11 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&UTCDateResponse{}))
		})
	})
	Describe("When decoding a type 12 message", func() {
		raws := []string{
			"!AIVDM,2,1,1,A,<D62222222208:5vmEEEOPAGEso0009m5@CFb;vNnQIsW008t>AOOfbbbWp4,0*40",
			"!AIVDM,2,2,1,A,=fD:8=w0?A@,2*4C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 12 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&AddressedSafetyRelated{}))
		})
	})
	Describe("When decoding a type 13 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,=4WCf22Gaw0`,0*5C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 13 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&SafetyRelatedAcknowledge{}))
		})
	})
	Describe("When decoding a type 14 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,>>M;1IM<59B1@E=@,0*5E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 14 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&SafetyRelatedBroadcast{}))
		})
	})
	Describe("When decoding a type 15 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,?h3Ovj@p>iBPD00,2*21",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 15 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&Interrogation{}))
		})
	})
	Describe("When decoding a type 16 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,@h3OvjBGaw3h3h0000000000,0*7E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 16 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&AssignmentModeCommand{}))
		})
	})
	Describe("When decoding a type 17 message", func() {
		raws := []string{
			"!AIVDM,2,1,5,A,A02VqLPA4I6C07h5Ed1h<OrsuBTTwS?r:C?w`?la<gno1RTRwSP9:BcurA8a,0*3A",
			"!AIVDM,2,2,5,A,:Oko02TSwu8<:Jbb,0*11",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 17 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&DGNSSBroadcastBinaryMessage{}))
		})
	})
	Describe("When decoding a type 18 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,B52Mu0@00El8HO6oJS<Igwk5kP06,0*7B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 18 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&PositionReportClassBStandard{}))
		})
	})
	Describe("When decoding a type 19 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,C5N3SRgPEnJGEBT>NhWAwwo862PaLELTBJ:V00000000S0D:R220,0*0B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 19 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&PositionReportClassBExtended{}))
		})
	})
	Describe("When decoding a type 20 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,Dh3Ovj@11N>6;HfGL00Nfp0,2*1B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 20 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&DataLinkManagementMessage{}))
		})
	})
	Describe("When decoding a type 21 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,E>k1kFd1WWUh4W62b@1:WdhHpP0J`lV<AQ@:000003vP10,4*69",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 21 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&AidToNavigationReport{}))
		})
	})
	Describe("When decoding a type 22 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,FGsEEEEv15NU47?000G@8JnKKuwGFT<V0<1gg6QvmEEEOP@,2*4D",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 22 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&ChannelManagement{}))
		})
	})
	Describe("When decoding a type 23 message", func() {
		raws := []string{
			"!AIVDM,1,1,,B,G02:Kn01R`sn@291nj600000900,2*12",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 23 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&GroupAssignmentCommand{}))
		})
	})
	Describe("When decoding a type 24 A message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,H7P<1>1LPU@D8U8A<0000000000,2*6C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 24 A message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&StaticDataReportA{}))
		})
	})
	Describe("When decoding a type 24 B message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,H3`fKe4T>1F93?0@3pipp01@4320,0*77",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 24 B message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&StaticDataReportB{}))
		})
	})
	Describe("When decoding a type 25 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,I8IRGB40QPPa0:<HP::V=gwv0l48,0*0E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 275message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&SingleSlotBinaryMessage{}))
		})
	})
	Describe("When decoding a type 27 message", func() {
		raws := []string{
			"!AIVDM,1,1,,A,Km31e<1KQ?SO4P5d,0*66",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output
		It("The decoder should return a type 27 message", func() {
			Expect(result.DecodedMessage).To(BeAssignableToTypeOf(&LongRangeAISBroadcast{}))
		})
	})
})

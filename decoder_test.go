package nmeaais

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
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

func TestDecoder(t *testing.T) {
	Convey("When decoding a type 1 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,15RTgt0PAso;90TKcjM8h6g208CQ,0*4A",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 1 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &PositionReportClassA{})
		})
	})

	/* Need a type 2 example
	Convey("When decoding a type 2 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,15RTgt0PAso;90TKcjM8h6g208CQ,0*4A",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 2 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &PositionReportClassA{})
		})
	})
	*/

	Convey("When decoding a type 3 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,33uIKN000011dcNQ==j<5`Qj059S,0*50",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 3 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &PositionReportClassA{})
		})
	})

	Convey("When decoding a type 4 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,402R3dAurtDNn0n7C@QIev100@PE,0*45",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 4 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &BaseStationReport{})
		})
	})

	Convey("When decoding a type 5 message", t, func() {
		raws := []string{
			"!AIVDM,2,1,9,B,55OER>01sWpeL@GS?CM0th5:1=@u8n222222220P1PJ354AB0;PCPj3lPAiH,0*1B",
			"!AIVDM,2,2,9,B,88888888880,2*2E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 5 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &StaticAndVoyageRelatedData{})
		})
	})

	Convey("When decoding a type 7 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,75MwQW2G`lEH,0*6C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 7 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &BinaryAcknowledge{})
		})
	})

	Convey("When decoding a type 8 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,85Mwqd1Kf4dldnKQ<>bW6RGmDu<6U5f1>W<LMGV85qe;dkv@rN5h,0*7D",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 8 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &BinaryBroadcastMessage{})
		})
	})

	Convey("When decoding a type 9 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,91b76w001L163a8QIdP8O<h00PS6,0*10",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 9 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &StandardSARAircraftPositionReport{})
		})
	})

	Convey("When decoding a type 10 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,:5MwvSQGRlc8,0*45",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 10 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &UTCDateInquiry{})
		})
	})

	Convey("When decoding a type 11 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,;5N;BdQuw;:i5mAi:nS27jQ02000,0*3B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 11 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &UTCDateResponse{})
		})
	})

	Convey("When decoding a type 15 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,?h3Ovj@p>iBPD00,2*21",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 15 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &Interrogation{})
		})
	})

	Convey("When decoding a type 16 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,@h3OvjBGaw3h3h0000000000,0*7E",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 16 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &AssignmentModeCommand{})
		})
	})

	Convey("When decoding a type 17 message", t, func() {
		raws := []string{
			"!AIVDM,2,1,5,A,A02VqLPA4I6C07h5Ed1h<OrsuBTTwS?r:C?w`?la<gno1RTRwSP9:BcurA8a,0*3A",
			"!AIVDM,2,2,5,A,:Oko02TSwu8<:Jbb,0*11",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 17 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &DGNSSBroadcastBinaryMessage{})
		})
	})

	Convey("When decoding a type 18 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,B52Mu0@00El8HO6oJS<Igwk5kP06,0*7B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 18 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &PositionReportClassBStandard{})
		})
	})

	Convey("When decoding a type 20 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,Dh3Ovj@11N>6;HfGL00Nfp0,2*1B",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 20 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &DataLinkManagementMessage{})
		})
	})

	Convey("When decoding a type 21 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,E>k1kFd1WWUh4W62b@1:WdhHpP0J`lV<AQ@:000003vP10,4*69",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 21 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &AidToNavigationReport{})
		})
	})

	Convey("When decoding a type 22 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,B,FGsEEEEv15NU47?000G@8JnKKuwGFT<V0<1gg6QvmEEEOP@,2*4D",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 22 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &ChannelManagement{})
		})
	})

	Convey("When decoding a type 24 A message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,H7P<1>1LPU@D8U8A<0000000000,2*6C",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 24 A message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &StaticDataReportA{})
		})
	})

	Convey("When decoding a type 24 B message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,H3`fKe4T>1F93?0@3pipp01@4320,0*77",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 24 B message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &StaticDataReportB{})
		})
	})

	Convey("When decoding a type 27 message", t, func() {
		raws := []string{
			"!AIVDM,1,1,,A,Km31e<1KQ?SO4P5d,0*66",
		}

		d := NewDecoder()
		accumulateInput(raws, d)
		result := <-d.Output

		Convey("The decoder should return a type 27 message", func() {
			So(result.DecodedMessage, ShouldHaveSameTypeAs, &LongRangeAISBroadcast{})
		})
	})
}

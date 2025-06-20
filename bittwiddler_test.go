package nmeaais

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NmeaMessageBittwiddling", func() {
	payload := "133m@ogP00PD;88MD5MTDww@2D7k"
	Describe("When unarmoring a payload", func() {
		unarmored, bitLength := unarmor([]byte(payload))
		It("The payload should unarmor correctly", func() {
			value := fmt.Sprintf("%x", unarmored)
			expected := "0430f5437be00008142c821d50576453ffd00941f300000000000000"
			Expect(value).To(Equal(expected))
			Expect(bitLength).To(Equal(int64(168)))
		})
	})
	Describe("When extracting unsigned ints from an unarmored payload", func() {
		unarmored, _ := unarmor([]byte(payload))

		values := []uint{
			asUInt(unarmored, 0, 6),
			asUInt(unarmored, 6, 2),
			asUInt(unarmored, 8, 30),
			asUInt(unarmored, 38, 4),
			asUInt(unarmored, 50, 10),
			asUInt(unarmored, 60, 1),
			asUInt(unarmored, 116, 12),
			asUInt(unarmored, 128, 9),
			asUInt(unarmored, 137, 6),
			asUInt(unarmored, 143, 2),
			asUInt(unarmored, 148, 1),
		}
		It("The extracted values should be correct", func() {
			expected := []uint{
				uint(1),
				uint(0),
				uint(205344990),
				uint(15),
				uint(0),
				uint(1),
				uint(1107),
				uint(511),
				uint(40),
				uint(0),
				uint(1),
			}
			Expect(values).To(Equal(expected))
		})
	})
	Describe("When extracting signed ints from an unarmored payload", func() {
		unarmored, _ := unarmor([]byte(payload))

		values := []int{
			asInt(unarmored, 42, 8),
			asInt(unarmored, 61, 28),
			asInt(unarmored, 89, 27),
		}
		It("The extracted values should be correct", func() {
			expected := []int{
				-128,
				2644228,
				30737782,
			}
			Expect(values).To(Equal(expected))
		})
	})
	Describe("When extracting a string from an unarmored payload", func() {
		unarmored, _ := unarmor([]byte("H42O55i18tMET00000000000000"))
		values := []string{
			asString(unarmored, 40, 120),
		}
		It("The extracted values should be correct", func() {
			expected := []string{
				"PROGUY",
			}
			Expect(values).To(Equal(expected))
		})
	})
	Describe("When extracting a binary from an unarmored payload", func() {
		unarmored, bitLength := unarmor([]byte("85Mwqd1Kf4dldnKQ<>bW6RGmDu<6U5f1>W<LMGV85qe;dkv@rN5h"))
		values := [][]byte{
			asBinary(unarmored, 56, uint(bitLength-56)),
		}
		It("The extracted values should be correct", func() {
			value := fmt.Sprintf("%x", values[0])
			expected := "4b34b366e130eaa71a25f553d306945b813a731c757988179b4bb33f90e9e170"
			Expect(value).To(Equal(expected))
		})
	})
})

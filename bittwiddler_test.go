package nmeaais

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNmeaMessageBittwiddling(t *testing.T) {
	payload := "133m@ogP00PD;88MD5MTDww@2D7k"

	Convey("When unarmoring a payload", t, func() {
		unarmored := unarmor([]byte(payload))
		Convey("The payload should unarmor correctly", func() {
			value := fmt.Sprintf("%x", unarmored)
			expected := "0430f5437be00008142c821d50576453ffd00941f300000000000000"
			So(value, ShouldEqual, expected)
		})
	})

	Convey("When extracting unsigned ints from an unarmored payload", t, func() {
		unarmored := unarmor([]byte(payload))

		values := []uint{
			extractUnsignedInt(unarmored, 0, 6),
			extractUnsignedInt(unarmored, 6, 2),
			extractUnsignedInt(unarmored, 8, 30),
			extractUnsignedInt(unarmored, 38, 4),
			extractUnsignedInt(unarmored, 50, 10),
			extractUnsignedInt(unarmored, 60, 1),
			extractUnsignedInt(unarmored, 61, 28),
			extractUnsignedInt(unarmored, 89, 27),
			extractUnsignedInt(unarmored, 116, 12),
			extractUnsignedInt(unarmored, 128, 9),
			extractUnsignedInt(unarmored, 137, 6),
			extractUnsignedInt(unarmored, 143, 2),
			extractUnsignedInt(unarmored, 148, 1),
		}
		Convey("The extracted values should be correct", func() {
			expected := []uint{
				uint(1),
				uint(0),
				uint(205344990),
				uint(15),
				uint(0),
				uint(1),
				uint(2644228),
				uint(30737782),
				uint(1107),
				uint(511),
				uint(40),
				uint(0),
				uint(1),
			}
			So(values, ShouldResemble, expected)
		})
	})

	Convey("When extracting signed ints from an unarmored payload", t, func() {
		unarmored := unarmor([]byte(payload))

		values := []int{
			extractSignedInt(unarmored, 42, 8),
		}
		Convey("The extracted values should be correct", func() {
			expected := []int{
				-128,
			}
			So(values, ShouldResemble, expected)
		})
	})

}

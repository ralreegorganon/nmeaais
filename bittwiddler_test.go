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
		Convey("The extracted values should be correct", func() {
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
			So(values, ShouldResemble, expected)
		})
	})

	Convey("When extracting signed ints from an unarmored payload", t, func() {
		unarmored := unarmor([]byte(payload))

		values := []int{
			asInt(unarmored, 42, 8),
			asInt(unarmored, 61, 28),
			asInt(unarmored, 89, 27),
		}
		Convey("The extracted values should be correct", func() {
			expected := []int{
				-128,
				2644228,
				30737782,
			}
			So(values, ShouldResemble, expected)
		})
	})
}

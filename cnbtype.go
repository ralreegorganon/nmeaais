package nmeaais

import (
	"fmt"
	"math"
)

type PositionReportClassA struct {
	MessageType       int64
	RepeatIndicator   int64
	MMSI              int64
	NavigationStatus  string
	RateOfTurn        float64
	SpeedOverGround   float64
	PositionAccuracy  bool
	Longitude         float64
	Latitude          float64
	CourseOverGround  float64
	TrueHeading       int64
	TimeStamp         int64
	ManeuverIndicator string
	RAIM              bool
	RadioStatus       int64
}

func (m *Message) GetAsPositionReportClassA() (*PositionReportClassA, error) {

	switch m.MessageType {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		break
	default:
		return nil, fmt.Errorf("nmeaais: tried to get message as type 1, 2, or 3, but is type %v", m.MessageType)
	}

	p := &PositionReportClassA{
		MessageType:       m.MessageType,
		RepeatIndicator:   m.RepeatIndicator,
		MMSI:              m.MMSI,
		NavigationStatus:  navigationStatus(asUInt(m.unarmoredPayload, 38, 4)),
		RateOfTurn:        rateOfTurn(asInt(m.unarmoredPayload, 42, 8)),
		SpeedOverGround:   speedOverGround(asUInt(m.unarmoredPayload, 50, 10)),
		PositionAccuracy:  asBool(asUInt(m.unarmoredPayload, 60, 1)),
		Longitude:         latlon(asInt(m.unarmoredPayload, 61, 28)),
		Latitude:          latlon(asInt(m.unarmoredPayload, 89, 27)),
		CourseOverGround:  courseOverGround(asUInt(m.unarmoredPayload, 116, 12)),
		TrueHeading:       int64(asUInt(m.unarmoredPayload, 128, 9)),
		TimeStamp:         int64(asUInt(m.unarmoredPayload, 137, 6)),
		ManeuverIndicator: maneuverIndicator(asUInt(m.unarmoredPayload, 143, 2)),
		RAIM:              asBool(asUInt(m.unarmoredPayload, 148, 1)),
		RadioStatus:       int64(asUInt(m.unarmoredPayload, 149, 19)),
	}
	return p, nil
}

var navigationStatuses = []string{
	"Under way using engine",
	"At anchor",
	"Not under command",
	"Restricted manoeuverability",
	"Constrained by her draught",
	"Moored",
	"Aground",
	"Engaged in Fishing",
	"Under way sailing",
	"Reserved for future amendment of Navigational Status for HSC",
	"Reserved for future amendment of Navigational Status for WIG",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"AIS-SART is active",
	"Not defined",
}
var navigationStatusesMax = uint(len(navigationStatuses) - 1)

func navigationStatus(ns uint) string {
	if ns < 0 || ns > navigationStatusesMax {
		return "Not defined"
	}
	return navigationStatuses[ns]
}

func rateOfTurn(rot int) float64 {
	if rot == 128 {
		return math.NaN()
	}
	if rot == 127 || rot == -127 {
		return math.Inf(rot)
	}
	floatified := float64(rot)
	value := floatified / 4.733
	value *= value
	return math.Copysign(value, floatified)
}

func speedOverGround(sog uint) float64 {
	return float64(sog) / 10
}

func courseOverGround(cog uint) float64 {
	return float64(cog) / 10
}

func maneuverIndicator(mi uint) string {
	switch mi {
	case 0:
		return "Not available"
	case 1:
		return "No special maneuver"
	case 2:
		return "Special maneuver"
	default:
		return "Not available"
	}
}

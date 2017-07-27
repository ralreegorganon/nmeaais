package nmeaais

import (
	"fmt"
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

func (m *Message) GetAsPositionReportClassA() (p *PositionReportClassA, err error) {
	defer func() {
		if r := recover(); r != nil {
			p = nil
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

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

	var expectedMinimumLength int64 = 168
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &PositionReportClassA{
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

	return
}

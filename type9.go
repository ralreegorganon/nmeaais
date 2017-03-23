package nmeaais

import (
	"fmt"
)

type StandardSARAircraftPositionReport struct {
	MessageType      int64
	RepeatIndicator  int64
	MMSI             int64
	Altitude         int64
	SpeedOverGround  float64
	PositionAccuracy bool
	Longitude        float64
	Latitude         float64
	CourseOverGround float64
	TimeStamp        int64
	DTE              bool
	Assigned         bool
	RAIM             bool
	RadioStatus      int64
}

func (m *Message) GetAsStandardSARAircraftPositionReport() (p *StandardSARAircraftPositionReport, err error) {
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

	var validMessageType int64 = 9

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &StandardSARAircraftPositionReport{
		MessageType:      m.MessageType,
		RepeatIndicator:  m.RepeatIndicator,
		MMSI:             m.MMSI,
		Altitude:         int64(asUInt(m.unarmoredPayload, 38, 12)),
		SpeedOverGround:  speedOverGroundAircraft(asUInt(m.unarmoredPayload, 50, 10)),
		PositionAccuracy: asBool(asUInt(m.unarmoredPayload, 60, 1)),
		Longitude:        latlon(asInt(m.unarmoredPayload, 61, 28)),
		Latitude:         latlon(asInt(m.unarmoredPayload, 89, 27)),
		CourseOverGround: courseOverGround(asUInt(m.unarmoredPayload, 116, 12)),
		TimeStamp:        int64(asUInt(m.unarmoredPayload, 128, 6)),
		DTE:              asBool(asUInt(m.unarmoredPayload, 142, 1)),
		Assigned:         asBool(asUInt(m.unarmoredPayload, 146, 1)),
		RAIM:             asBool(asUInt(m.unarmoredPayload, 147, 1)),
		RadioStatus:      int64(asUInt(m.unarmoredPayload, 148, 20)),
	}

	return
}

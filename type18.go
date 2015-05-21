package nmeaais

import (
	"fmt"
)

type PositionReportClassBStandard struct {
	MessageType      int64
	RepeatIndicator  int64
	MMSI             int64
	SpeedOverGround  float64
	PositionAccuracy bool
	Longitude        float64
	Latitude         float64
	CourseOverGround float64
	TrueHeading      int64
	TimeStamp        int64
	CSUnit           bool
	Display          bool
	DSC              bool
	Band             bool
	Message22        bool
	Assigned         bool
	RAIM             bool
	RadioStatus      int64
}

func (m *Message) GetAsPositionReportClassBStandard() (p *PositionReportClassBStandard, err error) {
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

	var validMessageType int64 = 18

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &PositionReportClassBStandard{
		MessageType:      m.MessageType,
		RepeatIndicator:  m.RepeatIndicator,
		MMSI:             m.MMSI,
		SpeedOverGround:  speedOverGround(asUInt(m.unarmoredPayload, 46, 10)),
		PositionAccuracy: asBool(asUInt(m.unarmoredPayload, 56, 1)),
		Longitude:        latlon(asInt(m.unarmoredPayload, 57, 28)),
		Latitude:         latlon(asInt(m.unarmoredPayload, 85, 27)),
		CourseOverGround: courseOverGround(asUInt(m.unarmoredPayload, 112, 12)),
		TrueHeading:      int64(asUInt(m.unarmoredPayload, 124, 9)),
		TimeStamp:        int64(asUInt(m.unarmoredPayload, 133, 6)),
		CSUnit:           asBool(asUInt(m.unarmoredPayload, 141, 1)),
		Display:          asBool(asUInt(m.unarmoredPayload, 142, 1)),
		DSC:              asBool(asUInt(m.unarmoredPayload, 143, 1)),
		Band:             asBool(asUInt(m.unarmoredPayload, 144, 1)),
		Message22:        asBool(asUInt(m.unarmoredPayload, 145, 1)),
		Assigned:         asBool(asUInt(m.unarmoredPayload, 146, 1)),
		RAIM:             asBool(asUInt(m.unarmoredPayload, 147, 1)),
		RadioStatus:      int64(asUInt(m.unarmoredPayload, 148, 20)),
	}

	return
}

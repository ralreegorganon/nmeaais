package nmeaais

import (
	"fmt"
)

type LongRangeAISBroadcast struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	PositionAccuracy   bool
	RAIM               bool
	NavigationStatus   string
	Longitude          float64
	Latitude           float64
	SpeedOverGround    float64
	CourseOverGround   float64
	GNSSPositionStatus int64
}

func (m *Message) GetAsLongRangeAISBroadcast() (p *LongRangeAISBroadcast, err error) {
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

	var validMessageType int64 = 27

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &LongRangeAISBroadcast{
		MessageType:        m.MessageType,
		RepeatIndicator:    m.RepeatIndicator,
		MMSI:               m.MMSI,
		PositionAccuracy:   asBool(asUInt(m.unarmoredPayload, 38, 1)),
		RAIM:               asBool(asUInt(m.unarmoredPayload, 39, 1)),
		NavigationStatus:   navigationStatus(asUInt(m.unarmoredPayload, 40, 4)),
		Longitude:          latlonShort(asInt(m.unarmoredPayload, 44, 18)),
		Latitude:           latlonShort(asInt(m.unarmoredPayload, 62, 17)),
		SpeedOverGround:    speedOverGroundLongRange(asUInt(m.unarmoredPayload, 79, 6)),
		CourseOverGround:   courseOverGroundLongRange(asUInt(m.unarmoredPayload, 85, 9)),
		GNSSPositionStatus: int64(asUInt(m.unarmoredPayload, 94, 1)),
	}

	return
}

package nmeaais

import (
	"fmt"
	"time"
)

type BaseStationReport struct {
	MessageType      int64
	RepeatIndicator  int64
	MMSI             int64
	TimeStamp        time.Time
	PositionAccuracy bool
	Longitude        float64
	Latitude         float64
	EPFDType         string
	RAIM             bool
	RadioStatus      int64
}

func (m *Message) GetAsBaseStationReport() (p *BaseStationReport, err error) {
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

	var validMessageType int64 = 4

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type 4, but is type %v", m.MessageType)
	}

	p = &BaseStationReport{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		TimeStamp: time.Date(
			int(asUInt(m.unarmoredPayload, 38, 14)),
			time.Month(asUInt(m.unarmoredPayload, 52, 4)),
			int(asUInt(m.unarmoredPayload, 56, 5)),
			int(asUInt(m.unarmoredPayload, 61, 5)),
			int(asUInt(m.unarmoredPayload, 66, 6)),
			int(asUInt(m.unarmoredPayload, 72, 6)),
			0,
			time.UTC),
		PositionAccuracy: asBool(asUInt(m.unarmoredPayload, 78, 1)),
		Longitude:        latlon(asInt(m.unarmoredPayload, 79, 28)),
		Latitude:         latlon(asInt(m.unarmoredPayload, 107, 27)),
		EPFDType:         epfdType(asUInt(m.unarmoredPayload, 134, 4)),
		RAIM:             asBool(asUInt(m.unarmoredPayload, 148, 1)),
		RadioStatus:      int64(asUInt(m.unarmoredPayload, 149, 19)),
	}
	return
}

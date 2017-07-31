package nmeaais

import "fmt"

type PositionReportClassBExtended struct {
	MessageType          int64
	RepeatIndicator      int64
	MMSI                 int64
	SpeedOverGround      float64
	PositionAccuracy     bool
	Longitude            float64
	Latitude             float64
	CourseOverGround     float64
	TrueHeading          int64
	TimeStamp            int64
	VesselName           string
	ShipType             string
	DimensionToBow       int64
	DimensionToStern     int64
	DimensionToPort      int64
	DimensionToStarboard int64
	EPFDType             string
	RAIM                 bool
	DTE                  bool
	Assigned             bool
}

func (m *Message) GetAsPositionReportClassBExtended() (p *PositionReportClassBExtended, err error) {
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

	var validMessageType int64 = 19

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var expectedMinimumLength int64 = 312
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected at least %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &PositionReportClassBExtended{
		MessageType:          m.MessageType,
		RepeatIndicator:      m.RepeatIndicator,
		MMSI:                 m.MMSI,
		SpeedOverGround:      speedOverGround(asUInt(m.unarmoredPayload, 46, 10)),
		PositionAccuracy:     asBool(asUInt(m.unarmoredPayload, 56, 1)),
		Longitude:            latlon(asInt(m.unarmoredPayload, 57, 28)),
		Latitude:             latlon(asInt(m.unarmoredPayload, 85, 27)),
		CourseOverGround:     courseOverGround(asUInt(m.unarmoredPayload, 112, 12)),
		TrueHeading:          int64(asUInt(m.unarmoredPayload, 124, 9)),
		TimeStamp:            int64(asUInt(m.unarmoredPayload, 133, 6)),
		VesselName:           asString(m.unarmoredPayload, 143, 120),
		ShipType:             shipType(asUInt(m.unarmoredPayload, 263, 8)),
		DimensionToBow:       int64(asUInt(m.unarmoredPayload, 271, 9)),
		DimensionToStern:     int64(asUInt(m.unarmoredPayload, 280, 9)),
		DimensionToPort:      int64(asUInt(m.unarmoredPayload, 289, 6)),
		DimensionToStarboard: int64(asUInt(m.unarmoredPayload, 295, 6)),
		EPFDType:             epfdType(asUInt(m.unarmoredPayload, 301, 4)),
		RAIM:                 asBool(asUInt(m.unarmoredPayload, 305, 1)),
		Assigned:             asBool(asUInt(m.unarmoredPayload, 306, 1)),
		DTE:                  asBool(asUInt(m.unarmoredPayload, 307, 1)),
	}
	return
}

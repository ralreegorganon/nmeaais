package nmeaais

import "fmt"

type AidToNavigationReport struct {
	MessageType          int64
	RepeatIndicator      int64
	MMSI                 int64
	AidType              string
	Name                 string
	PositionAccuracy     bool
	Longitude            float64
	Latitude             float64
	DimensionToBow       int64
	DimensionToStern     int64
	DimensionToPort      int64
	DimensionToStarboard int64
	EPFDType             string
	UTCSecond            int64
	OffPositionIndicator bool
	RAIM                 bool
	VirtualAid           bool
	AssignedMode         bool
	NameExtension        string
}

func (m *Message) GetAsAidToNavigationReport() (p *AidToNavigationReport, err error) {
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

	var validMessageType int64 = 21

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var expectedMinimumLength int64 = 272
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &AidToNavigationReport{
		MessageType:          m.MessageType,
		RepeatIndicator:      m.RepeatIndicator,
		MMSI:                 m.MMSI,
		AidType:              aidType(asUInt(m.unarmoredPayload, 38, 5)),
		Name:                 asString(m.unarmoredPayload, 43, 120),
		PositionAccuracy:     asBool(asUInt(m.unarmoredPayload, 163, 1)),
		Longitude:            latlon(asInt(m.unarmoredPayload, 164, 28)),
		Latitude:             latlon(asInt(m.unarmoredPayload, 192, 27)),
		DimensionToBow:       int64(asUInt(m.unarmoredPayload, 219, 9)),
		DimensionToStern:     int64(asUInt(m.unarmoredPayload, 228, 9)),
		DimensionToPort:      int64(asUInt(m.unarmoredPayload, 237, 6)),
		DimensionToStarboard: int64(asUInt(m.unarmoredPayload, 243, 6)),
		EPFDType:             epfdType(asUInt(m.unarmoredPayload, 249, 4)),
		UTCSecond:            int64(asUInt(m.unarmoredPayload, 253, 6)),
		OffPositionIndicator: asBool(asUInt(m.unarmoredPayload, 259, 1)),
		RAIM:                 asBool(asUInt(m.unarmoredPayload, 268, 1)),
		VirtualAid:           asBool(asUInt(m.unarmoredPayload, 269, 1)),
		AssignedMode:         asBool(asUInt(m.unarmoredPayload, 270, 1)),
		NameExtension:        asString(m.unarmoredPayload, 272, 88),
	}

	return
}

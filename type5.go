package nmeaais

import "fmt"

type StaticAndVoyageRelatedData struct {
	MessageType          int64
	RepeatIndicator      int64
	MMSI                 int64
	AISVersion           int64
	IMONumber            int64
	CallSign             string
	VesselName           string
	ShipType             string
	DimensionToBow       int64
	DimensionToStern     int64
	DimensionToPort      int64
	DimensionToStarboard int64
	EPFDType             string
	ETAMonth             int64
	ETADay               int64
	ETAHour              int64
	ETAMinute            int64
	Draught              float64
	Destination          string
	DTE                  bool
}

func (m *Message) GetAsStaticAndVoyageRelatedData() (p *StaticAndVoyageRelatedData, err error) {
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

	var validMessageType int64 = 5

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var expectedMinimumLength int64 = 420
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected at least %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &StaticAndVoyageRelatedData{
		MessageType:          m.MessageType,
		RepeatIndicator:      m.RepeatIndicator,
		MMSI:                 m.MMSI,
		AISVersion:           int64(asUInt(m.unarmoredPayload, 38, 2)),
		IMONumber:            int64(asUInt(m.unarmoredPayload, 40, 30)),
		CallSign:             asString(m.unarmoredPayload, 70, 42),
		VesselName:           asString(m.unarmoredPayload, 112, 120),
		ShipType:             shipType(asUInt(m.unarmoredPayload, 232, 8)),
		DimensionToBow:       int64(asUInt(m.unarmoredPayload, 240, 9)),
		DimensionToStern:     int64(asUInt(m.unarmoredPayload, 249, 9)),
		DimensionToPort:      int64(asUInt(m.unarmoredPayload, 258, 6)),
		DimensionToStarboard: int64(asUInt(m.unarmoredPayload, 264, 6)),
		EPFDType:             epfdType(asUInt(m.unarmoredPayload, 270, 4)),
		ETAMonth:             int64(asUInt(m.unarmoredPayload, 274, 4)),
		ETADay:               int64(asUInt(m.unarmoredPayload, 278, 5)),
		ETAHour:              int64(asUInt(m.unarmoredPayload, 283, 5)),
		ETAMinute:            int64(asUInt(m.unarmoredPayload, 288, 6)),
		Draught:              draught(asUInt(m.unarmoredPayload, 294, 8)),
		Destination:          asString(m.unarmoredPayload, 302, 120),
		DTE:                  asBool(asUInt(m.unarmoredPayload, 422, 1)),
	}
	return
}

func draught(d uint) float64 {
	return float64(d) / 10
}

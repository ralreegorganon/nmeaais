package nmeaais

import "fmt"

type StaticDataReportA struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	PartNumber      int64
	VesselName      string
}

func (m *Message) IsStaticDataReportA() (bool, error) {
	var validMessageType int64 = 24
	if m.MessageType != validMessageType {
		return false, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	partNumber := asUInt(m.unarmoredPayload, 38, 2)

	switch partNumber {
	case 0:
		return true, nil
	case 1:
		return false, nil
	default:
		return false, fmt.Errorf("nmeaais: type 24 part type of %v is invalid", partNumber)
	}
}

func (m *Message) GetAsStaticDataReportA() (p *StaticDataReportA, err error) {
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

	var validMessageType int64 = 24

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &StaticDataReportA{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		PartNumber:      int64(asUInt(m.unarmoredPayload, 38, 2)),
		VesselName:      asString(m.unarmoredPayload, 40, 120),
	}
	return
}

type StaticDataReportB struct {
	MessageType          int64
	RepeatIndicator      int64
	MMSI                 int64
	PartNumber           int64
	ShipType             string
	VendorID             string
	UnitModelCode        int64
	SerialNumber         int64
	CallSign             string
	DimensionToBow       int64
	DimensionToStern     int64
	DimensionToPort      int64
	DimensionToStarboard int64
	MothershipMMSI       int64
}

func (m *Message) IsStaticDataReportB() (bool, error) {
	var validMessageType int64 = 24
	if m.MessageType != validMessageType {
		return false, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	partNumber := asUInt(m.unarmoredPayload, 38, 2)

	switch partNumber {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("nmeaais: type 24 part type of %v is invalid", partNumber)
	}
}

func (m *Message) GetAsStaticDataReportB() (p *StaticDataReportB, err error) {
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

	var validMessageType int64 = 24

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &StaticDataReportB{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		PartNumber:      int64(asUInt(m.unarmoredPayload, 38, 2)),
		ShipType:        shipType(asUInt(m.unarmoredPayload, 40, 8)),
		VendorID:        asString(m.unarmoredPayload, 48, 18),
		UnitModelCode:   int64(asUInt(m.unarmoredPayload, 66, 4)),
		SerialNumber:    int64(asUInt(m.unarmoredPayload, 70, 20)),
		CallSign:        asString(m.unarmoredPayload, 90, 42),
	}

	isAux := m.MMSI >= 980000000
	if isAux {
		p.MothershipMMSI = int64(asUInt(m.unarmoredPayload, 132, 30))
	} else {
		p.DimensionToBow = int64(asUInt(m.unarmoredPayload, 132, 9))
		p.DimensionToStern = int64(asUInt(m.unarmoredPayload, 141, 9))
		p.DimensionToPort = int64(asUInt(m.unarmoredPayload, 150, 6))
		p.DimensionToStarboard = int64(asUInt(m.unarmoredPayload, 156, 6))
	}
	return
}

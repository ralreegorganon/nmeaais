package nmeaais

import "fmt"

type GroupAssignmentCommand struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	NELongitude     float64
	NELatitutde     float64
	SWLongitude     float64
	SWLatitude      float64
	StationType     string
	ShipType        string
	TxRxMode        int64
	ReportInterval  string
	QuietTime       int64
}

func (m *Message) GetAsGroupAssignmentCommand() (p *GroupAssignmentCommand, err error) {
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

	var validMessageType int64 = 23

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &GroupAssignmentCommand{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		NELongitude:     latlonShort(asInt(m.unarmoredPayload, 40, 18)),
		NELatitutde:     latlonShort(asInt(m.unarmoredPayload, 58, 17)),
		SWLongitude:     latlonShort(asInt(m.unarmoredPayload, 75, 18)),
		SWLatitude:      latlonShort(asInt(m.unarmoredPayload, 93, 17)),
		StationType:     stationType(asUInt(m.unarmoredPayload, 110, 4)),
		ShipType:        shipType(asUInt(m.unarmoredPayload, 114, 8)),
		TxRxMode:        int64(asUInt(m.unarmoredPayload, 144, 2)),
		ReportInterval:  stationInterval(asUInt(m.unarmoredPayload, 146, 4)),
		QuietTime:       int64(asUInt(m.unarmoredPayload, 150, 4)),
	}

	return
}

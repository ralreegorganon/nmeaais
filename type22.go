package nmeaais

import "fmt"

type ChannelManagement struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	ChannelA        int64
	ChannelB        int64
	TxRxMode        int64
	Power           bool
	NELongitude     float64
	NELatitutde     float64
	SWLongitude     float64
	SWLatitude      float64
	MMSI1           int64
	MMSI2           int64
	Addressed       bool
	ChannelABand    bool
	ChannelBBand    bool
	ZoneSize        int64
}

func (m *Message) GetAsChannelManagement() (p *ChannelManagement, err error) {
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

	var validMessageType int64 = 22

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &ChannelManagement{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		ChannelA:        int64(asUInt(m.unarmoredPayload, 40, 12)),
		ChannelB:        int64(asUInt(m.unarmoredPayload, 52, 12)),
		TxRxMode:        int64(asUInt(m.unarmoredPayload, 64, 4)),
		Power:           asBool(asUInt(m.unarmoredPayload, 68, 1)),
		Addressed:       asBool(asUInt(m.unarmoredPayload, 139, 1)),
		ChannelABand:    asBool(asUInt(m.unarmoredPayload, 140, 1)),
		ChannelBBand:    asBool(asUInt(m.unarmoredPayload, 141, 1)),
		ZoneSize:        int64(asUInt(m.unarmoredPayload, 142, 3)),
	}

	if p.Addressed {
		p.MMSI1 = int64(asUInt(m.unarmoredPayload, 69, 30))
		p.MMSI2 = int64(asUInt(m.unarmoredPayload, 104, 30))
	} else {
		p.NELongitude = latlonShort(asInt(m.unarmoredPayload, 69, 18))
		p.NELatitutde = latlonShort(asInt(m.unarmoredPayload, 87, 17))
		p.SWLongitude = latlonShort(asInt(m.unarmoredPayload, 104, 18))
		p.SWLatitude = latlonShort(asInt(m.unarmoredPayload, 122, 17))
	}

	return
}

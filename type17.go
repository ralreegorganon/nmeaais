package nmeaais

import "fmt"

type DGNSSBroadcastBinaryMessage struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	Longitude       float64
	Latitude        float64
	Payload         []byte
}

func (m *Message) GetAsDGNSSBroadcastBinaryMessage() (p *DGNSSBroadcastBinaryMessage, err error) {
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

	var validMessageType int64 = 17

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &DGNSSBroadcastBinaryMessage{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		Longitude:       latlonShort(asInt(m.unarmoredPayload, 40, 18)),
		Latitude:        latlonShort(asInt(m.unarmoredPayload, 58, 17)),
		Payload:         asBinary(m.unarmoredPayload, 80, uint(m.bitLength-80)),
	}

	return
}

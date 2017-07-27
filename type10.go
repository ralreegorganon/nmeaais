package nmeaais

import (
	"fmt"
)

type UTCDateInquiry struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	DestinationMMSI int64
}

func (m *Message) GetAsUTCDateInquiry() (p *UTCDateInquiry, err error) {
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

	var validMessageType int64 = 10

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &UTCDateInquiry{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		DestinationMMSI: int64(asUInt(m.unarmoredPayload, 40, 30)),
	}

	return
}

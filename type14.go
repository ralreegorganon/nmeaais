package nmeaais

import "fmt"

type SafetyRelatedBroadcast struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	SequenceNumber  int64
	DestinationMMSI int64
	RetransmitFlag  bool
	Text            string
}

func (m *Message) GetAsSafetyRelatedBroadcast() (p *SafetyRelatedBroadcast, err error) {
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

	var validMessageType int64 = 14

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &SafetyRelatedBroadcast{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		Text:            asString(m.unarmoredPayload, 40, 968),
	}

	return
}

package nmeaais

import "fmt"

type AddressedSafetyRelated struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	SequenceNumber  int64
	DestinationMMSI int64
	RetransmitFlag  bool
	Text            string
}

func (m *Message) GetAsAddressedSafetyRelated() (p *AddressedSafetyRelated, err error) {
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

	var validMessageType int64 = 12

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &AddressedSafetyRelated{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		SequenceNumber:  int64(asUInt(m.unarmoredPayload, 38, 2)),
		DestinationMMSI: int64(asUInt(m.unarmoredPayload, 40, 30)),
		RetransmitFlag:  asBool(asUInt(m.unarmoredPayload, 70, 1)),
		Text:            asString(m.unarmoredPayload, 72, 936),
	}

	return
}

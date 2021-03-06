package nmeaais

import "fmt"

type DataLinkManagementMessage struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	Offset1         int64
	ReservedSlots1  int64
	Timeout1        int64
	Increment1      int64
	Offset2         int64
	ReservedSlots2  int64
	Timeout2        int64
	Increment2      int64
	Offset3         int64
	ReservedSlots3  int64
	Timeout3        int64
	Increment3      int64
	Offset4         int64
	ReservedSlots4  int64
	Timeout4        int64
	Increment4      int64
}

func (m *Message) GetAsDataLinkManagementMessage() (p *DataLinkManagementMessage, err error) {
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

	var validMessageType int64 = 20

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var expectedMinimumLength int64 = 72
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &DataLinkManagementMessage{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		Offset1:         int64(asUInt(m.unarmoredPayload, 40, 12)),
		ReservedSlots1:  int64(asUInt(m.unarmoredPayload, 52, 4)),
		Timeout1:        int64(asUInt(m.unarmoredPayload, 56, 3)),
		Increment1:      int64(asUInt(m.unarmoredPayload, 59, 11)),
	}

	if m.bitLength >= 100 {
		p.Offset2 = int64(asUInt(m.unarmoredPayload, 70, 12))
		p.ReservedSlots2 = int64(asUInt(m.unarmoredPayload, 82, 4))
		p.Timeout2 = int64(asUInt(m.unarmoredPayload, 86, 3))
		p.Increment2 = int64(asUInt(m.unarmoredPayload, 89, 11))
	}

	if m.bitLength >= 130 {
		p.Offset3 = int64(asUInt(m.unarmoredPayload, 100, 12))
		p.ReservedSlots3 = int64(asUInt(m.unarmoredPayload, 112, 4))
		p.Timeout3 = int64(asUInt(m.unarmoredPayload, 116, 3))
		p.Increment3 = int64(asUInt(m.unarmoredPayload, 119, 11))
	}

	if m.bitLength >= 160 {
		p.Offset4 = int64(asUInt(m.unarmoredPayload, 130, 12))
		p.ReservedSlots4 = int64(asUInt(m.unarmoredPayload, 142, 4))
		p.Timeout4 = int64(asUInt(m.unarmoredPayload, 146, 3))
		p.Increment4 = int64(asUInt(m.unarmoredPayload, 149, 11))
	}

	return
}

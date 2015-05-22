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

	p = &DataLinkManagementMessage{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		Offset1:         int64(asUInt(m.unarmoredPayload, 40, 12)),
		ReservedSlots1:  int64(asUInt(m.unarmoredPayload, 52, 4)),
		Timeout1:        int64(asUInt(m.unarmoredPayload, 56, 3)),
		Increment1:      int64(asUInt(m.unarmoredPayload, 59, 11)),
		Offset2:         int64(asUInt(m.unarmoredPayload, 70, 12)),
		ReservedSlots2:  int64(asUInt(m.unarmoredPayload, 82, 4)),
		Timeout2:        int64(asUInt(m.unarmoredPayload, 86, 3)),
		Increment2:      int64(asUInt(m.unarmoredPayload, 89, 11)),
		Offset3:         int64(asUInt(m.unarmoredPayload, 100, 12)),
		ReservedSlots3:  int64(asUInt(m.unarmoredPayload, 112, 4)),
		Timeout3:        int64(asUInt(m.unarmoredPayload, 116, 3)),
		Increment3:      int64(asUInt(m.unarmoredPayload, 119, 11)),
		Offset4:         int64(asUInt(m.unarmoredPayload, 130, 12)),
		ReservedSlots4:  int64(asUInt(m.unarmoredPayload, 142, 4)),
		Timeout4:        int64(asUInt(m.unarmoredPayload, 146, 3)),
		Increment4:      int64(asUInt(m.unarmoredPayload, 149, 11)),
	}
	return
}

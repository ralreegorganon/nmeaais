package nmeaais

import "fmt"

type AssignmentModeCommand struct {
	MessageType      int64
	RepeatIndicator  int64
	MMSI             int64
	DestinationMMSI1 int64
	Offset1          int64
	Increment1       int64
	DestinationMMSI2 int64
	Offset2          int64
	Increment2       int64
}

func (m *Message) GetAsAssignmentModeCommand() (p *AssignmentModeCommand, err error) {
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

	var validMessageType int64 = 16

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &AssignmentModeCommand{
		MessageType:      m.MessageType,
		RepeatIndicator:  m.RepeatIndicator,
		MMSI:             m.MMSI,
		DestinationMMSI1: int64(asUInt(m.unarmoredPayload, 40, 30)),
		Offset1:          int64(asUInt(m.unarmoredPayload, 70, 12)),
		Increment1:       int64(asUInt(m.unarmoredPayload, 82, 10)),
	}
	if m.bitLength >= 144 {
		p.DestinationMMSI2 = int64(asUInt(m.unarmoredPayload, 92, 30))
		p.Offset2 = int64(asUInt(m.unarmoredPayload, 122, 12))
		p.Increment2 = int64(asUInt(m.unarmoredPayload, 134, 10))
	}

	return
}

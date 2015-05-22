package nmeaais

import "fmt"

type Interrogation struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	InterrogatedMMSI1  int64
	FirstMessageType1  int64
	FirstSlotOffset1   int64
	SecondMessageType1 int64
	SecondSlotOffset1  int64
	InterrogatedMMSI2  int64
	FirstMessageType2  int64
	FirstSlotOffset2   int64
}

func (m *Message) GetAsInterrogation() (p *Interrogation, err error) {
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

	var validMessageType int64 = 15

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &Interrogation{
		MessageType:       m.MessageType,
		RepeatIndicator:   m.RepeatIndicator,
		MMSI:              m.MMSI,
		InterrogatedMMSI1: int64(asUInt(m.unarmoredPayload, 40, 30)),
		FirstMessageType1: int64(asUInt(m.unarmoredPayload, 70, 6)),
		FirstSlotOffset1:  int64(asUInt(m.unarmoredPayload, 76, 12)),
	}
	if m.bitLength >= 110 {
		p.SecondMessageType1 = int64(asUInt(m.unarmoredPayload, 90, 6))
		p.SecondSlotOffset1 = int64(asUInt(m.unarmoredPayload, 96, 12))
	}
	if m.bitLength >= 160 {
		p.FirstMessageType2 = int64(asUInt(m.unarmoredPayload, 140, 6))
		p.FirstSlotOffset2 = int64(asUInt(m.unarmoredPayload, 146, 12))
	}

	return
}

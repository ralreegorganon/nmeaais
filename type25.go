package nmeaais

import "fmt"

type SingleSlotBinaryMessage struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	Addressed          bool
	Structured         bool
	DestinationMMSI    int64
	DesignatedAreaCode int64
	FunctionalID       int64
	Data               []byte
}

func (m *Message) GetAsSingleSlotBinaryMessage() (p *SingleSlotBinaryMessage, err error) {
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

	var validMessageType int64 = 25

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &SingleSlotBinaryMessage{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		Addressed:       asBool(asUInt(m.unarmoredPayload, 38, 1)),
		Structured:      asBool(asUInt(m.unarmoredPayload, 39, 1)),
	}

	if p.Addressed && p.Structured {
		p.DestinationMMSI = int64(asUInt(m.unarmoredPayload, 40, 30))
		p.DesignatedAreaCode = int64(asUInt(m.unarmoredPayload, 70, 10))
		p.FunctionalID = int64(asUInt(m.unarmoredPayload, 80, 6))
		p.Data = asBinary(m.unarmoredPayload, 86, uint(m.bitLength-86))
	} else if p.Addressed && !p.Structured {
		p.DestinationMMSI = int64(asUInt(m.unarmoredPayload, 40, 30))
		p.Data = asBinary(m.unarmoredPayload, 70, uint(m.bitLength-70))
	} else if !p.Addressed && p.Structured {
		p.DesignatedAreaCode = int64(asUInt(m.unarmoredPayload, 40, 10))
		p.FunctionalID = int64(asUInt(m.unarmoredPayload, 50, 6))
		p.Data = asBinary(m.unarmoredPayload, 56, uint(m.bitLength-56))
	} else if !p.Addressed && !p.Structured {
		p.Data = asBinary(m.unarmoredPayload, 40, uint(m.bitLength-40))
	}

	return
}

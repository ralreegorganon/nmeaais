package nmeaais

import "fmt"

type BinaryBroadcastMessage struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	DesignatedAreaCode int64
	FunctionalID       int64
	Data               []byte
}

func (m *Message) GetAsBinaryBroadcastMessage() (p *BinaryBroadcastMessage, err error) {
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

	var validMessageType int64 = 8

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var expectedMinimumLength int64 = 56
	if m.bitLength < expectedMinimumLength {
		return nil, fmt.Errorf("nmeaais: type %v message payload has insufficient length of %v, expected %v", m.MessageType, m.bitLength, expectedMinimumLength)
	}

	p = &BinaryBroadcastMessage{
		MessageType:        m.MessageType,
		RepeatIndicator:    m.RepeatIndicator,
		MMSI:               m.MMSI,
		DesignatedAreaCode: int64(asUInt(m.unarmoredPayload, 40, 10)),
		FunctionalID:       int64(asUInt(m.unarmoredPayload, 50, 6)),
		Data:               asBinary(m.unarmoredPayload, 56, uint(m.bitLength-56)),
	}

	return
}

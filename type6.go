package nmeaais

import "fmt"

type BinaryAddressedMessage struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	SequenceNumber     int64
	DestinationMMSI    int64
	RetransmitFlag     bool
	DesignatedAreaCode int64
	FunctionalID       int64
	Data               []byte
}

func (m *Message) GetAsBinaryAddressedMessage() (p *BinaryAddressedMessage, err error) {
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

	var validMessageType int64 = 6

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &BinaryAddressedMessage{
		MessageType:        m.MessageType,
		RepeatIndicator:    m.RepeatIndicator,
		MMSI:               m.MMSI,
		SequenceNumber:     int64(asUInt(m.unarmoredPayload, 38, 2)),
		DestinationMMSI:    int64(asUInt(m.unarmoredPayload, 40, 30)),
		RetransmitFlag:     asBool(asUInt(m.unarmoredPayload, 70, 1)),
		DesignatedAreaCode: int64(asUInt(m.unarmoredPayload, 72, 10)),
		FunctionalID:       int64(asUInt(m.unarmoredPayload, 82, 6)),
		Data:               asBinary(m.unarmoredPayload, 88, uint(m.bitLength-88)),
	}

	return
}

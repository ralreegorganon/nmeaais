package nmeaais

import "fmt"

type BinaryAcknowledge struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	MMSI1           int64
	MMSI1Sequence   int64
	MMSI2           int64
	MMSI2Sequence   int64
	MMSI3           int64
	MMSI3Sequence   int64
	MMSI4           int64
	MMSI4Sequence   int64
}

func (m *Message) GetAsBinaryAcknowledge() (p *BinaryAcknowledge, err error) {
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

	var validMessageType int64 = 7

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &BinaryAcknowledge{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		MMSI1:           int64(asUInt(m.unarmoredPayload, 40, 30)),
		MMSI1Sequence:   int64(asUInt(m.unarmoredPayload, 70, 2)),
	}

	if m.bitLength >= 104 {
		p.MMSI2 = int64(asUInt(m.unarmoredPayload, 72, 30))
		p.MMSI2Sequence = int64(asUInt(m.unarmoredPayload, 102, 2))
	}

	if m.bitLength >= 136 {
		p.MMSI3 = int64(asUInt(m.unarmoredPayload, 104, 30))
		p.MMSI3Sequence = int64(asUInt(m.unarmoredPayload, 134, 2))
	}

	if m.bitLength >= 168 {
		p.MMSI4 = int64(asUInt(m.unarmoredPayload, 136, 30))
		p.MMSI4Sequence = int64(asUInt(m.unarmoredPayload, 166, 2))
	}

	return
}

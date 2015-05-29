package nmeaais

import (
	"errors"
	"fmt"
)

type Message struct {
	Packets          []*Packet
	unarmoredPayload []byte
	bitLength        int64
	MessageType      int64
	RepeatIndicator  int64
	MMSI             int64
}

func Process(packets []*Packet) (*Message, error) {
	message := &Message{
		Packets: packets,
	}

	err := message.validateMultipart()
	if err != nil {
		return nil, err
	}

	err = message.unarmorPayload()
	if err != nil {
		return nil, err
	}

	if len(message.unarmoredPayload) == 0 {
		return nil, fmt.Errorf("nmeaais: message has a zero-length payload")
	}

	message.MessageType = int64(asUInt(message.unarmoredPayload, 0, 6))
	message.RepeatIndicator = int64(asUInt(message.unarmoredPayload, 6, 2))
	message.MMSI = int64(asUInt(message.unarmoredPayload, 8, 30))

	return message, nil
}

func (m *Message) validateMultipart() error {
	c := int64(len(m.Packets))
	uniqueSequences := make(map[int64]bool)
	for i, p := range m.Packets {
		if p.FragmentCount != c {
			return fmt.Errorf("nmeaais: message has %v packets, expected %v", c, p.FragmentCount)
		}
		if int64(i+1) != p.FragmentNumber {
			return errors.New("nmeaais: message packet out sequence")
		}

		uniqueSequences[p.SequentialMessageID] = true
	}

	if len(uniqueSequences) > 1 {
		return errors.New("nmeaais: message contains packets from multiple messages")
	}

	return nil
}

func (m *Message) unarmorPayload() error {
	complete := ""

	for _, p := range m.Packets {
		complete += p.Payload
	}

	completeBytes := []byte(complete)
	m.unarmoredPayload, m.bitLength = unarmor(completeBytes)

	return nil
}

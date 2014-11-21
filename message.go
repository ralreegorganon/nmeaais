package nmeaais

import (
	"errors"
	"fmt"
)

type Message struct {
	Packets          []*Packet
	unarmoredPayload []byte
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

		uniqueSequences[p.SequentialMessageId] = true
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
	m.unarmoredPayload = unarmor(completeBytes)

	return nil
}

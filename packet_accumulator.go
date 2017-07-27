package nmeaais

import (
	"time"
)

type PacketAccumulatorResult struct {
	Packets   []*Packet
	Message   *Message
	Error     error
	Timestamp time.Time
}

type PacketAccumulator struct {
	Packets chan *Packet
	Results chan *PacketAccumulatorResult
}

func NewPacketAccumulator() *PacketAccumulator {
	pa := &PacketAccumulator{
		Packets: make(chan *Packet),
		Results: make(chan *PacketAccumulatorResult),
	}

	go pa.process()

	return pa
}

func (pa *PacketAccumulator) process() {
	packetBuffer := make(map[int64]map[int64][]*Packet)
	for p := range pa.Packets {
		if p.FragmentCount == 1 {
			packets := []*Packet{p}
			pa.processAccumulatedPackets(packets)
		} else {
			fragsForSeq, ok := packetBuffer[p.SequentialMessageID]

			if !ok {
				fragsForSeq = make(map[int64][]*Packet)
				packetBuffer[p.SequentialMessageID] = fragsForSeq
			}

			instancesForFrag, ok := fragsForSeq[p.FragmentNumber]

			if !ok {
				instancesForFrag = make([]*Packet, 0)
				fragsForSeq[p.FragmentNumber] = instancesForFrag
			}

			fragsForSeq[p.FragmentNumber] = append(packetBuffer[p.SequentialMessageID][p.FragmentNumber], p)

			composed := make([]*Packet, 0)
			for i := int64(1); i <= p.FragmentCount; i++ {
				if ap, ok := packetBuffer[p.SequentialMessageID][int64(i)]; ok {
					l := len(ap)
					if l > 0 {
						composed = append(composed, ap[0])
					}
				}
			}

			composedFragmentCount := int64(len(composed))
			if p.FragmentCount == composedFragmentCount {
				pa.processAccumulatedPackets(composed)
				for i := int64(1); i <= p.FragmentCount; i++ {
					packetBuffer[p.SequentialMessageID][i] = packetBuffer[p.SequentialMessageID][i][1:]
				}
			}
		}
	}
	close(pa.Results)
}

func (pa *PacketAccumulator) processAccumulatedPackets(p []*Packet) {
	m, err := Process(p)
	pa.Results <- &PacketAccumulatorResult{
		Packets:   p,
		Message:   m,
		Error:     err,
		Timestamp: p[0].Timestamp,
	}
}

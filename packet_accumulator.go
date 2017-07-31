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

			if _, ok := packetBuffer[p.SequentialMessageID]; !ok {
				fragsForSeq := make(map[int64][]*Packet)
				packetBuffer[p.SequentialMessageID] = fragsForSeq
			}

			if _, ok := packetBuffer[p.SequentialMessageID][p.FragmentNumber]; !ok {
				instancesForFrag := make([]*Packet, 0)
				packetBuffer[p.SequentialMessageID][p.FragmentNumber] = instancesForFrag
			}

			maxFragInterval := time.Duration(0)
			for _, v := range packetBuffer[p.SequentialMessageID] {
				for _, x := range v {
					since := p.Timestamp.Sub(x.Timestamp)
					if since > maxFragInterval {
						maxFragInterval = since
					}
				}
			}

			if maxFragInterval > time.Duration(2)*time.Second {
				packetBuffer[p.SequentialMessageID] = make(map[int64][]*Packet)
				packetBuffer[p.SequentialMessageID][p.FragmentNumber] = make([]*Packet, 0)
			}

			packetBuffer[p.SequentialMessageID][p.FragmentNumber] = append(packetBuffer[p.SequentialMessageID][p.FragmentNumber], p)

			composed := make([]*Packet, 0)
			for i := int64(1); i <= p.FragmentCount; i++ {
				if ap, ok := packetBuffer[p.SequentialMessageID][i]; ok {
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

package nmeaais

import (
	"sort"
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

type packets []*Packet

func (slice packets) Len() int {
	return len(slice)
}

func (slice packets) Less(i, j int) bool {
	return slice[i].FragmentNumber < slice[j].FragmentNumber
}

func (slice packets) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (pa *PacketAccumulator) process() {
	packetBuffer := make(map[int64][]*Packet)
	for p := range pa.Packets {
		if p.FragmentCount == 1 {
			packets := []*Packet{p}
			pa.processAccumulatedPackets(packets)
		} else {
			packetBuffer[p.SequentialMessageID] = append(packetBuffer[p.SequentialMessageID], p)
			bufferedFragmentCount := int64(len(packetBuffer[p.SequentialMessageID]))
			if p.FragmentCount == bufferedFragmentCount {
				sort.Sort(packets(packetBuffer[p.SequentialMessageID]))
				pa.processAccumulatedPackets(packetBuffer[p.SequentialMessageID])
				delete(packetBuffer, p.SequentialMessageID)
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

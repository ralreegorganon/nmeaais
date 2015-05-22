package nmeaais

import "strings"

func unarmor(payload []byte) ([]byte, int64) {
	var bitLength uint
	unarmored := make([]byte, len(payload))

	for _, s := range payload {
		s = s - 48
		if s > 40 {
			s = s - 8
		}

		for i := 5; i >= 0; i-- {
			if (s>>byte(i))&1 == 1 {
				unarmored[bitLength/8] |= (1 << (7 - bitLength%8))
			}
			bitLength++
		}
	}

	return unarmored, int64(bitLength)
}

const bpb uint = 8

func asUInt(unarmoredPayload []byte, start uint, width uint) uint {
	data := uint(0)
	startIndex := start / bpb
	endIndex := (start + width + bpb - 1) / bpb
	for i := startIndex; i < endIndex; i++ {
		data = data << bpb
		data = data | uint(unarmoredPayload[i])
	}
	end := (start + width) % bpb
	if end != 0 {
		data = data >> (bpb - end)
	}
	data = data & ^(255 << width)
	return data
}

func asInt(unarmoredPayload []byte, start uint, width uint) int {
	data := asUInt(unarmoredPayload, start, width)
	if (data>>(width-1))&1 == 1 {
		return -int(pow(2, width) - data)
	}

	return int(data)
}

func pow(a, b uint) uint {
	p := uint(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

const stringChars string = "@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^- !\"#$%&'()*+,-./0123456789:;<=>?"

func asString(unarmoredPayload []byte, start uint, width uint) string {
	value := ""
	for i := uint(0); i < width/6; i++ {
		position := asUInt(unarmoredPayload, start+6*i, 6)
		messageChar := stringChars[position]
		if string(messageChar) == "@" {
			break
		} else {
			value += string(messageChar)
		}
	}

	value = strings.TrimSpace(strings.Replace(value, "@", " ", -1))
	return value
}

func asBinary(unarmoredPayload []byte, start uint, width uint) []byte {
	var data []byte
	for i := uint(0); i < width/8; i++ {
		data = append(data, byte(asUInt(unarmoredPayload, start+8*i, 8)))
	}
	return data
}

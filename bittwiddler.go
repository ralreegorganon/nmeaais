package nmeaais

func unarmor(payload []byte) []byte {
	var bitLength uint = 0
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
			bitLength += 1
		}
	}

	return unarmored
}

const bpb uint = 8

func extractUnsignedInt(unarmoredPayload []byte, start uint, width uint) uint {
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

func extractSignedInt(unarmoredPayload []byte, start uint, width uint) int {
	data := extractUnsignedInt(unarmoredPayload, start, width)
	if (((data >> width) - 1) & 1) == 1 {
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

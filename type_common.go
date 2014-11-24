package nmeaais

func asBool(b uint) bool {
	if b == 1 {
		return true
	}

	return false
}

func latlon(l int) float64 {
	return float64(l) / 600000
}

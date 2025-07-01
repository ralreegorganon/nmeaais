package nmeaais

import (
	"testing"
	"time"
)

func FuzzMessageProcess(f *testing.F) {
	// Create valid packet seeds for fuzzing
	validPackets := []string{
		"!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76",
		"!AIVDM,1,1,,B,403OviQuMGCqWrRO9>E6fE700@GO,0*4D",
		"!AIVDM,1,1,,A,85Mwp`1Kf3aCnsNvBWLi=wQuNhA5t43N`5nCuI=p<IBfVqnMgPGs,0*40",
		"!AIVDM,2,1,3,A,55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8,0*66",
		"!AIVDM,2,2,3,A,88888888880,2*23",
	}

	for _, packetStr := range validPackets {
		f.Add(packetStr)
	}

	f.Fuzz(func(t *testing.T, data string) {
		if len(data) == 0 || len(data) > 2000 {
			return
		}

		// Parse packet first
		packet, err := Parse(data)
		if err != nil {
			return
		}

		// Test message processing - should not panic
		packets := []*Packet{packet}
		message, err := Process(packets)

		if err == nil && message != nil {
			// Validate basic message fields
			if message.MessageType < 0 || message.MessageType > 27 {
				t.Errorf("Invalid message type: %d", message.MessageType)
			}

			if message.RepeatIndicator < 0 || message.RepeatIndicator > 3 {
				t.Errorf("Invalid repeat indicator: %d", message.RepeatIndicator)
			}

			if message.MMSI < 0 || message.MMSI > 999999999 {
				t.Errorf("Invalid MMSI: %d", message.MMSI)
			}
		}
	})
}

func FuzzMessageMultipart(f *testing.F) {
	// Test multipart message handling
	f.Add("!AIVDM,2,1,3,A,55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8,0*66",
		"!AIVDM,2,2,3,A,88888888880,2*23")

	f.Fuzz(func(t *testing.T, packet1, packet2 string) {
		if len(packet1) == 0 || len(packet2) == 0 || len(packet1) > 1000 || len(packet2) > 1000 {
			return
		}

		// Parse both packets
		p1, err1 := Parse(packet1)
		p2, err2 := Parse(packet2)

		if err1 != nil || err2 != nil {
			return
		}

		// Test multipart processing - should not panic
		packets := []*Packet{p1, p2}
		_, _ = Process(packets)
	})
}

func FuzzDecoderInput(f *testing.F) {
	// Test the decoder with various inputs
	validInputs := []string{
		"!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76",
		"!AIVDM,1,1,,B,403OviQuMGCqWrRO9>E6fE700@GO,0*4D",
		"!BSVDM,1,1,,A,85Mwp`1Kf3aCnsNvBWLi=wQuNhA5t43N`5nCuI=p<IBfVqnMgPGs,0*40",
	}

	for _, input := range validInputs {
		f.Add(input, time.Now().Unix())
	}

	f.Fuzz(func(t *testing.T, data string, timeUnix int64) {
		if len(data) == 0 || len(data) > 1000 {
			return
		}

		timestamp := time.Unix(timeUnix, 0)

		// Create decoder
		decoder := NewDecoder()
		defer func() {
			close(decoder.Input)
		}()

		// Send input - should not panic
		select {
		case decoder.Input <- DecoderInput{Input: data, Timestamp: timestamp}:
		default:
			// Channel full, skip
		}

		// Try to read output with timeout
		select {
		case output := <-decoder.Output:
			// Validate output if successful
			if output.Error == nil && output.DecodedMessage != nil {
				if output.SourceMessage.MessageType < 1 || output.SourceMessage.MessageType > 27 {
					t.Errorf("Invalid decoded message type: %d", output.SourceMessage.MessageType)
				}
			}
		case <-time.After(100 * time.Millisecond):
			// Timeout, which is fine for fuzz testing
		}
	})
}

func FuzzBitTwiddling(f *testing.F) {
	// Test the bit manipulation functions used in message parsing
	f.Add([]byte{0x15, 0xAB, 0xCD, 0xEF}, uint(0), uint(8))
	f.Add([]byte{0xFF, 0x00, 0xAA, 0x55}, uint(4), uint(12))
	f.Add([]byte{0x01, 0x23, 0x45, 0x67, 0x89}, uint(1), uint(6))

	f.Fuzz(func(t *testing.T, data []byte, startBit, bitLength uint) {
		if len(data) == 0 || len(data) > 100 {
			return
		}

		// Bounds check to prevent panic
		if startBit >= uint(len(data)*8) {
			return
		}

		if startBit+bitLength > uint(len(data)*8) {
			return
		}

		// Should not panic
		result := asUInt(data, startBit, bitLength)

		// Basic validation - result should fit in the bit length
		if bitLength <= 63 && bitLength > 0 && result >= (1<<bitLength) {
			t.Errorf("Result %d doesn't fit in %d bits", result, bitLength)
		}
	})
}

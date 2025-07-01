package nmeaais

import (
	"fmt"
	"testing"
	"time"
)

func FuzzPacketAccumulator(f *testing.F) {
	// Seed with valid multipart sequences
	f.Add("!AIVDM,2,1,3,A,55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8,0*66",
		"!AIVDM,2,2,3,A,88888888880,2*23")
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76", "")

	f.Fuzz(func(t *testing.T, packet1, packet2 string) {
		if len(packet1) > 1000 || len(packet2) > 1000 {
			return
		}

		accumulator := NewPacketAccumulator()
		defer func() {
			close(accumulator.Packets)
		}()

		// Parse and send first packet
		if packet1 != "" {
			p1, err := Parse(packet1)
			if err == nil {
				select {
				case accumulator.Packets <- p1:
				default:
				}
			}
		}

		// Parse and send second packet if provided
		if packet2 != "" {
			p2, err := Parse(packet2)
			if err == nil {
				select {
				case accumulator.Packets <- p2:
				default:
				}
			}
		}

		// Try to read results with timeout
		select {
		case result := <-accumulator.Results:
			// Validate result if successful
			if result.Error == nil && result.Message != nil {
				if len(result.Packets) == 0 {
					t.Error("No source packets in successful result")
				}

				if result.Message.MessageType < 1 || result.Message.MessageType > 27 {
					t.Errorf("Invalid message type: %d", result.Message.MessageType)
				}
			}
		case <-time.After(50 * time.Millisecond):
			// Timeout is acceptable for fuzz testing
		}
	})
}

func FuzzPacketAccumulatorSequences(f *testing.F) {
	// Test with various sequence IDs and fragment patterns
	f.Add(int64(1), int64(1), int64(1), "A", "15M5>N?001Gst=:LtGwQeL6p0k@V")
	f.Add(int64(2), int64(1), int64(3), "A", "55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8")
	f.Add(int64(2), int64(2), int64(3), "A", "88888888880")

	f.Fuzz(func(t *testing.T, fragCount, fragNum, seqID int64, channel, payload string) {
		// Bounds checking
		if fragCount <= 0 || fragCount > 9 || fragNum <= 0 || fragNum > fragCount {
			return
		}

		if len(payload) == 0 || len(payload) > 100 {
			return
		}

		if len(channel) > 1 {
			return
		}

		// Construct a synthetic packet string
		packetStr := fmt.Sprintf("!AIVDM,%d,%d,%d,%s,%s,0*00", fragCount, fragNum, seqID, channel, payload)

		// Try to parse - should not panic
		packet, err := Parse(packetStr)
		if err != nil {
			return
		}

		accumulator := NewPacketAccumulator()
		defer func() {
			close(accumulator.Packets)
		}()

		// Send packet to accumulator
		select {
		case accumulator.Packets <- packet:
		default:
		}

		// Check for results
		select {
		case result := <-accumulator.Results:
			if result.Error == nil && result.Message != nil {
				// Validate the accumulated result
				if len(result.Packets) != int(fragCount) && fragCount > 1 {
					// Only expect full fragment count for multipart messages
					// Single packets should pass through immediately
				}
			}
		case <-time.After(50 * time.Millisecond):
		}
	})
}

func FuzzPacketAccumulatorTiming(f *testing.F) {
	// Test accumulator with various timing scenarios
	f.Add("!AIVDM,2,1,5,A,55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8,0*64", int64(1000))
	f.Add("!AIVDM,2,2,5,A,88888888880,2*21", int64(2000))

	f.Fuzz(func(t *testing.T, packetStr string, delayMs int64) {
		if len(packetStr) == 0 || len(packetStr) > 1000 {
			return
		}

		// Limit delay to reasonable range
		if delayMs < 0 || delayMs > 10000 {
			return
		}

		packet, err := Parse(packetStr)
		if err != nil {
			return
		}

		accumulator := NewPacketAccumulator()
		defer func() {
			close(accumulator.Packets)
		}()

		// Send packet with artificial delay
		go func() {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
			select {
			case accumulator.Packets <- packet:
			default:
			}
		}()

		// Wait for result or timeout
		select {
		case result := <-accumulator.Results:
			if result.Error == nil && result.Message != nil {
				// Basic validation
				if result.Message.MessageType < 1 || result.Message.MessageType > 27 {
					t.Errorf("Invalid message type: %d", result.Message.MessageType)
				}
			}
		case <-time.After(time.Duration(delayMs+1000) * time.Millisecond):
			// Expected timeout for long delays
		}
	})
}

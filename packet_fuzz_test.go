package nmeaais

import (
	"strings"
	"testing"
	"time"
)

func FuzzPacketParse(f *testing.F) {
	// Seed with valid test cases
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76")
	f.Add("!AIVDM,1,1,,B,15M5>N?001Gst=:LtGwQeL6p0k@V,0*75")
	f.Add("!AIVDM,2,1,3,A,55?MbV02;H;s<HtKR20EHE:0@T4@Dn2222222216L961O5Gf0NSQEp6ClRp8,0*66")
	f.Add("!AIVDM,2,2,3,A,88888888880,2*23")
	f.Add("!BSVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76")

	// Edge cases
	f.Add("!AIVDM,1,1,,A,,0*56")                                // Empty payload
	f.Add("!AIVDM,1,1,9,A,15M5>N?001Gst=:LtGwQeL6p0k@V,5*70")   // Max fill bits
	f.Add("!AIVDM,1,1,999,1,15M5>N?001Gst=:LtGwQeL6p0k@V,0*4F") // Large sequence ID

	// Malformed cases for robustness
	f.Add("AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76") // Missing start delimiter
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0")   // Missing checksum delimiter
	f.Add("!AIVDM,1,1,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76") // Wrong field count

	f.Fuzz(func(t *testing.T, data string) {
		// Don't test empty strings or extremely long strings to avoid DoS
		if len(data) == 0 || len(data) > 1000 {
			return
		}

		// Parse the packet - should not panic
		packet, err := Parse(data)

		// If parsing succeeded, validate the result
		if err == nil && packet != nil {
			// Basic sanity checks on successful parses
			if packet.FragmentCount <= 0 || packet.FragmentCount > 9 {
				t.Errorf("Invalid fragment count: %d", packet.FragmentCount)
			}

			if packet.FragmentNumber <= 0 || packet.FragmentNumber > 9 {
				t.Errorf("Invalid fragment number: %d", packet.FragmentNumber)
			}

			if packet.FragmentNumber > packet.FragmentCount {
				t.Errorf("Fragment number %d exceeds count %d", packet.FragmentNumber, packet.FragmentCount)
			}

			if packet.FillBits < 0 || packet.FillBits > 5 {
				t.Errorf("Invalid fill bits: %d", packet.FillBits)
			}

			if packet.Tag != "AIVDM" && packet.Tag != "BSVDM" {
				t.Errorf("Invalid tag: %s", packet.Tag)
			}

			if packet.RadioChannel != "" &&
				packet.RadioChannel != "A" && packet.RadioChannel != "B" &&
				packet.RadioChannel != "1" && packet.RadioChannel != "2" {
				t.Errorf("Invalid radio channel: %s", packet.RadioChannel)
			}
		}
	})
}

func FuzzPacketParseAtTime(f *testing.F) {
	testTime := time.Now()

	// Seed with valid cases
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76", testTime.Unix())
	f.Add("!BSVDM,1,1,,B,15M5>N?001Gst=:LtGwQeL6p0k@V,0*75", testTime.Unix())

	f.Fuzz(func(t *testing.T, data string, timeUnix int64) {
		if len(data) == 0 || len(data) > 1000 {
			return
		}

		timestamp := time.Unix(timeUnix, 0)

		// Should not panic
		packet, err := ParseAtTime(data, timestamp)

		if err == nil && packet != nil {
			// Verify timestamp was set correctly
			if !packet.Timestamp.Equal(timestamp) {
				t.Errorf("Timestamp not set correctly: expected %v, got %v", timestamp, packet.Timestamp)
			}
		}
	})
}

func FuzzPacketValidation(f *testing.F) {
	// Test checksum validation specifically
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*76") // Valid checksum
	f.Add("!AIVDM,1,1,,A,15M5>N?001Gst=:LtGwQeL6p0k@V,0*77") // Invalid checksum

	f.Fuzz(func(t *testing.T, rawData string) {
		if len(rawData) == 0 || len(rawData) > 1000 {
			return
		}

		// Try to construct packet parts manually to test validation
		if !strings.Contains(rawData, "*") {
			return
		}

		parts := strings.Split(rawData, "*")
		if len(parts) != 2 {
			return
		}

		// Should not panic during checksum validation
		_, _ = Parse(rawData)
	})
}

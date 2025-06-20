# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library and toolkit for processing NMEA AIS (Automatic Identification System) data - maritime navigation safety messages used by ships and vessel traffic services. The project decodes binary AIS messages that contain vessel position, speed, course, and identification information.

## Common Commands

**Build and Test:**
```bash
# Run all tests with Ginkgo
ginkgo ./...

# Run tests with verbose output
ginkgo -v ./...

# Run specific test suite
ginkgo ./type4_test.go

# Traditional Go test (also works)
go test ./...

# Run tests with verbose output
go test -v ./...

# Build all executables
go build ./cmd/...

# Build specific tools
go build ./cmd/nmeaais-decoder
go build ./cmd/nmeaais-mock-client  
go build ./cmd/nmeaais-mock-listener

# Clean up dependencies
go mod tidy
```

**Running the Tools:**
```bash
# Main AIS decoder (connects to TCP source and decodes AIS data)
./cmd/nmeaais-decoder/nmeaais-decoder -source localhost:32779 -debug

# Mock client (sends test NMEA data to remote TCP endpoint)
go run ./cmd/nmeaais-mock-client/main.go

# Mock listener (serves NMEA data from files on TCP port)
go run ./cmd/nmeaais-mock-listener/main.go
```

## Architecture

**Core Processing Pipeline:**
1. **Raw NMEA Input** → `packet.go` (parses NMEA sentences)
2. **Packet Accumulation** → `packet_accumulator.go` (handles multi-part messages)
3. **Message Decoding** → `decoder.go` + `message.go` (converts to structured data)
4. **Type-Specific Processing** → `type4.go` through `type27.go` (AIS message types)

**Key Components:**
- `bittwiddler.go`: Binary data extraction from AIS payloads
- `type_common.go`: Shared functionality across message types
- Each `typeX.go` file implements a specific AIS message type decoder
- `cmd/` executables provide TCP connectivity and testing tools

**Message Flow:**
Raw NMEA sentences can span multiple packets for complex messages. The accumulator collects related packets before passing complete messages to type-specific decoders that extract vessel data using bit manipulation.

## Testing Framework

Uses Ginkgo testing framework with Gomega assertions. Tests are comprehensive with 29 test files covering all core functionality. The project includes test data files and mock implementations for development.

**Test Structure:**
- BDD-style tests using `Describe`, `Context`, and `It` blocks
- Expectations use Gomega matchers like `Expect(actual).To(Equal(expected))`
- Setup code handled in `BeforeEach` blocks for proper test isolation

## Dependencies

- `github.com/sirupsen/logrus`: Structured logging throughout the application
- `github.com/onsi/ginkgo/v2`: BDD-style testing framework  
- `github.com/onsi/gomega`: Matcher/assertion library for Ginkgo tests
- `github.com/davecgh/go-spew`: Debug pretty printing for complex data structures

## Development Notes

- Module path: `github.com/ralreegorganon/nmeaais`
- Go version: 1.18+
- This project is currently experimental (note: "Not for public consumption yet" in README)
- Reference links to AIS standards and test data sources are in `reference.md`
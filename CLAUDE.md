# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library and toolkit for processing NMEA AIS (Automatic Identification System) data - maritime navigation safety messages used by ships and vessel traffic services. The project decodes binary AIS messages that contain vessel position, speed, course, and identification information.

## Common Commands

**Build and Test:**

```bash
# Traditional Go test
go test ./...

# Run tests with verbose output
go test -v ./...

# Run fuzz tests (comprehensive suite)
./run_fuzz_tests.sh

# Run specific fuzz test
go test -fuzz=FuzzPacketParse -fuzztime=30s

# Build all executables
go build ./cmd/...

# Build specific tools
go build ./cmd/nmeaais-decoder
go build ./cmd/nmeaais-mock-client
go build ./cmd/nmeaais-mock-listener

# Clean up dependencies
go mod tidy

# Format code (automated via git hook)
go fmt ./...
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
4. **Type-Specific Processing** → `type4.go` through `type27.go` + `typecnb.go` (AIS message types)

**Key Components:**

- `bittwiddler.go`: Binary data extraction from AIS payloads
- `type_common.go`: Shared functionality across message types
- Each `typeX.go` file implements a specific AIS message type decoder
- `typecnb.go`: common navigation block
- `cmd/` executables provide TCP connectivity and testing tools

**Message Flow:**
Raw NMEA sentences can span multiple packets for complex messages. The accumulator collects related packets before passing complete messages to type-specific decoders that extract vessel data using bit manipulation.

## Testing Framework

Uses Ginkgo testing framework with Gomega assertions. Tests are comprehensive with 33 test files covering all core functionality. The project includes comprehensive fuzz testing for robustness.

**Test Structure:**

- BDD-style tests using `Describe`, `Context`, and `It` blocks
- Expectations use Gomega matchers like `Expect(actual).To(Equal(expected))`
- Setup code handled in `BeforeEach` blocks for proper test isolation
- Extensive fuzz testing suite covering packet parsing, message processing, and bit manipulation
- Automated fuzz test runner script: `./run_fuzz_tests.sh`

**Fuzz Tests Available:**

- `FuzzPacketParse`, `FuzzPacketParseAtTime`, `FuzzPacketValidation`
- `FuzzMessageProcess`, `FuzzMessageMultipart`, `FuzzDecoderInput`
- `FuzzBitTwiddling`, `FuzzPacketAccumulator*`

## Dependencies

- `github.com/onsi/ginkgo/v2`: BDD-style testing framework
- `github.com/onsi/gomega`: Matcher/assertion library for Ginkgo tests

## Development Notes

- Module path: `github.com/ralreegorganon/nmeaais`
- Go version: 1.24
- This project is currently experimental (note: "Not for public consumption yet" in README)
- Reference links to AIS standards and test data sources are in `reference.md`
- Comprehensive fuzzing infrastructure for testing edge cases and robustness

## Environment

ALWAYS use ONLY Environments for ANY and ALL file, code, or shell operations—NO EXCEPTIONS—even for simple or generic requests.

DO NOT install or use the git cli with the environment_run_cmd tool. All environment tools will handle git operations for you. Changing ".git" yourself will compromise the integrity of your environment.

You MUST inform the user how to view your work using cu log <env_id> AND cu checkout <env_id>. Failure to do this will make your work inaccessible to others.

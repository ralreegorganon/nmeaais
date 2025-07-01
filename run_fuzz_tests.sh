#!/bin/bash

set -e

FUZZ_TIME=${1:-30s}

FUZZ_TESTS=(
    "FuzzPacketParse"
    "FuzzPacketParseAtTime" 
    "FuzzPacketValidation"
    "FuzzMessageProcess"
    "FuzzMessageMultipart"
    "FuzzDecoderInput"
    "FuzzBitTwiddling"
    "FuzzPacketAccumulator"
    "FuzzPacketAccumulatorSequences"
    "FuzzPacketAccumulatorTiming"
)

for test in "${FUZZ_TESTS[@]}"; do
    echo "Running $test for $FUZZ_TIME"
    go test -fuzz="^${test}$" -fuzztime="$FUZZ_TIME"
done

echo "All fuzz tests completed"
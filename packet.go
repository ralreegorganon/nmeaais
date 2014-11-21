package nmeaais

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	Raw                 string
	StartDelimiter      string
	Tag                 string
	FragmentCount       int64
	FragmentNumber      int64
	SequentialMessageId int64
	RadioChannel        string
	Payload             string
	FillBits            int64
	Checksum            string
}

const aivdmTag string = "AIVDM"
const startDelimiter string = "!"
const validTagAndDelimiter string = startDelimiter + aivdmTag

func Parse(raw string) (*Packet, error) {
	raw = strings.TrimSpace(raw)

	if !strings.HasPrefix(raw, startDelimiter) {
		return nil, errors.New("nmeaais: invalid start delimiter")
	}

	parts := strings.Split(raw, ",")
	partsCount := len(parts)
	if partsCount != 7 {
		return nil, fmt.Errorf("nmeaais: has %v parts instead of 7", partsCount)
	}
	for i, _ := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	packet := &Packet{
		Raw:            raw,
		StartDelimiter: startDelimiter,
	}

	err := packet.parseTag(parts[0])
	if err != nil {
		return nil, err
	}

	err = packet.parseFragmentCount(parts[1])
	if err != nil {
		return nil, err
	}

	err = packet.parseFragmentNumber(parts[2])
	if err != nil {
		return nil, err
	}

	err = packet.parseSequentialMessageId(parts[3])
	if err != nil {
		return nil, err
	}

	err = packet.parseRadioChannel(parts[4])
	if err != nil {
		return nil, err
	}

	packet.Payload = parts[5]

	fillBitsAndChecksumParts := strings.Split(parts[6], "*")
	if len(fillBitsAndChecksumParts) != 2 {
		return nil, fmt.Errorf("nmeaais: invalid fillbits and checksum of '%v'", parts[6])
	}

	err = packet.parseFillBits(fillBitsAndChecksumParts[0])
	if err != nil {
		return nil, err
	}

	packet.Checksum = fillBitsAndChecksumParts[1]

	err = packet.validate()
	if err != nil {
		return nil, err
	}

	return packet, nil
}

func (p *Packet) parseTag(part string) error {
	tag := strings.TrimLeft(part, startDelimiter)
	p.Tag = tag
	return nil
}

func (p *Packet) parseFragmentCount(part string) error {
	fragmentCount, err := strconv.ParseInt(part, 10, 64)
	if err != nil {
		return err
	}
	p.FragmentCount = fragmentCount
	return nil
}

func (p *Packet) parseFragmentNumber(part string) error {
	fragmentNumber, err := strconv.ParseInt(part, 10, 64)
	if err != nil {
		return err
	}
	p.FragmentNumber = fragmentNumber
	return nil
}

func (p *Packet) parseSequentialMessageId(part string) error {
	if part != "" {
		sequentialMessageId, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return err
		}
		p.SequentialMessageId = sequentialMessageId
	}
	return nil
}

func (p *Packet) parseRadioChannel(part string) error {
	switch part {
	case "A":
		fallthrough
	case "B":
		fallthrough
	case "1":
		fallthrough
	case "2":
		p.RadioChannel = part
		return nil
	case "":
		return nil
	default:
		return fmt.Errorf("nmeaais: invalid radio channel '%v'", part)
	}
}

func (p *Packet) parseFillBits(part string) error {
	fillBits, err := strconv.ParseInt(part, 10, 64)
	if err != nil {
		return err
	}
	if fillBits < 0 || fillBits > 5 {
		return fmt.Errorf("nmeaais: fill bits of '%v' outside valid range of 0-5", fillBits)
	}
	p.FillBits = fillBits
	return nil
}

func (p *Packet) validate() error {
	rawForChecksum := strings.TrimSuffix(strings.TrimPrefix(p.Raw, p.StartDelimiter), "*"+p.Checksum)

	var checksum uint8 = 0
	for i := 0; i < len(rawForChecksum); i++ {
		checksum = checksum ^ rawForChecksum[i]
	}

	checksumBytes := []byte{
		checksum,
	}

	crc := strings.ToUpper(hex.EncodeToString(checksumBytes))

	if crc != p.Checksum {
		return fmt.Errorf("nmeaais: checksum '%v' doesn't match expected '%v'", crc, p.Checksum)
	}

	return nil
}

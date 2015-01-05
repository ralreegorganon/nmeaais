package nmeaais

import "fmt"

type StaticDataReportA struct {
	MessageType     int64
	RepeatIndicator int64
	MMSI            int64
	PartNumber      int64
	VesselName      string
}

func (m *Message) IsStaticDataReportA() (bool, error) {
	var validMessageType int64 = 24
	if m.MessageType != validMessageType {
		return false, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	partNumber := asUInt(m.unarmoredPayload, 38, 2)

	switch partNumber {
	case 0:
		return true, nil
	case 1:
		return false, nil
	default:
		return false, fmt.Errorf("nmeaais: type 24 part type of %v is invalid", partNumber)
	}
}

func (m *Message) GetAsStaticDataReportA() (*StaticDataReportA, error) {
	var validMessageType int64 = 24

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p := &StaticDataReportA{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		PartNumber:      int64(asUInt(m.unarmoredPayload, 38, 2)),
		VesselName:      asString(m.unarmoredPayload, 40, 120),
	}
	return p, nil
}

type StaticDataReportB struct {
	MessageType          int64
	RepeatIndicator      int64
	MMSI                 int64
	PartNumber           int64
	ShipType             string
	VendorID             string
	UnitModelCode        int64
	SerialNumber         int64
	CallSign             string
	DimensionToBow       int64
	DimensionToStern     int64
	DimensionToPort      int64
	DimensionToStarboard int64
	MothershipMMSI       int64
}

func (m *Message) IsStaticDataReportB() (bool, error) {
	var validMessageType int64 = 24
	if m.MessageType != validMessageType {
		return false, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	partNumber := asUInt(m.unarmoredPayload, 38, 2)

	switch partNumber {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("nmeaais: type 24 part type of %v is invalid", partNumber)
	}
}

func (m *Message) GetAsStaticDataReportB() (*StaticDataReportB, error) {
	var validMessageType int64 = 24

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p := &StaticDataReportB{
		MessageType:     m.MessageType,
		RepeatIndicator: m.RepeatIndicator,
		MMSI:            m.MMSI,
		PartNumber:      int64(asUInt(m.unarmoredPayload, 38, 2)),
		ShipType:        shipType(asUInt(m.unarmoredPayload, 40, 8)),
		VendorID:        asString(m.unarmoredPayload, 48, 18),
		UnitModelCode:   int64(asUInt(m.unarmoredPayload, 66, 4)),
		SerialNumber:    int64(asUInt(m.unarmoredPayload, 70, 20)),
		CallSign:        asString(m.unarmoredPayload, 90, 42),
	}

	isAux := m.MMSI >= 980000000
	if isAux {
		p.MothershipMMSI = int64(asUInt(m.unarmoredPayload, 132, 30))
	} else {
		p.DimensionToBow = int64(asUInt(m.unarmoredPayload, 132, 9))
		p.DimensionToStern = int64(asUInt(m.unarmoredPayload, 141, 9))
		p.DimensionToPort = int64(asUInt(m.unarmoredPayload, 150, 6))
		p.DimensionToStarboard = int64(asUInt(m.unarmoredPayload, 156, 6))
	}
	return p, nil
}

func shipType(st uint) string {
	if st < 0 || st > shipTypesMax {
		return "Undefined"
	}
	return shipTypes[st]
}

var shipTypesMax = uint(len(shipTypes) - 1)
var shipTypes = []string{
	"Not available",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"Wing in ground (WIG), all ships of this type",
	"Wing in ground (WIG), Hazardous category A",
	"Wing in ground (WIG), Hazardous category B",
	"Wing in ground (WIG), Hazardous category C",
	"Wing in ground (WIG), Hazardous category D",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Wing in ground (WIG), Reserved for future use",
	"Fishing",
	"Towing",
	"Towing: length exceeds 200m or breadth exceeds 25m",
	"Dredging or underwater ops",
	"Diving ops",
	"Military ops",
	"Sailing",
	"Pleasure Craft",
	"Reserved",
	"Reserved",
	"High speed craft (HSC), all ships of this type",
	"High speed craft (HSC), Hazardous category A",
	"High speed craft (HSC), Hazardous category B",
	"High speed craft (HSC), Hazardous category C",
	"High speed craft (HSC), Hazardous category D",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), Reserved for future use",
	"High speed craft (HSC), No additional information",
	"Pilot Vessel",
	"Search and Rescue vessel",
	"Tug",
	"Port Tender",
	"Anti-pollution equipment",
	"Law Enforcement",
	"Spare - Local Vessel",
	"Spare - Local Vessel",
	"Medical Transport",
	"Noncombatant ship according to RR Resolution No. 18",
	"Passenger, all ships of this type",
	"Passenger, Hazardous category A",
	"Passenger, Hazardous category B",
	"Passenger, Hazardous category C",
	"Passenger, Hazardous category D",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, Reserved for future use",
	"Passenger, No additional information",
	"Cargo, all ships of this type",
	"Cargo, Hazardous category A",
	"Cargo, Hazardous category B",
	"Cargo, Hazardous category C",
	"Cargo, Hazardous category D",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, Reserved for future use",
	"Cargo, No additional information",
	"Tanker, all ships of this type",
	"Tanker, Hazardous category A",
	"Tanker, Hazardous category B",
	"Tanker, Hazardous category C",
	"Tanker, Hazardous category D",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, Reserved for future use",
	"Tanker, No additional information",
	"Other Type, all ships of this type",
	"Other Type, Hazardous category A",
	"Other Type, Hazardous category B",
	"Other Type, Hazardous category C",
	"Other Type, Hazardous category D",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, Reserved for future use",
	"Other Type, no additional information",
}

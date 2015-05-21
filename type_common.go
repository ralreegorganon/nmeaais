package nmeaais

import (
	"math"
)

func asBool(b uint) bool {
	if b == 1 {
		return true
	}

	return false
}

func latlon(l int) float64 {
	return float64(l) / 600000
}

func rateOfTurn(rot int) float64 {
	if rot == 128 {
		return math.NaN()
	}
	if rot == 127 || rot == -127 {
		return math.Inf(rot)
	}
	floatified := float64(rot)
	value := floatified / 4.733
	value *= value
	return math.Copysign(value, floatified)
}

func speedOverGround(sog uint) float64 {
	return float64(sog) / 10
}

func courseOverGround(cog uint) float64 {
	return float64(cog) / 10
}

func maneuverIndicator(mi uint) string {
	switch mi {
	case 0:
		return "Not available"
	case 1:
		return "No special maneuver"
	case 2:
		return "Special maneuver"
	default:
		return "Not available"
	}
}

var aidTypesMax = uint(len(aidTypes) - 1)
var aidTypes = []string{
	"Default, Type of Aid to Navigation not specified",
	"Reference point",
	"RACON (radar transponder marking a navigation hazard)",
	"Fixed structure off shore, such as oil platforms, wind farms,rigs.",
	"Spare, Reserved for future use.",
	"Light, without sectors",
	"Light, with sectors",
	"Leading Light Front",
	"Leading Light Rear",
	"Beacon, Cardinal N",
	"Beacon, Cardinal E",
	"Beacon, Cardinal S",
	"Beacon, Cardinal W",
	"Beacon, Port hand",
	"Beacon, Starboard hand",
	"Beacon, Preferred Channel port hand",
	"Beacon, Preferred Channel starboard hand",
	"Beacon, Isolated danger",
	"Beacon, Safe water",
	"Beacon, Special mark",
	"Cardinal Mark N",
	"Cardinal Mark E",
	"Cardinal Mark S",
	"Cardinal Mark W",
	"Port hand Mark",
	"Starboard hand Mark",
	"Preferred Channel Port hand",
	"Preferred Channel Starboard hand",
	"Isolated danger",
	"Safe Water",
	"Special Mark",
	"Light Vessel / LANBY / Rigs",
}

func aidType(at uint) string {
	if at < 0 || at > aidTypesMax {
		return "Undefined"
	}
	return aidTypes[at]
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

func shipType(st uint) string {
	if st < 0 || st > shipTypesMax {
		return "Undefined"
	}
	return shipTypes[st]
}

var epfdTypesMax = uint(len(epfdTypes) - 1)
var epfdTypes = []string{
	"Undefined",
	"GPS",
	"GLONASS",
	"Combined GPS/GLONASS",
	"Loran-C",
	"Chayka",
	"Integrated navigation system",
	"Surveyed",
	"Galileo",
	"Undefined",
	"Undefined",
	"Undefined",
	"Undefined",
	"Undefined",
	"Undefined",
	"Undefined",
}

func epfdType(et uint) string {
	if et < 0 || et > epfdTypesMax {
		return "Undefined"
	}
	return epfdTypes[et]
}

var navigationStatusesMax = uint(len(navigationStatuses) - 1)
var navigationStatuses = []string{
	"Under way using engine",
	"At anchor",
	"Not under command",
	"Restricted manoeuverability",
	"Constrained by her draught",
	"Moored",
	"Aground",
	"Engaged in Fishing",
	"Under way sailing",
	"Reserved for future amendment of Navigational Status for HSC",
	"Reserved for future amendment of Navigational Status for WIG",
	"Reserved for future use",
	"Reserved for future use",
	"Reserved for future use",
	"AIS-SART is active",
	"Not defined",
}

func navigationStatus(ns uint) string {
	if ns < 0 || ns > navigationStatusesMax {
		return "Not defined"
	}
	return navigationStatuses[ns]
}

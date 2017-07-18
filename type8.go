package nmeaais

import "fmt"

type BinaryBroadcastMessage struct {
	MessageType        int64
	RepeatIndicator    int64
	MMSI               int64
	DesignatedAreaCode int64
	FunctionalID       int64
	Data               []byte
}

func (m *Message) GetAsBinaryBroadcastMessage() (p *BinaryBroadcastMessage, err error) {
	defer func() {
		if r := recover(); r != nil {
			p = nil
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	var validMessageType int64 = 8

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	p = &BinaryBroadcastMessage{
		MessageType:        m.MessageType,
		RepeatIndicator:    m.RepeatIndicator,
		MMSI:               m.MMSI,
		DesignatedAreaCode: int64(asUInt(m.unarmoredPayload, 40, 10)),
		FunctionalID:       int64(asUInt(m.unarmoredPayload, 50, 6)),
		Data:               asBinary(m.unarmoredPayload, 56, uint(m.bitLength-56)),
	}

	return
}

type MeteorologicalAndHydrologicalData struct {
	MessageType              int64
	RepeatIndicator          int64
	MMSI                     int64
	DesignatedAreaCode       int64
	FunctionalID             int64
	Longitude                float64
	Latitude                 float64
	PositionAccuracy         bool
	Day                      int64
	Hour                     int64
	Minute                   int64
	AverageWindSpeed         int64
	GustSpeed                int64
	WindDirection            int64
	WindGustDirection        int64
	AirTemperature           float64
	RelativeHumidity         int64
	DewPoint                 float64
	AirPressure              int64
	PressureTendency         string
	MaxVisibility            bool
	HorizontalVisibility     float64
	WaterLevel               float64
	WaterLevelTrend          string
	SurfaceCurrentSpeed      int64
	SurfaceCurrentDirection  int64
	CurrentSpeed1            int64
	CurrentDirection1        int64
	CurrentMeasurementDepth1 int64
	CurrentSpeed2            int64
	CurrentDirection2        int64
	CurrentMeasurementDepth2 int64
	WaveHeight               int64
	WavePeriod               int64
	WaveDirection            int64
	SwellHeight              int64
	SwellPeriod              int64
	SwellDirection           int64
	SeaState                 string
	WaterTemperature         float64
	Precipitation            string
	Salinity                 int64
	Ice                      int64
}

func (m *Message) GetAsMeteorologicalAndHydrologicalData() (p *MeteorologicalAndHydrologicalData, err error) {
	var validMessageType int64 = 8

	if m.MessageType != validMessageType {
		return nil, fmt.Errorf("nmeaais: tried to get message as type %v, but is type %v", validMessageType, m.MessageType)
	}

	var validDesignatedAreaCode int64 = 1
	var validFunctionalID int64 = 31

	dac := int64(asUInt(m.unarmoredPayload, 40, 10))
	fid := int64(asUInt(m.unarmoredPayload, 50, 6))

	if dac != validDesignatedAreaCode || fid != validFunctionalID {
		return nil, fmt.Errorf("nmeaais: tried to get message type 8 binary payload for DAC %v FID %v, but is DAC %v FID %v", validDesignatedAreaCode, validFunctionalID, dac, fid)
	}

	p = &MeteorologicalAndHydrologicalData{
		MessageType:          m.MessageType,
		RepeatIndicator:      m.RepeatIndicator,
		MMSI:                 m.MMSI,
		DesignatedAreaCode:   dac,
		FunctionalID:         fid,
		Longitude:            latlonMeteorologicalAndHydrological(asInt(m.unarmoredPayload, 56, 25)),
		Latitude:             latlonMeteorologicalAndHydrological(asInt(m.unarmoredPayload, 81, 24)),
		PositionAccuracy:     asBool(asUInt(m.unarmoredPayload, 105, 1)),
		Day:                  int64(asUInt(m.unarmoredPayload, 106, 5)),
		Hour:                 int64(asUInt(m.unarmoredPayload, 111, 5)),
		Minute:               int64(asUInt(m.unarmoredPayload, 116, 6)),
		AverageWindSpeed:     int64(asUInt(m.unarmoredPayload, 122, 7)),
		GustSpeed:            int64(asUInt(m.unarmoredPayload, 129, 7)),
		WindDirection:        int64(asUInt(m.unarmoredPayload, 136, 9)),
		WindGustDirection:    int64(asUInt(m.unarmoredPayload, 145, 9)),
		AirTemperature:       airTemperature(asInt(m.unarmoredPayload, 154, 11)),
		RelativeHumidity:     int64(asUInt(m.unarmoredPayload, 165, 7)),
		DewPoint:             airTemperature(asInt(m.unarmoredPayload, 172, 10)),
		AirPressure:          int64(asUInt(m.unarmoredPayload, 182, 9)),
		PressureTendency:     pressureTendency(asUInt(m.unarmoredPayload, 191, 2)),
		MaxVisibility:        asBool(asUInt(m.unarmoredPayload, 193, 1)),
		HorizontalVisibility: horizontalVisibility(asUInt(m.unarmoredPayload, 194, 6)),
	}

	return
}

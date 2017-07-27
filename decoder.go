package nmeaais

import (
	"fmt"
	"time"
)

type DecoderInput struct {
	Input     string
	Timestamp time.Time
}

type DecoderOutput struct {
	SourcePackets  []*Packet
	SourceMessage  *Message
	DecodedMessage interface{}
	Error          error
	Timestamp      time.Time
}

type Decoder struct {
	Input             chan DecoderInput
	Output            chan DecoderOutput
	packetAccumulator *PacketAccumulator
}

func NewDecoder() *Decoder {
	d := &Decoder{
		Input:             make(chan DecoderInput),
		Output:            make(chan DecoderOutput),
		packetAccumulator: NewPacketAccumulator(),
	}

	go d.parse()
	go d.decode()

	return d
}

func (d *Decoder) parse() {
	for s := range d.Input {
		packet, err := ParseAtTime(s.Input, s.Timestamp)
		if err != nil {
			continue
		}
		d.packetAccumulator.Packets <- packet
	}
	close(d.packetAccumulator.Packets)
}

func (d *Decoder) decode() {
	for r := range d.packetAccumulator.Results {

		if r.Error != nil {
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: nil,
				Error:          r.Error,
				Timestamp:      r.Timestamp,
			}
			continue
		}

		switch r.Message.MessageType {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			x, err := r.Message.GetAsPositionReportClassA()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 4:
			x, err := r.Message.GetAsBaseStationReport()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 5:
			x, err := r.Message.GetAsStaticAndVoyageRelatedData()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 7:
			x, err := r.Message.GetAsBinaryAcknowledge()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 8:
			x, err := r.Message.GetAsBinaryBroadcastMessage()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 9:
			x, err := r.Message.GetAsStandardSARAircraftPositionReport()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 10:
			x, err := r.Message.GetAsUTCDateInquiry()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 11:
			x, err := r.Message.GetAsUTCDateResponse()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 15:
			x, err := r.Message.GetAsInterrogation()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 16:
			x, err := r.Message.GetAsAssignmentModeCommand()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 17:
			x, err := r.Message.GetAsDGNSSBroadcastBinaryMessage()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 18:
			x, err := r.Message.GetAsPositionReportClassBStandard()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 20:
			x, err := r.Message.GetAsDataLinkManagementMessage()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break

		case 21:
			x, err := r.Message.GetAsAidToNavigationReport()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 22:
			x, err := r.Message.GetAsChannelManagement()
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
			break
		case 24:
			if ok, _ := r.Message.IsStaticDataReportA(); ok {
				x, err := r.Message.GetAsStaticDataReportA()
				d.Output <- DecoderOutput{
					SourcePackets:  r.Packets,
					SourceMessage:  r.Message,
					DecodedMessage: x,
					Error:          err,
					Timestamp:      r.Timestamp,
				}
			}
			if ok, _ := r.Message.IsStaticDataReportB(); ok {
				x, err := r.Message.GetAsStaticDataReportB()
				d.Output <- DecoderOutput{
					SourcePackets:  r.Packets,
					SourceMessage:  r.Message,
					DecodedMessage: x,
					Error:          err,
					Timestamp:      r.Timestamp,
				}
			}
			break
		default:
			err := fmt.Errorf("nmeaais: unsupported message of type %v from %v", r.Message.MessageType, r.Message.MMSI)
			d.Output <- DecoderOutput{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: nil,
				Error:          err,
				Timestamp:      r.Timestamp,
			}
		}
	}
	close(d.Output)
}

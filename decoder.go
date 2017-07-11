package nmeaais

import "fmt"

type DecoderResult struct {
	SourcePackets  []*Packet
	SourceMessage  *Message
	DecodedMessage interface{}
	Error          error
}

type Decoder struct {
	Input             chan string
	Output            chan *DecoderResult
	packetAccumulator *PacketAccumulator
}

func NewDecoder() *Decoder {
	d := &Decoder{
		Input:             make(chan string),
		Output:            make(chan *DecoderResult),
		packetAccumulator: NewPacketAccumulator(),
	}

	go d.parse()
	go d.decode()

	return d
}

func (d *Decoder) parse() {
	for s := range d.Input {
		packet, err := Parse(s)
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
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: nil,
				Error:          r.Error,
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
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 4:
			x, err := r.Message.GetAsBaseStationReport()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 5:
			x, err := r.Message.GetAsStaticAndVoyageRelatedData()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 7:
			x, err := r.Message.GetAsBinaryAcknowledge()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 8:
			x, err := r.Message.GetAsBinaryBroadcastMessage()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 9:
			x, err := r.Message.GetAsStandardSARAircraftPositionReport()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 15:
			x, err := r.Message.GetAsInterrogation()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 16:
			x, err := r.Message.GetAsAssignmentModeCommand()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 17:
			x, err := r.Message.GetAsDGNSSBroadcastBinaryMessage()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 18:
			x, err := r.Message.GetAsPositionReportClassBStandard()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 20:
			x, err := r.Message.GetAsDataLinkManagementMessage()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break

		case 21:
			x, err := r.Message.GetAsAidToNavigationReport()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 22:
			x, err := r.Message.GetAsChannelManagement()
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: x,
				Error:          err,
			}
			break
		case 24:
			if ok, _ := r.Message.IsStaticDataReportA(); ok {
				x, err := r.Message.GetAsStaticDataReportA()
				d.Output <- &DecoderResult{
					SourcePackets:  r.Packets,
					SourceMessage:  r.Message,
					DecodedMessage: x,
					Error:          err,
				}
			}
			if ok, _ := r.Message.IsStaticDataReportB(); ok {
				x, err := r.Message.GetAsStaticDataReportB()
				d.Output <- &DecoderResult{
					SourcePackets:  r.Packets,
					SourceMessage:  r.Message,
					DecodedMessage: x,
					Error:          err,
				}
			}
			break
		default:
			err := fmt.Errorf("nmeaais: unsupported message of type %v from %v", r.Message.MessageType, r.Message.MMSI)
			d.Output <- &DecoderResult{
				SourcePackets:  r.Packets,
				SourceMessage:  r.Message,
				DecodedMessage: nil,
				Error:          err,
			}
		}
	}
	close(d.Output)
}

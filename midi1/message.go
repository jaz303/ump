package midi1

import (
	"github.com/jaz303/ump"
)

const (
	Clock         = ump.Word(ump.MsgTypeSystem | (0xF8 << 16))
	Start         = ump.Word(ump.MsgTypeSystem | (0xFA << 16))
	Continue      = ump.Word(ump.MsgTypeSystem | (0xFB << 16))
	Stop          = ump.Word(ump.MsgTypeSystem | (0xFC << 16))
	ActiveSensing = ump.Word(ump.MsgTypeSystem | (0xFE << 16))
	Reset         = ump.Word(ump.MsgTypeSystem | (0xFF << 16))
)

const (
	noteOff         = ump.MsgTypeMIDIv1 | (0b1000 << 20)
	noteOn          = ump.MsgTypeMIDIv1 | (0b1001 << 20)
	polyPressure    = ump.MsgTypeMIDIv1 | (0b1010 << 20)
	controlChange   = ump.MsgTypeMIDIv1 | (0b1011 << 20)
	programChange   = ump.MsgTypeMIDIv1 | (0b1100 << 20)
	channelPressure = ump.MsgTypeMIDIv1 | (0b1101 << 20)
	pitchBend       = ump.MsgTypeMIDIv1 | (0b1110 << 20)

	channelShift = 16
)

func NoteOff(channel uint8, note int8, velocity int8) ump.Word {
	return noteOff |
		(ump.Word(channel&0x0F) << channelShift) |
		(ump.Word(note) << 8) |
		ump.Word(velocity)
}

func NoteOn(channel uint8, note int8, velocity int8) ump.Word {
	return noteOn |
		(ump.Word(channel&0x0F) << channelShift) |
		(ump.Word(note) << 8) |
		ump.Word(velocity)
}

// TODO: poly pressure

func ControlChange(channel uint8, controller, value int8) ump.Word {
	return controlChange |
		(ump.Word(channel&0x0F) << channelShift) |
		(ump.Word(controller) << 8) |
		ump.Word(value)
}

// TODO: program change
// TODO: channel pressure
// TODO: pitch bend

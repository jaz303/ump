package midi1

import "github.com/jaz303/ump"

func IsNoteOn(w ump.Word) bool        { return (w & 0xFFF00000) == noteOn }
func IsNoteOff(w ump.Word) bool       { return (w & 0xFFF00000) == noteOff }
func IsControlChange(w ump.Word) bool { return (w & 0xFFF00000) == controlChange }

func Channel(w ump.Word) int { return int(w>>channelShift) & 0x0F }

func Note(w ump.Word) int8     { return int8((w >> 8) & 0x7F) }
func Velocity(w ump.Word) int8 { return int8(w & 0x7F) }

func Controller(w ump.Word) int8 { return int8((w >> 8) & 0x7F) }
func Data(w ump.Word) int8       { return int8(w & 0x7F) }

package unicode

import (
	"github.com/reiver/go-ascii"
)

// Unicode short-codes.
const (
	NUL rune = ascii.NUL // 0x00
	SOH rune = ascii.SOH // 0x01
	STX rune = ascii.STX // 0x02
	ETX rune = ascii.ETX // 0x03
	EOT rune = ascii.ETX // 0x04
	ENQ rune = ascii.ENQ // 0x05
	ACK rune = ascii.ACK // 0x06
	BEL rune = ascii.BEL // 0x07
	BS  rune = ascii.BS  // 0x08
	HT  rune = ascii.HT  // 0x09
	LF  rune = ascii.LF  // 0x0A
	VT  rune = ascii.VT  // 0x0B
	FF  rune = ascii.FF  // 0x0C
	CR  rune = ascii.CR  // 0x0D
	SO  rune = ascii.SO  // 0x0E
	SI  rune = ascii.SI  // 0x0F
	DLE rune = ascii.DLE // 0x10
	DC1 rune = ascii.DC1 // 0x11
	DC2 rune = ascii.DC2 // 0x12
	DC3 rune = ascii.DC3 // 0x13
	DC4 rune = ascii.DC4 // 0x14
	NAK rune = ascii.NAK // 0x15
	SYN rune = ascii.SYN // 0x16
	ETB rune = ascii.ETB // 0x17
	CAN rune = ascii.CAN // 0x18
	EM  rune = ascii.EM  // 0x19
	SUB rune = ascii.SUB // 0x1A
	ESC rune = ascii.ESC // 0x1B
	FS  rune = ascii.FS  // 0x1C
	GS  rune = ascii.GS  // 0x1D
	RS  rune = ascii.RS  // 0x1E
	US  rune = ascii.US  // 0x1F
	SP  rune = ascii.SP  // 0x20
	// ...
	DEL rune = ascii.DEL // 0x7F
	// ---
	PAD rune = 0x80
	HOP rune = 0x81
	BPH rune = 0x82
	NBH rune = 0x83
	IND rune = 0x84
	NEL rune = 0x85
	SSA rune = 0x86
	ESA rune = 0x87
	HTS rune = 0x88
	HTJ rune = 0x89
	VTS rune = 0x8A
	PLD rune = 0x8B
	PLU rune = 0x8C
	RI  rune = 0x8D
	SS2 rune = 0x8E
	SS3 rune = 0x8F
	DCS rune = 0x90
	PU1 rune = 0x91
	PU2 rune = 0x92
	STS rune = 0x93
	CCH rune = 0x94
	MW  rune = 0x95
	SPA rune = 0x96
	EPA rune = 0x97
	SOS rune = 0x98
	SGC rune = 0x99
	SCI rune = 0x9A
	CSI rune = 0x9B
	ST  rune = 0x9C
	OSC rune = 0x9D
	PM  rune = 0x9E
	APC rune = 0x9F
)

const (
	XON  rune = ascii.XON  // DC1
	XOFF rune = ascii.XOFF // DC3
)

package unicode

import (
	"github.com/reiver/go-ascii"
)

// Unicode short-codes.
const (
	NUL = ascii.NUL // 0x00
	SOH = ascii.SOH // 0x01
	STX = ascii.STX // 0x02
	ETX = ascii.ETX // 0x03
	EOT = ascii.ETX // 0x04
	ENQ = ascii.ENQ // 0x05
	ACK = ascii.ACK // 0x06
	BEL = ascii.BEL // 0x07
	BS  = ascii.BS  // 0x08
	HT  = ascii.HT  // 0x09
	LF  = ascii.LF  // 0x0A
	VT  = ascii.VT  // 0x0B
	FF  = ascii.FF  // 0x0C
	CR  = ascii.CR  // 0x0D
	SO  = ascii.SO  // 0x0E
	SI  = ascii.SI  // 0x0F
	DLE = ascii.DLE // 0x10
	DC1 = ascii.DC1 // 0x11
	DC2 = ascii.DC2 // 0x12
	DC3 = ascii.DC3 // 0x13
	DC4 = ascii.DC4 // 0x14
	NAK = ascii.NAK // 0x15
	SYN = ascii.SYN // 0x16
	ETB = ascii.ETB // 0x17
	CAN = ascii.CAN // 0x18
	EM  = ascii.EM  // 0x19
	SUB = ascii.SUB // 0x1A
	ESC = ascii.ESC // 0x1B
	FS  = ascii.FS  // 0x1C
	GS  = ascii.GS  // 0x1D
	RS  = ascii.RS  // 0x1E
	US  = ascii.US  // 0x1F
	SP  = ascii.SP  // 0x20
	// ...
	DEL = ascii.DEL // 0x7F
	// ---
	PAD = 0x80
	HOP = 0x81
	BPH = 0x82
	NBH = 0x83
	IND = 0x84
	NEL = 0x85
	SSA = 0x86
	ESA = 0x87
	HTS = 0x88
	HTJ = 0x89
	VTS = 0x8A
	PLD = 0x8B
	PLU = 0x8C
	RI  = 0x8D
	SS2 = 0x8E
	SS3 = 0x8F
	DCS = 0x90
	PU1 = 0x91
	PU2 = 0x92
	STS = 0x93
	CCH = 0x94
	MW  = 0x95
	SPA = 0x96
	EPA = 0x97
	SOS = 0x98
	SGC = 0x99
	SCI = 0x9A
	CSI = 0x9B
	ST  = 0x9C
	OSC = 0x9D
	PM  = 0x9E
	APC = 0x9F
)

const (
	XON  = ascii.XON  // DC1
	XOFF = ascii.XOFF // DC3
)

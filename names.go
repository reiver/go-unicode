package unicode

// ASCII names.
const (
	Null                                 rune =  NUL
	StartOfHeading                       rune =  SOH
	StartOfText                          rune =  STX
	EndOfText                            rune =  ETX
	EndOfTransmission                    rune =  EOT
	Enquiry                              rune =  ENQ
	Acknowledgment                       rune =  ACK
	Bell                                 rune =  BEL
	Backspace                            rune =  BS
	HorizontalTab                        rune =  HT
	LineFeed                             rune =  LF
	VerticalTab                          rune =  VT
	FormFeed                             rune =  FF
	CarriageReturn                       rune =  CR
	ShiftOut                             rune =  SO
	ShiftIn                              rune =  SI
	DataLinkEscape                       rune =  DLE
	DeviceControl1                       rune =  DC1
	DeviceControl2                       rune =  DC2
	DeviceControl3                       rune =  DC3
	DeviceControl4                       rune =  DC4
	NegativeAcknowledgement              rune =  NAK
	SynchronousIdle                      rune =  SYN
	EndOfTransmitBlock                   rune =  ETB
	Cancel                               rune =  CAN
	EndOfMedium                          rune =  EM
	Substitute                           rune =  SUB
	Escape                               rune =  ESC
	FileSeparator                        rune =  FS
	GroupSeparator                       rune =  GS
	RecordSeparator                      rune =  RS
	UnitSeparator                        rune =  US
	Space                                rune =  SP
	// ...
	Delete                               rune = DEL
	// ---
	PaddingCharacter                     rune = PAD
	HighOctetPreset                      rune = HOP
	BreakPermittedHere                   rune = BPH
	NoBreakHere                          rune = NBH
	Index                                rune = IND
	NextLine                             rune = NEL
	StartOfSelectedArea                  rune = SSA
	EndOfSelectedArea                    rune = ESA
	CharacterTabulationSet               rune = HTS ; HorizontalTabulationSet               rune = HTS
	CharacterTabulationWithJustification rune = HTJ ; HorizontalTabulationWithJustification rune = HTJ
	LineTabulationSet                    rune = VTS ; VerticalTabulationSet                 rune = VTS
	PartialLineDown                      rune = PLD ; PartialLineForward                    rune = PLD
	PartialLineBackward                  rune = PLU ; PartialLineUp                         rune = PLU
	ReverseIndex                         rune = RI  ; ReverseLineFeed                       rune = RI
	SingleShiftTwo                       rune = SS2
	SingleShiftThree                     rune = SS3
	DeviceControlString                  rune = DCS
	PrivateUseOne                        rune = PU1 ; PrivateUse1                           rune = PU1
	PrivateUseTwo                        rune = PU2 ; PrivateUse2                           rune = PU2
	SetTransmitState                     rune = STS
	CancelCharacter                      rune = CCH
	MessageWaiting                       rune = MW
	StartOfGuardedArea                   rune = SPA ; StartOfProtectedArea                  rune = SPA
	EndOfGuardedArea                     rune = EPA ; EndOfProtectedArea                    rune = EPA
	StartOfString                        rune = SOS
	SingleGraphicCharacterIntroducer     rune = SGC
	SingleCharacterIntroducer            rune = SCI
	ControlSequenceIntroducer            rune = CSI
	StringTerminator                     rune = ST
	OperatingSystemCommand               rune = OSC
	PrivacyMessage                       rune = PM
	ApplicationProgramCommand            rune = APC
)

const (
	TransmitOn  rune = XON
	TransmitOff rune = XOFF
)

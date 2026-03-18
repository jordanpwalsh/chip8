package main

const (
	MemSize			= 4096
	ROMStart		= 0x200
	StackDepth	= 16
	NumRegs			= 16
	DisplayW		= 64
	DisplayH		= 32
)

type CPU struct {
	mem [MemSize]byte

	V	  [NumRegs]byte
	I 	uint16
	PC	uint16

	stack [StackDepth]uint16
	SP	uint8

	delay byte
	sound byte

	display [DisplayH][DisplayW]bool
	keys [16]bool
}

func NewCPU() *CPU {
	cpu := &CPU{
		PC: ROMStart,
	}

	cpu.loadFonts()
	return cpu
}



package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: chip8 <path-to-rom>")
		os.Exit(1)
	}

	cpu := NewCPU()

	if err := cpu.LoadROM(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	//Peek at first opcode so we can verify it loaded
	b0 := cpu.mem[cpu.PC]
	b1 := cpu.mem[cpu.PC+1]
	fmt.Printf("loaded OK - first opcode: 0x%02X%02X\n", b0, b1)

	//Verify fonts loaded correctly: digit '0' starts at 0x000
	fmt.Printf("font[0] first byte: 0x%02X (expect 0xF0)\n", cpu.mem[0x000])
	fmt.Printf("font[F] first byte: 0x%02X (expect 0xF0)\n", cpu.mem[0x050])


}

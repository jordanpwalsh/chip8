package main

import (
	"fmt"
	"os"
)

func (c *CPU) LoadROM(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("LoadROM: %w", err)
	}

	available := MemSize - ROMStart //3584 bytes
	if len(data) > available {
		return fmt.Errorf("ROM is %d bytes, max is %d", len(data), available)
	}

	copy(c.mem[ROMStart:], data)
	return nil
}

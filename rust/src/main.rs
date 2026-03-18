struct Cpu {
    memory: [u8; 4096],
    pc: u16,
    registers: [u8; 16]
}

impl Cpu {
    fn new() -> Self {
        Cpu {
            memory: [0; 4096],
            pc: 0x200,
            registers: [0; 16],
        }
    }

    fn load_rom(&mut self, data: &[u8]) {
        self.memory[0x200..0x200 + data.len()]
            .copy_from_slice(data)
    }
}

fn main() {
    let mut cpu = Cpu::new();
    let rom: Vec<u8> = vec![0x00, 0xE0, 0x12, 0x00];
    cpu.load_rom(&rom);
    println!("First memory byte at 0x200: {:#04x}",
        cpu.memory[0x200]);

    println!("PC: {:#06x}", cpu.pc);
}



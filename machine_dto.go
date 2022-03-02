package main

import "io"

func NewMachine(instructions []Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		instruction: instructions,
		input:       in,
		output:      out,
		buf:         make([]byte, 1),
		memory:      make([]int, len(instructions)),
	}
}

func (m *Machine) Execute() {
	for m.size < len(m.instruction) {
		ins := m.instruction[m.size]

		switch ins.Type {
		case PLUS:
			m.memory[m.pointer] += ins.Count
			if m.memory[m.pointer] == 256 { // max
				m.memory[m.pointer] = 0
			}

		case SUB:
			m.memory[m.pointer] -= ins.Count
			if m.memory[m.pointer] == -1 { // min
				m.memory[m.pointer] = 255
			}

		case PLUSDATAPOINTER:
			m.pointer += ins.Count

		case SUBDATAPOINTER:
			m.pointer -= ins.Count

		case WRITECHAR:
			for i := 0; i < ins.Count; i++ {
				m.Write()
			}

		case READCHAR:
			for i := 0; i < ins.Count; i++ {
				m.Read()
			}

		case STARTLOOP:
			if m.memory[m.pointer] == 0 {
				m.size = ins.Count
				continue
			}

		case ENDLOOP:
			if m.memory[m.pointer] != 0 {
				m.size = ins.Count
				continue
			}
		}

		m.size++
	}
}

func (m *Machine) Read() {
	n, err := m.input.Read(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.pointer] = int(m.buf[0])
}

func (m *Machine) Write() {
	m.buf[0] = byte(m.memory[m.pointer])

	n, err := m.output.Write(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}

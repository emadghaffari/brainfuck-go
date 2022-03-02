package main

import "io"

func NewCompiler(reader io.Reader) *Compiler {
	buf := make([]byte, 1)
	size := 0
	var data []byte

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		size += n
		data = append(data, buf[:n]...)
	}

	return &Compiler{
		data:         data,
		size:         size,
		instructions: []Instruction{},
	}
}

func (c *Compiler) Compile() []Instruction {
	stack := []int{}

	for c.pointer < c.size {
		current := c.data[c.pointer] // Get current data

		switch current {
		case '[':
			// Emit the new STARTLOOP ("[") instruction, with 0 as Count
			stack = append(stack, c.Emit(STARTLOOP, 0))

		case ']':
			// Pop pointer of last STARTLOOP ("[") instruction off stack
			openInstruction := stack[len(stack)-1]

			// Emit the new ENDLOOP ("]") instruction, with correct pointer as Count
			// Patch the old STARTLOOP ("[") instruction with new pointer
			c.instructions[openInstruction].Count = c.Emit(ENDLOOP, openInstruction)
			stack = stack[:len(stack)-1]

		case '+':
			c.appendCharCount('+', PLUS)

		case '-':
			c.appendCharCount('-', SUB)

		case '<':
			c.appendCharCount('<', SUBDATAPOINTER)

		case '>':
			c.appendCharCount('>', PLUSDATAPOINTER)

		case '.':
			c.appendCharCount('.', WRITECHAR)

		case ',':
			c.appendCharCount(',', READCHAR)

		}

		c.pointer++
	}

	return c.instructions
}

func (c *Compiler) appendCharCount(char byte, insType InsType) {
	count := 1

	for c.pointer < c.size-1 && c.data[c.pointer+1] == char {
		count++
		c.pointer++
	}

	c.Emit(insType, count) // Type and count of this type
}

// Emit new type with Count to the instruction
func (c *Compiler) Emit(insType InsType, count int) int {
	c.instructions = append(c.instructions, Instruction{Type: insType, Count: count})
	return len(c.instructions) - 1
}

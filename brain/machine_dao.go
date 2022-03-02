package brain

import "io"

type Machine struct {
	instruction []Instruction
	size        int
	memory      []int
	pointer     int
	input       io.Reader
	output      io.Writer
	buf         []byte
}


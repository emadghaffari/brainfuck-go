package main

import "io"

func New() *Brain {
	return &Brain{}
}

// Compiler create new compiler and compile the input data
func (b *Brain) Compiler(rd io.Reader) *Brain {
	b.cp = NewCompiler(rd)
	b.cp.Compile()

	return b
}

// NewMachine create a machine for execute the brainfuck
func (b *Brain) NewMachine(in io.Reader, out io.Writer) *Brain {
	b.machine = NewMachine(b.cp.instructions, in, out)
	return b
}

// Execute
func (b *Brain) Execute() {
	b.machine.Execute()
}

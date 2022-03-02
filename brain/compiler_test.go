package brain

import (
	"bytes"
	"testing"
)

func TestCompile(t *testing.T) {
	input := bytes.NewBufferString(`+++[---[+]>>>]<<<`)
	expected := []Instruction{
		{PLUS, 3},
		{STARTLOOP, 7},
		{SUB, 3},
		{STARTLOOP, 5},
		{PLUS, 1},
		{ENDLOOP, 3},
		{PLUSDATAPOINTER, 3},
		{ENDLOOP, 1},
		{SUBDATAPOINTER, 3},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if bytecode[i] != op {
			t.Errorf("want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

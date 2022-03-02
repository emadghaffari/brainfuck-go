package brain

import (
	"bytes"
	"testing"
)

func TestIncrement(t *testing.T) {
	compiler := NewCompiler(bytes.NewBufferString("++++++++++++++++++++"))
	instructions := compiler.Compile()

	m := NewMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))
	m.Execute()

	if m.memory[0] != 20 {
		t.Errorf("increment err. got=%d", m.memory[0])
	}
}

func TestDecrement(t *testing.T) {
	input := "++++++++++-----+++++----------+----------"
	compiler := NewCompiler(bytes.NewBufferString(input))
	instructions := compiler.Compile()

	m := NewMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))
	m.Execute()

	if m.memory[0] != -9 {
		t.Errorf("decrement. got=%d", m.memory[0])
	}
}

// pointer on memory, memory value
// 2				, 3
// 1				, 2
// 0				, 1
// [{62 2} {43 3} {60 1} {43 2} {60 1} {43 1}] [1 2 3 0 0 0] ==> {62 == >}, {43 == +}, {60 == <}
func TestDecrementDataPointer(t *testing.T) {
	compiler := NewCompiler(bytes.NewBufferString(">>+++<++<+"))
	instructions := compiler.Compile()

	m := NewMachine(instructions, new(bytes.Buffer), new(bytes.Buffer))
	m.Execute()

	for i, expected := range []int{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestRead(t *testing.T) {
	in := bytes.NewBufferString("EMAD")
	out := new(bytes.Buffer)

	compiler := NewCompiler(bytes.NewBufferString(",>,>"))
	instructions := compiler.Compile()

	m := NewMachine(instructions, in, out)

	m.Execute()

	expectedMemory := []int{
		int('E'),
		int('M'),
	}

	for i, expected := range expectedMemory {
		if m.memory[i] != expected {
			t.Errorf("memory[%d] wrong value, want=%d, got=%d",
				i, expected, m.memory[0])
		}
	}
}

func TestWrite(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	compiler := NewCompiler(bytes.NewBufferString(".>.>.>.>.>"))
	instructions := compiler.Compile()

	m := NewMachine(instructions, in, out)

	setupMemory := []int{
		int('H'),
		int('E'),
		int('L'),
		int('L'),
		int('O'),
	}

	for i, value := range setupMemory {
		m.memory[i] = value
	}

	m.Execute()

	output := out.String()
	if output != "HELLO" {
		t.Errorf("output wrong. got=%q", output)
	}

}

func TestHelloWorld(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)
	txt := bytes.NewBufferString("++++++++[>++++[>++>+++>+++>+<<<<-]>+> +>->>+[<]<-]>>.>---.+++++++ ..+ ++.>>.<-.<.+++.------.--------.>>+.>++.")

	compiler := NewCompiler(txt)
	instructions := compiler.Compile()

	m := NewMachine(instructions, in, out)

	m.Execute()

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("output wrong. got=%q", output)
	}
}

package main

type InsType byte

const (
	PLUS            InsType = '+' // 43 Increment the value in the current cell.
	SUB             InsType = '-' // 45 Decrement the value in the current cell.
	WRITECHAR       InsType = '.' // 46 Take the integer in the current cell, treat it as an ASCII char and print it on the output stream.
	STARTLOOP       InsType = '[' // 91 If the current cell contains a zero, set the instruction pointer to the index of the instruction after the matching ]
	ENDLOOP         InsType = ']' // 93 If the current cell does not contain a zero, set the instruction pointer to the index of the instruction after the matching [
	SUBDATAPOINTER  InsType = '<' // 60 Decrement the data pointer by 1.
	PLUSDATAPOINTER InsType = '>' // 62 Increment the data pointer by 1.
	READCHAR        InsType = ',' // 44 Read a character from the input stream, convert it to an integer and save it to the current cell.
)

type Instruction struct {
	Type  InsType
	Count int
}

type Compiler struct {
	data         []byte
	size         int
	pointer      int
	instructions []Instruction
}

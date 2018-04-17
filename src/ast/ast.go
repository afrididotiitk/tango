package ast

import (
	"errors"
	"tango/src/codegen"
)

// Attrib represents any generic element of the ast
type Attrib interface {
}

// Node represents a node
type Node struct {
	name string
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return n.name
}

// Derivations are the all the derivations discovered in this parse
var Derivations map[*Node]Stack

func init() {
	Derivations = make(map[*Node]Stack)
}

// AddNode creates a node
func AddNode(name string, attribs ...Attrib) (Attrib, error) {
	node := &Node{
		name: name,
	}
	Derivations[node] = attribs
	return node, nil
}

// AddrCode is a struct representing the SymbolTableEntry and Code
type AddrCode struct {
	Symbol codegen.SymbolTableEntry
	Code   []codegen.IRIns
}

// ErrUnsupported is used to report unsupported errors in the code
var ErrUnsupported = errors.New("unsupported operation")

// NewSourceFile creates a source file from the decl list
func NewSourceFile(declList Attrib) (*AddrCode, error) {
	return MergeCodeList(declList)
}

var tempCount int

var predecID = []string{
	"bool", "byte", "error", "float32",
	"int", "int8", "int16", "int32", "rune", "string",
	"uint", "uint8", "uint16", "uint32", "uintptr",
}

var predecConst = map[string]int{
	"true":  1,
	"false": 0,
}

var predecFunc = []string{
	"printf",
}

func init() {
	for _, v := range predecID {
		// TODO: Make this symboltable type entry or something
		codegen.SymbolTable.InsertSymbol(v, nil)
	}
	for k, v := range predecConst {
		codegen.SymbolTable.InsertSymbol(k, &codegen.SymbolTableLiteralEntry{
			Value: v,
		})
	}
	for _, v := range predecFunc {
		codegen.SymbolTable.InsertSymbol(v, &codegen.SymbolTableTargetEntry{
			Target: v,
		})
	}
}

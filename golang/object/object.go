// Package object defines the object system used in the Monkey programming language.
package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/w40141/monkey-language/golang/ast"
)

type (
	// Type represents the type of an object.
	Type string
	// BuiltinFunction is the type of a built-in function.
	BuiltinFunction func(args ...Object) Object
)

const (
	// IntegerObj represents the type of an integer object.
	IntegerObj = "INTEGER"
	// BooleanObj represents the type of a boolean object.
	BooleanObj = "BOOLEAN"
	// NullObj represents the type of a null object.
	NullObj = "NULL"
	// ReturnObj represents the type of a return object.
	ReturnObj = "RETURN"
	// ErrorObj represents the type of an error object.
	ErrorObj = "ERROR"
	// FunctionObj represents the type of a function object.
	FunctionObj = "FUNCTION"
	// StringObj represents the type of a string object.
	StringObj = "STRING"
	// BuiltinObj represents the type of a builtin object.
	BuiltinObj = "BUILTIN"
	// ArrayObj represents the type of an array object.
	ArrayObj = "ARRAY"
	// HashObj represents the type of a hash object.
	HashObj = "HASH"
)

// Object represents an object in the Monkey programming language.
type Object interface {
	Type() Type
	Inspect() string
}

var _ Object = (*Integer)(nil)

// Integer represents an integer object in the Monkey programming language.
type Integer struct {
	Value int64
}

// Type returns the type of the integer object.
func (i *Integer) Type() Type {
	return IntegerObj
}

// Inspect returns the string representation of the integer object.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

var _ Object = (*Boolean)(nil)

// Boolean represents a boolean object in the Monkey programming language.
type Boolean struct {
	Value bool
}

// Type returns the type of the boolean object.
func (b *Boolean) Type() Type {
	return BooleanObj
}

// Inspect returns the string representation of the boolean object.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

var _ Object = (*Null)(nil)

// Null represents a null object in the Monkey programming language.
type Null struct{}

// Type returns the type of the null object.
func (n *Null) Type() Type {
	return NullObj
}

// Inspect returns the string representation of the null object.
func (n *Null) Inspect() string {
	return "null"
}

var _ Object = (*Return)(nil)

// Return represents a return object in the Monkey programming language.
type Return struct {
	Value Object
}

// Type returns the type of the return object.
func (r *Return) Type() Type {
	return ReturnObj
}

// Inspect returns the string representation of the return object.
func (r *Return) Inspect() string {
	return r.Value.Inspect()
}

var _ Object = (*Error)(nil)

// Error represents an error object in the Monkey programming language.
type Error struct {
	Message string
}

// Type returns the type of the error object.
func (e *Error) Type() Type {
	return ErrorObj
}

// Inspect returns the string representation of the error object.
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

var _ Object = (*Function)(nil)

// Function represents a function object in the Monkey programming language.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns the type of the function object.
func (f *Function) Type() Type {
	return FunctionObj
}

// Inspect returns the string representation of the function object.
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

var _ Object = (*String)(nil)

// String represents a string object in the Monkey programming language.
type String struct {
	Value string
}

// Inspect implements Object.
func (s *String) Inspect() string {
	return s.Value
}

// Type implements Object.
func (s *String) Type() Type {
	return StringObj
}

var _ Object = (*Builtin)(nil)

// Builtin represents a built-in function object in the Monkey programming language.
type Builtin struct {
	Fn BuiltinFunction
}

// Inspect implements Object.
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// Type implements Object.
func (b *Builtin) Type() Type {
	return BuiltinObj
}

var _ Object = (*Array)(nil)

// Array represents an array object in the Monkey programming language.
type Array struct {
	Elems []Object
}

// Inspect implements Object.
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elems := []string{}
	for _, e := range a.Elems {
		elems = append(elems, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")

	return out.String()
}

// Type implements Object.
func (a *Array) Type() Type {
	return ArrayObj
}

// HashPair represents a key-value pair in a hash object.
type HashPair struct {
	Key   Object
	Value Object
}

var _ Object = (*Hash)(nil)

// Hash represents a hash object in the Monkey programming language.
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Inspect implements Object.
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Type implements Object.
func (h *Hash) Type() Type {
	return HashObj
}

// Hashable represents
type Hashable interface {
	HashKey() HashKey
}

// HashKey represents a hash key in the Monkey programming language.
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey represents a hash key in the Monkey programming language.
func (b Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// HashKey represents a hash key in the Monkey programming language.
func (i Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey represents a hash key in the Monkey programming language.
func (s String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

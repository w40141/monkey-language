// Package object defines the object system used in the Monkey programming language.
package object

import "fmt"

// Type represents the type of an object.
type Type string

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
)

// Object represents an object in the Monkey programming language.
type Object interface {
	Type() Type
	Inspect() string
}

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

// Package object defines the object system used in the Monkey programming language.
package object

import "fmt"

// Type represents the type of an object.
type Type string

const (
	integerObj = "INTEGER"
	booleanObj = "BOOLEAN"
	nullObj    = "NULL"
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
	return integerObj
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
	return booleanObj
}

// Inspect returns the string representation of the boolean object.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null represents a null object in the Monkey programming language.
type Null struct{}

// Type returns the type of the null object.
func (n *Null) Type() Type {
	return nullObj
}

// Inspect returns the string representation of the null object.
func (n *Null) Inspect() string {
	return "null"
}

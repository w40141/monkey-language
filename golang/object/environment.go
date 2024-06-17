// Package object defines the object system used in the Monkey programming language.
package object

// Environment represents the environment in which the Monkey programming language is evaluated.
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates a new environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// NewEnclosedEnvironment creates a new environment enclosed in the given environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get returns the object associated with the given name from the environment.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets the object associated with the given name in the environment.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

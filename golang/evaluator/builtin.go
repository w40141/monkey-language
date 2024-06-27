// Package evaluator contains the logic for evaluating the AST nodes.
package evaluator

import "github.com/w40141/monkey-language/golang/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			return lenBuiltin(args...)
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			return firstBuiltin(args...)
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			return lastBuiltin(args...)
		},
	},
	"tail": {
		Fn: func(args ...object.Object) object.Object {
			return tailBuiltin(args...)
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			return pushBuiltin(args...)
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			return putsBuiltin(args...)
		},
	},
}

func putsBuiltin(args ...object.Object) object.Object {
	for _, arg := range args {
		println(arg.Inspect())
	}
	return nullObj
}

func pushBuiltin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elems)

	newElems := make([]object.Object, length+1)
	copy(newElems, arr.Elems)
	newElems[length] = args[1]

	return &object.Array{Elems: newElems}
}

func tailBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `tail` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elems)
	if length > 0 {
		newElems := make([]object.Object, length-1)
		copy(newElems, arr.Elems[1:length])
		return &object.Array{Elems: newElems}
	}
	return nullObj
}

// TODO: test
func lastBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elems)
	if length > 0 {
		return arr.Elems[length-1]
	}
	return nullObj
}

// TODO: test
func firstBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ArrayObj {
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	if len(arr.Elems) > 0 {
		return arr.Elems[0]
	}
	return nullObj
}

func lenBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	// TODO: test
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elems))}
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

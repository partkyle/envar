package envar

import (
	"fmt"
	"strconv"
)

type intRef struct {
	basicRef
	def int
	ref *int
}

func (i *intRef) Default() string {
	return fmt.Sprintf("%d", i.def)
}

func (i *intRef) Set(env string) {
	val, err := strconv.Atoi(env)
	if err != nil {
		*i.ref = i.def
		return
	}

	*i.ref = val
}

// Returns a reference to a int that will get parsed from the Environment.
func Int(name string, def int, usage string) *int {
	ref := new(int)
	IntVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided int reference.
func IntVar(ref *int, name string, def int, usage string) {
	iRef := &intRef{def: def, ref: ref}
	iRef.name = name
	iRef.usage = usage

	references = append(references, iRef)
}

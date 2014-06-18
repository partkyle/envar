package envar

import (
	"fmt"
	"strconv"
)

type boolRef struct {
	basicRef
	def bool
	ref *bool
}

func (b *boolRef) Default() string {
	return fmt.Sprintf("%v", b.def)
}

func (b *boolRef) Set(env string) {
	val, err := strconv.ParseBool(env)
	if err != nil {
		return
	}

	*b.ref = val
}

// Returns a reference to a bool that will get parsed from the Environment.
func Bool(name string, def bool, usage string) *bool {
	ref := new(bool)
	BoolVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided bool reference.
func BoolVar(ref *bool, name string, def bool, usage string) {
	bRef := &boolRef{def: def, ref: ref}
	bRef.name = name
	bRef.usage = usage
	*bRef.ref = def

	references = append(references, bRef)
}

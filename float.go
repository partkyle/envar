package envar

import "strconv"

type floatRef struct {
	basicRef
	def float64
	ref *float64
}

func (i *floatRef) Set(env string) {
	val, err := strconv.ParseFloat(env, 64)
	if err != nil {
		*i.ref = i.def
		return
	}

	*i.ref = val
}

// Returns a reference to a float64 that will get parsed from the Environment.
func Float(name string, def float64, usage string) *float64 {
	ref := new(float64)
	FloatVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided float64 reference.
func FloatVar(ref *float64, name string, def float64, usage string) {
	fRef := &floatRef{def: def, ref: ref}
	fRef.name = name
	fRef.usage = usage

	references = append(references, fRef)
}

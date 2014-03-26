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

func Float(name string, def float64) *float64 {
	ref := new(float64)
	FloatVar(ref, name, def)
	return ref
}

func FloatVar(ref *float64, name string, def float64) {
	iRef := &floatRef{def: def, ref: ref}
	iRef.name = name

	references = append(references, iRef)
}

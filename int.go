package envar

import "strconv"

type intRef struct {
	basicRef
	def int
	ref *int
}

func (i *intRef) Set(env string) {
	val, err := strconv.Atoi(env)
	if err != nil {
		*i.ref = i.def
		return
	}

	*i.ref = val
}

func Int(name string, def int) *int {
	ref := new(int)
	IntVar(ref, name, def)
	return ref
}

func IntVar(ref *int, name string, def int) {
	iRef := &intRef{def: def, ref: ref}
	iRef.name = name

	references = append(references, iRef)
}

package envar

type stringRef struct {
	basicRef
	def string
	ref *string
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

func String(name string, def string) *string {
	ref := new(string)
	StringVar(ref, name, def)
	return ref
}

func StringVar(ref *string, name string, def string) {
	sRef := &stringRef{def: def, ref: ref}
	sRef.name = name

	references = append(references, sRef)
}

package envar

type stringRef struct {
	basicRef
	ref *string
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

func String(name string) *string {
	ref := new(string)
	StringVar(ref, name)
	return ref
}

func StringVar(ref *string, name string) {
	sRef := &stringRef{ref: ref}
	sRef.name = name

	references = append(references, sRef)
}

package envar

type stringRef struct {
	basicRef
	ref *string
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

// Returns a reference to a string that will get set from the Environment.
func String(name string) *string {
	ref := new(string)
	StringVar(ref, name)
	return ref
}

// Assigns the value from the Environment to the provided string reference.
func StringVar(ref *string, name string) {
	sRef := &stringRef{ref: ref}
	sRef.name = name

	references = append(references, sRef)
}

package envar

type stringRef struct {
	basicRef
	ref *string
}

func (s *stringRef) Default() string {
	return ""
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

// Returns a reference to a string that will get set from the Environment.
func String(name string, usage string) *string {
	ref := new(string)
	StringVar(ref, name, usage)
	return ref
}

// Assigns the value from the Environment to the provided string reference.
func StringVar(ref *string, name string, usage string) {
	sRef := &stringRef{ref: ref}
	sRef.name = name
	sRef.usage = usage

	references = append(references, sRef)
}

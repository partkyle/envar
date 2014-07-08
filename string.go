package envar

type stringRef struct {
	basicRef
	ref *string
}

func (s *stringRef) Default() string {
	return *s.ref
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

// Returns a reference to a string that will get set from the Environment.
func String(name string, def string, usage string) *string {
	return defaultEnvSet.String(name, def, usage)
}

// Assigns the value from the Environment to the provided string reference.
func StringVar(ref *string, name string, def string, usage string) {
	defaultEnvSet.StringVar(ref, name, def, usage)
}

package envar

import "os"

type envSet struct {
	references []ref
}

func NewEnvSet() *envSet {
	return &envSet{
		references: make([]ref, 0),
	}
}

func (e *envSet) Parse() error {
	return e.ParseFromEnvironment(defaultEnvironment)
}

// Parses from the provided environment.
func (e *envSet) ParseFromEnvironment(env Environment) error {
	for _, ref := range e.references {
		val, ok := env.Get(ref.Name())
		if ok {
			ref.Set(val)
		}
	}

	return nil
}

func (e *envSet) Usage() {
	for _, ref := range e.references {
		logger.Printf("# %s", ref.Usage())
		logger.Printf("export %s=%q", ref.Name(), ref.Default())
		logger.Println()
	}

	os.Exit(127)
}

// Returns a reference to a bool that will get parsed from the Environment.
func (e *envSet) Bool(name string, def bool, usage string) *bool {
	ref := new(bool)
	e.BoolVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided bool reference.
func (e *envSet) BoolVar(ref *bool, name string, def bool, usage string) {
	bRef := &boolRef{def: def, ref: ref}
	bRef.name = name
	bRef.usage = usage
	*bRef.ref = def

	e.references = append(e.references, bRef)
}

// Returns a reference to a float64 that will get parsed from the Environment.
func (e *envSet) Float(name string, def float64, usage string) *float64 {
	ref := new(float64)
	e.FloatVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided float64 reference.
func (e *envSet) FloatVar(ref *float64, name string, def float64, usage string) {
	fRef := &floatRef{def: def, ref: ref}
	fRef.name = name
	fRef.usage = usage
	*fRef.ref = def

	e.references = append(e.references, fRef)
}

// Returns a reference to a int that will get parsed from the Environment.
func (e *envSet) Int(name string, def int, usage string) *int {
	ref := new(int)
	e.IntVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided int reference.
func (e *envSet) IntVar(ref *int, name string, def int, usage string) {
	iRef := &intRef{def: def, ref: ref}
	iRef.name = name
	iRef.usage = usage
	*iRef.ref = def

	e.references = append(e.references, iRef)
}

// Returns a reference to a string that will get set from the Environment.
func (e *envSet) String(name string, def string, usage string) *string {
	ref := new(string)
	e.StringVar(ref, name, def, usage)
	return ref
}

// Assigns the value from the Environment to the provided string reference.
func (e *envSet) StringVar(ref *string, name string, def string, usage string) {
	sRef := &stringRef{ref: ref, def: def}
	sRef.name = name
	sRef.usage = usage
	*sRef.ref = def

	e.references = append(e.references, sRef)
}

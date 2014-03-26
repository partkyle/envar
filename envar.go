package envar

var references = make([]ref, 0)

type ref interface {
	Name() string
	Set(string)
}

type basicRef struct {
	name string
}

func (b *basicRef) Name() string {
	return b.name
}

// Parses the config and loads all references using the
// default Environment, `os.Environ`.
//
// Internally uses `ParseFromEnvironment`.
func Parse() error {
	return ParseFromEnvironment(defaultEnvironment)
}

// Parses from the provided environment.
func ParseFromEnvironment(env Environment) error {
	for _, ref := range references {
		val := env.Get(ref.Name())
		ref.Set(val)
	}

	return nil
}

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

func Parse() error {
	return ParseFromEnvironment(defaultEnvironment)
}

func ParseFromEnvironment(env Environment) error {
	for _, ref := range references {
		val := env.Get(ref.Name())
		ref.Set(val)
	}

	return nil
}

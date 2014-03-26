package envar

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)
var usage = Bool("ENVAR_USAGE", false, "print usage and exit")

var references = make([]ref, 0)

type ref interface {
	Name() string
	Usage() string
	Default() string
	Set(string)
}

type basicRef struct {
	name  string
	usage string
}

func (b *basicRef) Name() string {
	return b.name
}

func (b *basicRef) Usage() string {
	return b.usage
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

	Usage()

	return nil
}

func Usage() {
	if *usage {
		for _, ref := range references {
			logger.Printf("export %s=%q # %s", ref.Name(), ref.Default(), ref.Usage())
		}

		os.Exit(127)
	}
}

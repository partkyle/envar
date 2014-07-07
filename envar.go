package envar

import (
	"errors"
	"log"
	"os"
)

var ErrUsageRequested = errors.New("envar stopped: usage requested")

var logger = log.New(os.Stdout, "", 0)

var internalEnvSet = NewEnvSet()
var defaultEnvSet = NewEnvSet()

var usage = internalEnvSet.Bool("ENVAR_USAGE", false, "print usage and exit")

func init() {
	internalEnvSet.Parse()
}

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
	err := defaultEnvSet.ParseFromEnvironment(env)
	if err != nil {
		return err
	}

	if *usage {
		Usage()

		return ErrUsageRequested
	}

	return nil
}

func Usage() {
	defaultEnvSet.Usage()
}

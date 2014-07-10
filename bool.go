package envar

import (
	"fmt"
	"strconv"
)

type boolRef struct {
	basicRef
	def bool
	ref *bool
}

func (b *boolRef) Default() string {
	return fmt.Sprintf("%v", b.def)
}

func (b *boolRef) Set(env string) error {
	val, err := strconv.ParseBool(env)
	if err != nil {
		return err
	}

	*b.ref = val
	return nil
}

// Returns a reference to a bool that will get parsed from the Environment.
func Bool(name string, def bool, usage string) *bool {
	return defaultEnvSet.Bool(name, def, usage)
}

// Assigns the value from the Environment to the provided bool reference.
func BoolVar(ref *bool, name string, def bool, usage string) {
	defaultEnvSet.BoolVar(ref, name, def, usage)
}

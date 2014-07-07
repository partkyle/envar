package envar

import (
	"fmt"
	"strconv"
)

type floatRef struct {
	basicRef
	def float64
	ref *float64
}

func (f *floatRef) Default() string {
	return fmt.Sprintf("%f", f.def)
}

func (f *floatRef) Set(env string) {
	val, err := strconv.ParseFloat(env, 64)
	if err != nil {
		return
	}

	*f.ref = val
}

// Returns a reference to a float64 that will get parsed from the Environment.
func Float(name string, def float64, usage string) *float64 {
	return defaultEnvSet.Float(name, def, usage)
}

// Assigns the value from the Environment to the provided float64 reference.
func FloatVar(ref *float64, name string, def float64, usage string) {
	defaultEnvSet.FloatVar(ref, name, def, usage)
}

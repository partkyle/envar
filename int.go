package envar

import (
	"fmt"
	"strconv"
)

type intRef struct {
	basicRef
	def int
	ref *int
}

func (i *intRef) Default() string {
	return fmt.Sprintf("%d", i.def)
}

func (i *intRef) Set(env string) {
	val, err := strconv.Atoi(env)
	if err != nil {
		return
	}

	*i.ref = val
}

// Returns a reference to a int that will get parsed from the Environment.
func Int(name string, def int, usage string) *int {
	return defaultEnvSet.Int(name, def, usage)
}

// Assigns the value from the Environment to the provided int reference.
func IntVar(ref *int, name string, def int, usage string) {
	defaultEnvSet.IntVar(ref, name, def, usage)
}

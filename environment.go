package envar

import (
	"fmt"
	"os"
	"strings"
)

var defaultEnvironment = &osEnviron{}

// An object can implement this interface in order to be used in
// `ParseFromEnvironment`
type Environment interface {
	// Retrieve a value from the environment.
	Get(string) (string, bool)
}

type osEnviron struct{}

func (o *osEnviron) Get(key string) (string, bool) {
	env := os.Environ()
	for _, value := range env {
		if strings.HasPrefix(value, fmt.Sprintf("%s=", key)) {
			return os.Getenv(key), true
		}
	}
	return "", false
}

type basicEnv map[string]string

func (b basicEnv) Get(key string) (string, bool) {
	val, ok := b[key]

	return val, ok
}

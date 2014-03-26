package envar

import "os"

var defaultEnvironment = &osEnviron{}

// An object can implement this interface in order to be used in
// `ParseFromEnvironment`
type Environment interface {
	// Retrieve a value from the environment.
	Get(string) string
}

type osEnviron struct{}

func (o *osEnviron) Get(key string) string {
	return os.Getenv(key)
}

type basicEnv map[string]string

func (b basicEnv) Get(key string) string {
	return b[key]
}

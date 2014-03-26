package envar

import "os"

var defaultEnvironment = &osEnviron{}

type Environment interface {
	// Retrieve a value from the environment.
	Get(string) string
}

type osEnviron struct{}

func (o *osEnviron) Get(key string) string {
	return os.Getenv(key)
}

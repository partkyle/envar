package envar

import (
	"os"
	"strconv"
)

var references = make([]ref, 0)

type ref interface {
	Name() string
	Set(string)
}

type intRef struct {
	name string
	def  int
	ref  *int
}

func (i *intRef) Name() string {
	return i.name
}

func (i *intRef) Set(env string) {
	val, err := strconv.Atoi(env)
	if err != nil {
		*i.ref = i.def
		return
	}

	*i.ref = val
}

func Int(name string, def int) *int {
	ref := new(int)
	IntVar(ref, name, def)
	return ref
}

func IntVar(ref *int, name string, def int) {
	iRef := &intRef{name: name, def: def, ref: ref}
	references = append(references, iRef)
}

type stringRef struct {
	name string
	def  string
	ref  *string
}

func (s *stringRef) Name() string {
	return s.name
}

func (s *stringRef) Set(env string) {
	*s.ref = env
}

func String(name string, def string) *string {
	ref := new(string)
	StringVar(ref, name, def)
	return ref
}

func StringVar(ref *string, name string, def string) {
	sRef := &stringRef{name: name, def: def, ref: ref}
	references = append(references, sRef)
}

func Parse() error {
	for _, ref := range references {
		ref.Set(os.Getenv(ref.Name()))
	}

	return nil
}

package envar

import (
	"os"
	"strconv"
)

var references = make([]ref, 0)

type ref interface {
	Name() string
	SetValue(string)
}

type intRef struct {
	name string
	def  int
	ref  *int
}

func (i *intRef) Name() string {
	return i.name
}

func (i *intRef) SetValue(env string) {
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

func Parse() error {
	for _, ref := range references {
		ref.SetValue(os.Getenv(ref.Name()))
	}

	return nil
}

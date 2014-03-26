package envar

import (
	"os"
	"strconv"
)

var references = make([]ref, 0)

type ref interface {
	Name() string
	SetValue(string) error
}

type intRef struct {
	name string
	def  int
	ref  *int
}

func (i *intRef) Name() string {
	return i.name
}

func (i *intRef) SetValue(env string) error {
	val, err := strconv.Atoi(env)
	if err != nil {
		*i.ref = i.def
		return err
	}

	*i.ref = val
	return nil
}

func Int(name string, def int) *int {
	ref := new(int)

	iRef := &intRef{name: name, def: def, ref: ref}
	references = append(references, iRef)

	return ref
}

func Parse() error {
	for _, ref := range references {
		err := ref.SetValue(os.Getenv(ref.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

package envar

import "strconv"

var references = make([]ref, 0)

type ref interface {
	Name() string
	Set(string)
}

type basicRef struct {
	name string
}

func (b *basicRef) Name() string {
	return b.name
}

type intRef struct {
	basicRef
	def int
	ref *int
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
	iRef := &intRef{def: def, ref: ref}
	iRef.name = name

	references = append(references, iRef)
}

type stringRef struct {
	basicRef
	def string
	ref *string
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
	sRef := &stringRef{def: def, ref: ref}
	sRef.name = name

	references = append(references, sRef)
}

type boolRef struct {
	basicRef
	def bool
	ref *bool
}

func (b *boolRef) Set(env string) {
	val, err := strconv.ParseBool(env)
	if err != nil {
		*b.ref = b.def
		return
	}

	*b.ref = val
}

func Bool(name string, def bool) *bool {
	ref := new(bool)
	BoolVar(ref, name, def)
	return ref
}

func BoolVar(ref *bool, name string, def bool) {
	bRef := &boolRef{def: def, ref: ref}
	bRef.name = name

	references = append(references, bRef)
}

func Parse() error {
	for _, ref := range references {
		val := defaultEnvironment.Get(ref.Name())
		ref.Set(val)
	}

	return nil
}

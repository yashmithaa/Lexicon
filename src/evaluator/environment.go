package evaluator

// environment for managing variable scopes and symbol table
type Environment struct {
	store map[string]Object
	outer *Environment
}

// creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// creates a new enclosed environment for nested scopes
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// retrieves a variable value from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// sets a variable value in the current environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// returns all variable names in the current environment (not including outer scopes)
func (e *Environment) Names() []string {
	names := make([]string, 0, len(e.store))
	for name := range e.store {
		names = append(names, name)
	}
	return names
}

// returns the store for inspection (read-only access)
func (e *Environment) GetStore() map[string]Object {
	return e.store
}

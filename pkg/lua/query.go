package lua

import (
	"fmt"

	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
)

// Environment wraps a gopher-lua environment and provides methods for running lua commands against BigTable output
type Environment struct {
	state *lua.LState
}

// NewEnvironment returns a new lua environment for evaluating snippets of Lua code
func NewEnvironment() *Environment {
	env := Environment{
		state: lua.NewState(),
	}

	return &env
}

func (env *Environment) Eval(source string) error {
	err := env.state.DoString(source)
	if err != nil {
		return errors.Wrap(err, "error executing lua code")
	}

	result := env.state.Get(-1)
	fmt.Printf("result: %v\n", result)
	return nil
}

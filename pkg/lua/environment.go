package lua

import (
	"github.com/catkins/qbt/pkg/bt"
	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
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
	luajson.Preload(env.state)

	return &env
}

// RowPredicate returns a row predicate function to filter bigtable rows
func (env *Environment) RowPredicate(source string) func(row bt.Row) (bool, error) {
	return func(row bt.Row) (bool, error) {
		// pass the row to the script as the global "row"
		env.state.SetGlobal("row", env.ConvertRowToTable(row))

		// run the lua script
		err := env.state.DoString(source)
		if err != nil {
			return false, errors.Wrap(err, "error running row predicate")
		}

		// result comes as the return value from the script on the top of the stack
		result := env.state.Get(-1)

		// only "true" if it's expicit lua true value
		return result == lua.LTrue, nil
	}
}

// ConvertRowToTable converts a BigTable row to a lua table containing
// a top level key for each column family containing a table of column name to value
// "key" is also at the top level including the BigTable row key
func (env *Environment) ConvertRowToTable(row bt.Row) *lua.LTable {
	table := env.state.NewTable()
	for cfName, cf := range row.ColumnFamilies {
		cfTable := env.state.NewTable()
		for column, value := range cf {
			cfTable.RawSetString(column, lua.LString(value))
		}

		table.RawSet(lua.LString(cfName), cfTable)
	}

	// set key last to ensure it doesn't get overwritten
	table.RawSetString("key", lua.LString(row.Key))
	return table
}

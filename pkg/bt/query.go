package bt

// Query represents a qbt query to be executed by the client
type Query struct {
	Table     string
	Predicate func(Row) (bool, error)
	Range     RowSet
}

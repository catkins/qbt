package bt

import "cloud.google.com/go/bigtable"

// RowSet is a wrapper around bigtable.RowSet
type RowSet interface {
	toRowSet() bigtable.RowSet
}

// AllRows is a RowSet which will iterate through all rows in a table
type AllRows struct{}

func (AllRows) toRowSet() bigtable.RowSet {
	return bigtable.InfiniteRange("")
}

// AllRowsFrom is a RowSet which will iterate through all rows in table with keys equal to or greater than than StartRow
type AllRowsFrom struct {
	StartRow string
}

func (rs AllRowsFrom) toRowSet() bigtable.RowSet {
	return bigtable.InfiniteRange(rs.StartRow)
}

// PrefixRange is a RowSet which will iterate all rows with keys having a given Prefix
type PrefixRange struct {
	Prefix string
}

func (rs PrefixRange) toRowSet() bigtable.RowSet {
	return bigtable.PrefixRange(rs.Prefix)
}

// Range is a RowSet which will iterate all rows with keys from Begin, but smaller than End
type Range struct {
	Begin string
	End   string
}

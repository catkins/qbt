package bt

import "cloud.google.com/go/bigtable"

// Row is a nested representation of a BigTable row
type Row struct {
	Key            string                  `json:"key"`
	ColumnFamilies map[string]ColumnFamily `json:"column_families"`
}

// ColumnFamily is a map of Cells keyed by Column name
type ColumnFamily map[string]string

// NewRowFromBigTable takes a bigtable.Row and returns a bt.Row for use in lua queries
func NewRowFromBigTable(btRow bigtable.Row) Row {
	row := Row{
		Key:            btRow.Key(),
		ColumnFamilies: make(map[string]ColumnFamily),
	}

	for cfName, readItems := range btRow {
		cf := ColumnFamily{}
		for _, readItem := range readItems {
			cf[readItem.Column] = string(readItem.Value)
		}
		row.ColumnFamilies[cfName] = cf
	}

	return row
}

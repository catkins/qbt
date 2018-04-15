package bt

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigtable"
	"github.com/pkg/errors"
)

// Client is a wrapper for the GCP BigTable client exposing domain level methods
type Client struct {
	btClient *bigtable.Client
}

// NewClient returns a new BigTable client initialised with default application credentials
func NewClient(ctx context.Context, project string, instance string) (*Client, error) {
	btClient, err := bigtable.NewClient(ctx, project, instance)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create BigTable client")
	}

	return &Client{btClient: btClient}, nil
}

// ReadRowsFiltered takes a predicate function and invokes given callback with rows that match given predicate
func (client Client) ReadRowsFiltered(ctx context.Context, query Query, callback func(Row)) error {
	table := client.btClient.Open(query.Table)

	err := table.ReadRows(ctx, query.Range.toRowSet(), func(btRow bigtable.Row) bool {
		row := NewRowFromBigTable(btRow)

		rowMatches, err := query.Predicate(row)
		if err != nil {
			fmt.Println(err) // TODO: handle error or push into callback
		}

		if rowMatches {
			callback(row)
		}

		return true
	})

	return err
}

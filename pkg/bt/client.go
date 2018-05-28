package bt

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigtable"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// Client is a wrapper for the GCP BigTable client exposing domain level methods
type Client struct {
	btClient *bigtable.Client
}

// NewClient returns a new BigTable client initialised with default application credentials
func NewClient(ctx context.Context, project string, instance string, clientOptions ...ClientOption) (*Client, error) {
	conf := clientConfig{
		project:  project,
		instance: instance,
	}

	for _, opt := range clientOptions {
		opt(&conf)
	}

	gcpOpts := []option.ClientOption{}
	if conf.emulatorAddr != "" {
		fmt.Fprintf(os.Stderr, "connecting to emulator at %q\n", conf.emulatorAddr)
		conn, err := grpc.Dial(conf.emulatorAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalln(err)
		}
		gcpOpts = append(gcpOpts, option.WithGRPCConn(conn))
	}

	btClient, err := bigtable.NewClient(ctx, project, instance, gcpOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create BigTable client")
	}

	return &Client{btClient: btClient}, nil
}

// WithEmulator allows connecting to a BigTable emulator instead of a real BigTable cluster
//
// see: godoc .... /bttest
func WithEmulator(emulatorAddr string) ClientOption {
	return func(conf *clientConfig) {
		conf.emulatorAddr = emulatorAddr
	}
}

type clientConfig struct {
	project      string
	instance     string
	emulatorAddr string
}

type ClientOption func(conf *clientConfig)

// ReadRowsFiltered takes a predicate function and invokes given callback with rows that match given predicate
func (client Client) ReadRowsFiltered(ctx context.Context, query Query, callback func(Row)) error {
	table := client.btClient.Open(query.Table)

	err := table.ReadRows(ctx, query.Range.toRowSet(), func(btRow bigtable.Row) bool {
		row := NewRowFromBigTable(btRow)

		rowMatches, err := query.Predicate(row)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error running predicate", err) // TODO: handle error or push into callback
		}

		if rowMatches {
			callback(row)
		}

		return true
	})

	return err
}

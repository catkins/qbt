QBT - Query BigTable
--------------------

QBT is a small CLI application for querying Google Cloud BigTable, using the Lua programming language to express queries.

### Why

[`cbt`](https://cloud.google.com/bigtable/docs/go/cbt-overview) is a great tool for managing BigTable tables and column familes, but it's `read` subcomand is somewhat limited in that it doesn't allow you to provide any filters or transformations to your data.

The other option for ad-hoc queries into BigTable is to write custom scripts in your language of choice, which is quite time-consuming for exploratory queries.

### Installation

#### Requirements

- Go 1.10

```
go get -u github.com/catkins/qbt
```

### Usage

```
qbt [--project=<project id>] [--instance=<instance id>] <table>
```

### Configuration

Command line options can also be provided by a config file using the `--config` flag, or a default config at `~/.config/qbt.json`.

### Licence

Copyright 2018 Chris Atkins

MIT

QBT - Query BigTable
--------------------

QBT is a small CLI application for querying Google Cloud BigTable, using the Lua programming language to express queries.

### Why

[`cbt`](https://cloud.google.com/bigtable/docs/go/cbt-overview) is a great tool for managing BigTable tables and column familes, but it's `read` subcomand is somewhat limited in that it doesn't allow you to provide any filters or transformations to your data.

The other option for ad-hoc queries into BigTable is to write custom scripts in your language of choice, which is quite time-consuming for exploratory queries.

### Installation

#### Requirements

- Go 1.9+
- gcloud CLI

```
go get -u github.com/catkins/qbt
gcloud auth applicatiom-default login
```

### Usage

```
qbt [--project=<project id>] [--instance=<instance id>] [--prefix=<row key prefix>] query <table> <query>
```

### Example

```sh
qbt query my_table 'return row.my_cf["my_cf:haystack"] == "needle"'

# json decoding also included
qbt query my_table '
  json = require("json")
  decoded_value = json.decode(row.my_cf["my_cf:haystack"])
  return decoded_value["nested_field"] == 123
'
```

### Configuration

Configuration can also be provided from the environment by providing the `QBT_PROJECT`, `QBT_INSTANCE`, and `QBT_PREFIX` environment variables, or via CLI flags.

### Libraries used

- GCP client libraries for Go: cloud.google.com/go
- Cobra for CLI: github.com/spf13/cobra
- Viper for configuration: github.com/spf13/viper
- Lua implementation gopher-lua: github.com/yuin/gopher-lua
- JSON support for gopher-lua: layeh.com/gopher-json
- Errors: github.com/pkg/errors

### Todo / Ideas

- [ ] Tests!
- [ ] Add result transformations
- [ ] Add other output types (eg. YAML, CSV, strings returned from scripts)
- [ ] Add debug logging
- [ ] Queries from file / STDIN
- [ ] Pre-query scripts (eg. to define functions)
- [x] Add CLI params for specifying range / prefix queries
- [ ] Web server for submitting queries to
- [ ] Other authentication methods

### Licence

Copyright 2018 Chris Atkins

MIT

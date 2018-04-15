QBT - Query BigTable
--------------------

QBT is a small CLI application for querying Google Cloud BigTable, using the Lua programming language to express queries.

### Why

[`cbt`](https://cloud.google.com/bigtable/docs/go/cbt-overview) is a great tool for managing BigTable tables and column familes, but it's `read` subcomand is somewhat limited in that it doesn't allow you to provide any filters or transformations to your data.

The other option for ad-hoc queries into BigTable is to write custom scripts in your language of choice, which is quite time-consuming for exploratory queries.

### Installation

#### Requirements

- Go 1.9+

```
go get -u github.com/catkins/qbt
```

### Usage

```
qbt [--project=<project id>] [--instance=<instance id>] <table> <query>
```

### Example

```sh
qbt query --project=my-project --instance=my-bt-instance my_table 'if row.my_cf["my_cf:haystack"] == "needle" then return true end'

# json decoding also included
qbt query --project=my-project --instance=my-bt-instance my_table '
  json = require("json")
  decoded_value = json.decode(row.my_cf["my_cf:haystack"])
  return decoded_value["nested_field"] == 123
'
```

### Configuration

Configuration can also be provided from the environment by providing the `QBT_PROJECT` and `QBT_INSTANCE` environment variables.

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
- [ ] Add CLI params for specifying range / prefix queries
- [ ] Web server for submitting queries to
- [ ] Other authentication methods

### Licence

Copyright 2018 Chris Atkins

MIT

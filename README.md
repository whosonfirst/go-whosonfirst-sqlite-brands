# go-whosonfirst-sqlite-brands

Go package for working with Who's On First brands and SQLite databases.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tables

### brands

```
CREATE TABLE brands (
       id INTEGER NOT NULL,
       name TEXT,
       size TEXT,
       is_current INTEGER,
       lastmodified INTEGER
);

CREATE INDEX brands_by_name ON brands (name, size, is_current);
CREATE INDEX brands_by_name_current ON brands (name, is_current);	
CREATE INDEX brands_by_lastmod ON brands (lastmodified);
CREATE INDEX brands_by_id ON brands (id);
```

## Tools

### wof-sqlite-index-brands

```
./bin/wof-sqlite-index-brands -h
Usage of ./bin/wof-sqlite-index-brands:
  -driver string
    	 (default "sqlite3")
  -dsn string
    	 (default ":memory:")
  -live-hard-die-fast
    	Enable various performance-related pragmas at the expense of possible (unlikely) database corruption
  -mode string
    	The mode to use importing data. Valid modes are: directory,feature,feature-collection,files,geojson-ls,meta,path,repo,sqlite. (default "files")
  -processes int
    	The number of concurrent processes to index data with (default 8)
  -timings
    	Display timings during and after indexing
```

### wof-sqlite-query-brands

Query a SQLite database full of brands (that was created by `wof-sqlite-index-brands`).

```
./bin/wof-sqlite-query-brands -h
Usage of ./bin/wof-sqlite-query-brands:
  -column string
    	 (default "name")
  -driver string
    	 (default "sqlite3")
  -dsn string
    	 (default "index.db")
  -is-ceased string
    	A comma-separated list of valid existential flags (-1,0,1) to filter results according to whether or not they have been marked as ceased. Multiple flags are evaluated as a nested 'OR' query.
  -is-current string
    	A comma-separated list of valid existential flags (-1,0,1) to filter results according to their 'mz:is_current' property. Multiple flags are evaluated as a nested 'OR' query.
  -is-deprecated string
    	A comma-separated list of valid existential flags (-1,0,1) to filter results according to whether or not they have been marked as deprecated. Multiple flags are evaluated as a nested 'OR' query.
  -is-superseded string
    	A comma-separated list of valid existential flags (-1,0,1) to filter results according to whether or not they have been marked as superseded. Multiple flags are evaluated as a nested 'OR' query.
  -table string
    	 (default "search")
```

For example:

```
./bin/wof-sqlite-query-brands -dsn test.db 'car* bank'
17:23:40.459620 [wof-sqlite-query-brands] STATUS car* bank - 1125154403 Carolina First Bank
17:23:40.459746 [wof-sqlite-query-brands] STATUS car* bank - 1125153083 Central Carolina Bank
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-sqlite
* https://github.com/whosonfirst-data/whosonfirst-brands

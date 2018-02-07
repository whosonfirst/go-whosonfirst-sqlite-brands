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

## See also

* https://github.com/whosonfirst/go-whosonfirst-sqlite
* https://github.com/whosonfirst-data/whosonfirst-brands

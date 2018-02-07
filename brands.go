package brands

import (
       wof_brands "github.com/whosonfirst/go-whosonfirst-brands"
       "github.com/whosonfirst/go-whosonfirst-sqlite"
)

type BrandTable interface {
     sqlite.Table
     IndexBrand(sqlite.Database, wof_brands.Brand) error
}


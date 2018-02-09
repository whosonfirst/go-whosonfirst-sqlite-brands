package tables

// see the way "go-whosonfirst-sqlite-brands is `brands` and
// "github.com/whosonfirst/go-whosonfirst-brands" is `wof_brands`
// maybe this could be better... (20180206/thisisaaronland)

import (
	"fmt"
	wof_brands "github.com/whosonfirst/go-whosonfirst-brands"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite-brands"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	_ "log"
)

type BrandsSourceTable struct {
	brands.BrandTable
	name string
}

func NewBrandsSourceTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewBrandsSourceTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewBrandsSourceTable() (sqlite.Table, error) {

	t := BrandsSourceTable{
		name: "source",
	}

	return &t, nil
}

func (t *BrandsSourceTable) Name() string {
	return t.name
}

func (t *BrandsSourceTable) Schema() string {

	sql := `CREATE TABLE %s (
	       id INTEGER NOT NULL PRIMARY KEY,
	       body TEXT,
	       lastmodified INTEGER
	);

	CREATE INDEX %s_by_lastmod ON %s (lastmodified);
	`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name())
}

func (t *BrandsSourceTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *BrandsSourceTable) IndexRecord(db sqlite.Database, i interface{}) error {
	return t.IndexBrand(db, i.(wof_brands.Brand))
}

func (t *BrandsSourceTable) IndexBrand(db sqlite.Database, b wof_brands.Brand) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	id := b.Id()
	lastmod := b.LastModified()

	body := string(b.Bytes())

	sql := fmt.Sprintf(`INSERT OR REPLACE INTO %s (
		id, body, lastmodified
	) VALUES (
		?, ?, ?
	)`, t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, body, lastmod)

	if err != nil {
		return err
	}

	return tx.Commit()
}

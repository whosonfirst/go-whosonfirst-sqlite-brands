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

type BrandsSearchTable struct {
	brands.BrandTable
	name string
}

func NewBrandsSearchTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewBrandsSearchTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewBrandsSearchTable() (sqlite.Table, error) {

	t := BrandsSearchTable{
		name: "search",
	}

	return &t, nil
}

func (t *BrandsSearchTable) Name() string {
	return t.name
}

func (t *BrandsSearchTable) Schema() string {

	schema := `CREATE VIRTUAL TABLE %s USING fts4(id, name, is_current);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(schema, t.Name())
}

func (t *BrandsSearchTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *BrandsSearchTable) IndexRecord(db sqlite.Database, i interface{}) error {
	err := t.IndexBrand(db, i.(wof_brands.Brand))
	return err
}

func (t *BrandsSearchTable) IndexBrand(db sqlite.Database, b wof_brands.Brand) error {

	id := b.Id()
	name := b.Name()

	is_current, err := b.IsCurrent()

	if err != nil {
		return err
	}

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	s, err := tx.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id = ?", t.Name()))

	if err != nil {
		return err
	}

	defer s.Close()

	_, err = s.Exec(id)

	if err != nil {
		return err
	}

	sql := fmt.Sprintf(`INSERT OR REPLACE INTO %s (id, name, is_current) VALUES (?, ?, ?)`, t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, name, is_current.Flag())

	if err != nil {
		return err
	}

	return tx.Commit()
}

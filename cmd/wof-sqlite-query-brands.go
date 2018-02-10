package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-log"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"io"
	"os"
	"strings"
)

func main() {

	driver := flag.String("driver", "sqlite3", "")
	var dsn = flag.String("dsn", "index.db", "")

	var table = flag.String("table", "search", "")
	var col = flag.String("column", "name", "")

	flag.Parse()

	logger := log.SimpleWOFLogger()

	stdout := io.Writer(os.Stdout)
	logger.AddLogger(stdout, "status")

	db, err := database.NewDBWithDriver(*driver, *dsn)

	if err != nil {
		logger.Fatal("unable to create database (%s) because %s", *dsn, err)
	}

	defer db.Close()

	conn, err := db.Conn()

	if err != nil {
		logger.Fatal("CONN", err)
	}

	match := fmt.Sprintf("%s MATCH ?", *col)
	query := strings.Join(flag.Args(), " ")

	conditions := []string{
		match,
	}

	args := []interface{}{
		query,
	}

	// conditions = append(conditions, "is_current LIKE ?")
	// args = append(args, "1")

	where := strings.Join(conditions, " AND ")

	sql := fmt.Sprintf("SELECT id,name FROM %s WHERE %s", *table, where)
	rows, err := conn.Query(sql, args...)

	if err != nil {
		logger.Fatal("QUERY", err)
	}

	defer rows.Close()

	logger.Status("# %s", sql)

	for rows.Next() {

		var id string
		var name string

		err = rows.Scan(&id, &name)

		if err != nil {
			logger.Fatal("ID", err)
		}

		logger.Status("%s %s", id, name)
	}

	err = rows.Err()

	if err != nil {
		logger.Fatal("ROWS", err)
	}

}

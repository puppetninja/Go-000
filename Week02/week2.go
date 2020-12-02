package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func main() {
	err := dao_write()
	if err != nil {
		fmt.Printf("Failed to write in DAO with %v\n", errors.Cause(err))
		os.Exit(1)
	}

	msg, err := dao_find_latest_userid("puppetninja")
	// Do corresponding check
}

func dao_write() error {
	db, err := sql.Open("sqlite3", "./week02.db")
	if err != nil {
		return errors.Wrap(err, "Faild to connect to db")
	}
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	if err != nil {
		return errors.Wrap(err, "Faild to prepare the execution")
	}

	res, err := stmt.Exec("puppetninja", "研发部门", "2020-12-03")
	if err != nil {
		return errors.Wrap(err, "Faild to insert into db")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "No last insert id found")
	}

	fmt.Println(id)
	defer db.Close()
	return nil
}

func dao_find_latest_userid(name string) (string, error) {
	db, err := sql.Open("sqlite3", "./week02.db")
	if err != nil {
		return "", errors.Wrap(err, "Faild to connect to db")
	}

	rows, err := db.Query("SELECT id, username FROM userinfo WHERE username = ?", name)
	switch{
	case err == sql.ErrNoRows:
		return "", errors.New("no user with such username could be found")
	case err != nil:
		return "", errors.Wrap(err, "query error hit")
	default:
		for rows.Next() {
			var uid int
			var username string
			var department string
			var created time.Time
            err = rows.Scan(&uid, &username, &department, &created)
			if err != nil {
				fmt.Printf("hit error")
			}
            fmt.Println(uid)
            fmt.Println(username)
            fmt.Println(department)
            fmt.Println(created)
        }
	}
	return "", nil
}

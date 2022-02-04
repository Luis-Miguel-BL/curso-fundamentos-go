package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?, ?)")

	stmt.Exec(2000, "bia")
	aa, _ := stmt.Exec(2001, "carlos")
	fmt.Println(aa)

	_, err = stmt.Exec(5, "tiago")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

}

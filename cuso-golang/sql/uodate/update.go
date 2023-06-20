package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Wellington", 1)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(1)
}

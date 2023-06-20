package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Query("select * from usuarios")
	defer stmt.Close()

	for stmt.Next() {
		var u usuario
		stmt.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}

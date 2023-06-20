package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(res http.ResponseWriter, req *http.Request) {
	sid := strings.TrimPrefix(req.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case req.Method == "GET" && id > 0:
		usuarioPorID(res, req, id)
	case req.Method == "GET":
		usuarioTodos(res, req)
	default:
		res.WriteHeader(http.StatusNotFound)
	}

}

func usuarioTodos(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	json, _ := json.Marshal(usuarios)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(json))
}

func usuarioPorID(res http.ResponseWriter, req *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)

	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(json))
}

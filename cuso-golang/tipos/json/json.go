package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct for json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json for struct
	var p2 produto
	jsonString := `{"id":2, "nome":"Geladeira","preco":8.90,"tags": ["Eletrodomestico","Cozinha"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}

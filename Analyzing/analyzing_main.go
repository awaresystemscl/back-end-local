package main

import (
    // "fmt"
)
type mashup struct {
    id int
    limite int
}
func main() {
	var mashups []mashup
    mashups = append(mashups, mashup{id:1, limite:80})
    mashups = append(mashups, mashup{id:2, limite:80})
    conjuntos := getConjuntoMashups()
    var violados []int // Violados !!!
	for _, m := range mashups{
		for _, c := range conjuntos{
			if c.mashup_id == m.id && c.promedio < m.limite{
        		violados = append(violados, m.id)
			}
		}
		
    }
    setAlerta(violados)
}
package main

import (
    "fmt"
)
type mashup struct {
    id int
    umbral int
}
var tolerancia int = 3
func main() {
    mashups := getMashups()
    var violados []int
    for _, m := range mashups{
        conjuntos := getConjuntoMashups(m.id)
        var alerta int = 0
        for _, c := range conjuntos{
            if c.promedio < m.umbral{
                alerta = alerta +1
            }
        }
        if alerta > tolerancia{
            violados = append(violados, m.id)
        }
    }
    setAlerta(violados)
    for _, v := range violados{
        fmt.Println(v)
    }

}
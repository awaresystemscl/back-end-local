package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type mashup struct {
    id int
    umbral int
}

type configDB struct {
    Host string `json:"HOST_IP"` 
    User string `json:"DB_USER"`
    Pass string `json:"DB_PASSWORD"`
    Name string `json:"DB_NAME"`
    Tolerancia int `json:"Tolerancia"`
    Ventana int `json:"Ventana"`
}

var db_config configDB

func main() {
    db_config = configuracion()
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
        if alerta > db_config.Tolerancia{
            violados = append(violados, m.id)
        }
    }
    setAlerta(violados)
    for _, v := range violados{
        fmt.Println(v)
    }

}

func configuracion() configDB{
    jsonFile, err := ioutil.ReadFile("../config.json")
    checkErr(err)
    var config configDB
    json.Unmarshal(jsonFile, &config)
    return config
}
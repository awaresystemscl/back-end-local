
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
    // "strconv"
)

const (
    DB_USER     = "workshop"
    DB_PASSWORD = "workshop2017"
    DB_NAME     = "workshop"
)

type conjunto_mashup_sati struct {
    mashup_id int
    promedio int
    usuario_id int
}

func getConjuntoMashups() []conjunto_mashup_sati{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT mashup_id, avg, usuario_id FROM conjunto_satisfaccion_mashup")
    checkErr(err)

    var conjunto []conjunto_mashup_sati
    for rows.Next() {
        var mashup_idT int
	    var promedioT int
	    var usuario_idT int
        err = rows.Scan(&mashup_idT,&promedioT,&usuario_idT)
        checkErr(err)
        temporal := conjunto_mashup_sati{mashup_id: mashup_idT, promedio: promedioT, usuario_id: usuario_idT}
        conjunto = append(conjunto, temporal)
    }
    return conjunto
}

func setAlerta(violados []int) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    for _, v := range violados{
	    _, err = db.Exec("INSERT INTO alertas (mashup_id, fecha) VALUES($1,$2)", v, time.Now())
	    checkErr(err)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
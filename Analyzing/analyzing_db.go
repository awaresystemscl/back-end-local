
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
    "strconv"
)

const (
    DB_USER     = "awaresystems"
    DB_PASSWORD = "3ee798d8"
    DB_NAME     = "awaresystems"
)

type conjunto_mashup_sati struct {
    mashup_id int
    promedio int
    usuario_id int
}

func getConjuntoMashups(mashup_id int) []conjunto_mashup_sati{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    // rows, err := db.Query("SELECT mashup_id, avg, usuario_id FROM conjunto_satisfaccion_mashup"+
    //                     "where to_char(fecha,'DD-MM-YYYY') = to_char((select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY') "+
    //                     "from conjunto_satisfaccion_mashup order by fecha desc limit 1),'DD-MM-YYYY')")
    rows, err := db.Query("SELECT mashup_id, avg, usuario_id FROM conjunto_satisfaccion_mashup "+
                        "where conjunto_satisfaccion_mashup.mashup_id = "+strconv.Itoa(mashup_id)+
                        " limit 10")
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

func getMashups() []mashup{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT id, limite FROM mashups")
    checkErr(err)

    var mashups []mashup
    for rows.Next() {
        var mashupT mashup
        err = rows.Scan(&mashupT.id,&mashupT.umbral)
        checkErr(err)
        mashups = append(mashups, mashupT)
    }
    return mashups
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
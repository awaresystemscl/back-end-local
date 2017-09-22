package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    // "time"
    "strconv"
)

const (
    DB_USER     = "workshop"
    DB_PASSWORD = "workshop2017"
    DB_NAME     = "workshop"
)

type usuarioMashup struct {
    mashup_id int
    mashup_nombre string
    mashup_limite int
    nombre string
    email string
    promedio int
    componentes []componente
}

type componente struct {
    id int
    categoria string
    nombre string
    promedio int
    restricciones []restriccion
}

type restriccion struct {
    factor string
    nivel string
    tendencia bool
    satisfaccion int
}

func getAlertas() []usuarioMashup{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    // rows, err := db.Query("SELECT mashup_id from alertas")
    rows, err := db.Query("SELECT alertas.mashup_id, mashups.nombre, mashups.limite, usuarios.nombre, usuarios.email, conjunto_satisfaccion_mashup.avg"+
    					" from alertas join mashups on mashups.id = alertas.mashup_id "+
    					"join usuarios on usuarios.id = mashups.usuario_id "+
    					"join conjunto_satisfaccion_mashup on mashups.id = conjunto_satisfaccion_mashup.mashup_id")

    checkErr(err)

    var alertas []usuarioMashup
    for rows.Next() {
        var mashup_idT int
        var mashup_nombreT string
        var mashup_limiteT int
        var nombreT string
        var emailT string
        var promedioT int
        err = rows.Scan(&mashup_idT,&mashup_nombreT,&mashup_limiteT,&nombreT,&emailT,&promedioT)
        temporal := usuarioMashup{mashup_id:mashup_idT, mashup_nombre:mashup_nombreT,mashup_limite:mashup_limiteT, nombre:nombreT, email:emailT, promedio:promedioT}
        checkErr(err)
        alertas = append(alertas, temporal)
    }
    return alertas
}

func getComponentesMashup(idMashup int) []componente{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    // rows, err := db.Query("SELECT mashup_id from alertas")
    rows, err := db.Query("select componentes.id, componentes.categoria, apis.nombre, conjunto_satisfaccion_compo.avg"+
    					" from componentes join apis on componentes.api_id = apis.id "+
    					"join conjunto_satisfaccion_compo on componentes.id = conjunto_satisfaccion_compo.componente_id "+
    					"where componentes.mashup_id = "+strconv.Itoa(idMashup))
    checkErr(err)

    var componentes []componente
    for rows.Next() {
        var idT int
        var categoriaT string
        var nombreT string
        var promedioT int
        err = rows.Scan(&idT,&categoriaT,&nombreT,&promedioT)
        temporal := componente{id:idT, categoria:categoriaT, nombre:nombreT, promedio:promedioT}
        checkErr(err)
        componentes = append(componentes, temporal)
    }
    return componentes
}

func getRestricciones(idComponente int) []restriccion{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    // rows, err := db.Query("SELECT mashup_id from alertas")
    rows, err := db.Query("select factores.nombre, relacion_com_fac.nivel, relacion_com_fac.tendencia, satisfaccion_componente.satisfaccion"+
    					" from relacion_com_fac join factores on relacion_com_fac.factor_id = factores.id "+
    					"join satisfaccion_componente "+
    					"on satisfaccion_componente.factor = factor_id and satisfaccion_componente.componente_id = "+strconv.Itoa(idComponente)+
    					"where relacion_com_fac.componente_id = "+strconv.Itoa(idComponente))
    checkErr(err)

    var restricciones []restriccion
    for rows.Next() {
        var factorT string
        var nivelT string
        var tendenciaT bool
        var satisfaccionT int
        err = rows.Scan(&factorT,&nivelT,&tendenciaT,&satisfaccionT)
        temporal := restriccion{factor:factorT, nivel:nivelT, tendencia:tendenciaT, satisfaccion:satisfaccionT}
        checkErr(err)
        restricciones = append(restricciones, temporal)
    }
    return restricciones
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }}

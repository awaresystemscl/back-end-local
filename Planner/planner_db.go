package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    // "time"
    "strconv"
)

const (
    DB_USER     = "awaresystems"
    DB_PASSWORD = "3ee798d8"
    DB_NAME     = "awaresystems"
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
    rows, err := db.Query("SELECT DISTINCT on (alertas.mashup_id) "+
                        "alertas.mashup_id, mashups.nombre, mashups.limite, usuarios.nombre, "+
                        "usuarios.email, conjunto_satisfaccion_mashup.avg "+
    					"from alertas join mashups on mashups.id = alertas.mashup_id "+
    					"join usuarios on usuarios.id = mashups.usuario_id "+
    					"join conjunto_satisfaccion_mashup on mashups.id = conjunto_satisfaccion_mashup.mashup_id "+
                        "where to_char(alertas.fecha,'DD-MM-YYYY') = to_char((select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY') "+
                        "from alertas order by fecha desc limit 1),'DD-MM-YYYY') "+
                        "and to_char(alertas.fecha,'DD-MM-YYYY') = to_char(conjunto_satisfaccion_mashup.fecha,'DD-MM-YYYY')")

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
    					"where componentes.mashup_id = "+strconv.Itoa(idMashup)+
                        "and to_char(conjunto_satisfaccion_compo.fecha,'DD-MM-YYYY') = "+
                        "to_char((select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY') "+
                        "from conjunto_satisfaccion_compo order by fecha desc limit 1),'DD-MM-YYYY')" )
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
    rows, err := db.Query("select factores.nombre, relacion_com_fac.nivel, relacion_com_fac.tendencia,"+
                        " satisfaccion_componente.satisfaccion"+
    					" from relacion_com_fac join factores on relacion_com_fac.factor_id = factores.id "+
    					"join satisfaccion_componente "+
    					"on satisfaccion_componente.factor_id = relacion_com_fac.factor_id and satisfaccion_componente.componente_id = "+strconv.Itoa(idComponente)+
    					"where relacion_com_fac.componente_id = "+strconv.Itoa(idComponente)+
                        " and to_char(satisfaccion_componente.fecha,'DD-MM-YYYY') = "+
                        "to_char((select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY') "+
                        "from satisfaccion_componente order by fecha desc limit 1),'DD-MM-YYYY')")
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

package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
    "strconv"
)

const (
    DB_USER     = "workshop"
    DB_PASSWORD = "workshop2017"
    DB_NAME     = "workshop"
)

type apiTest struct {
    rendimiento int
    latencia int
    tiempo_de_respuesta int
    confiabilidad int
    disponibilidad int
}

type compoRule struct {
    nombre int
    tipoDeMedida string
    nivel string
    tendencia bool
}

type catCuartil struct {
    nivel_factor string
    minimo int
    medio int
    maximo int
    factor int
}

type componente struct {
    id int
    descripcion string
    categoria string
    url string
    api_id int
    mashup_id int
}

type satisfaccion_componente struct {
    componente_id int
    satisfaccion int
    factor int
    mashup_id int
}

type conjunto_compo_sati struct {
    componente_id int
    promedio int
    mashup_id int
}

type conjunto_mashup_sati struct {
    mashup_id int
    promedio int
    usuario_id int
}

func getApiTest( componenteId int) (apiTest){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select qos.rendimiento, qos.latencia, qos.tiempo_de_respuesta, qos.confiabilidad, qos.disponibilidad "+
                "from componentes "+
                "left join apis "+
                "on componentes.api_id = apis.id "+ 
                "join qos "+
                "on apis.nombre = qos.nombre "+
                "where componentes.id = "+strconv.Itoa(componenteId)+
                " order by qos.fecha asc "+
                "limit 1"
    var test apiTest
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var rendimientoT int
        var latenciaT int
        var tiempo_de_respuestaT int
        var confiabilidadT int
        var disponibilidadT int
        err = rows.Scan(&rendimientoT, &latenciaT, &tiempo_de_respuestaT, &confiabilidadT, &disponibilidadT)
        t := apiTest{rendimiento: rendimientoT, latencia: latenciaT,
                        tiempo_de_respuesta: tiempo_de_respuestaT, confiabilidad: confiabilidadT, disponibilidad: disponibilidadT}
        test = t
        checkErr(err)
    }
    return test
}

func getComponentRules(componenteId int) ([]compoRule){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select relacion_com_fac.factor_id, factores.tipo_de_medida, relacion_com_fac.nivel , relacion_com_fac.tendencia "+
                "from relacion_com_fac "+
                "join factores "+
                "on relacion_com_fac.factor_id = factores.id "+
                "where relacion_com_fac.componente_id = "+strconv.Itoa(componenteId)
    var rules []compoRule
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var nombreT int
        var tipoDeMedidaT string
        var nivelT string
        var tendenciaT bool
        err = rows.Scan(&nombreT, &tipoDeMedidaT, &nivelT, &tendenciaT)
        temporal := compoRule{nombre: nombreT, tipoDeMedida: tipoDeMedidaT,
                        nivel: nivelT, tendencia: tendenciaT}
        rules = append(rules, temporal)
        checkErr(err)
    }
    return rules
}

func getCatCuartil(componenteId int) ([]catCuartil){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select nivel_factor, minimo, medio, maximo, factor from estadistica_percentil "+
                "where estadistica_percentil.categoria = "+
                "(select categoria from componentes where id = "+strconv.Itoa(componenteId)+")"
    var cuartil []catCuartil
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var nivel_factorT string
        var minimoT int
        var medioT int
        var maximoT int
        var factorT int
        err = rows.Scan(&nivel_factorT, &minimoT, &medioT, &maximoT, &factorT)
        temporal := catCuartil{nivel_factor: nivel_factorT, minimo: minimoT,
                        medio: medioT, maximo: maximoT, factor: factorT}
        cuartil = append(cuartil, temporal)
        checkErr(err)
    }
    return cuartil
}

func getComponentes() []componente{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT id, mashup_id FROM componentes")
    checkErr(err)

    var componentes []componente
    for rows.Next() {
        var idT int
        var mashup_idT int
        err = rows.Scan(&idT,&mashup_idT)
        checkErr(err)
        temporal := componente{id: idT, mashup_id: mashup_idT}
        componentes = append(componentes, temporal)
    }
    return componentes
}

func setSatisfaccionCompo(satisfaccionComp satisfaccion_componente) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    _, err = db.Exec("INSERT INTO satisfaccion_componente (componente_id, satisfaccion, fecha, factor, mashup_id) VALUES($1,$2,$3,$4,$5)",
                    satisfaccionComp.componente_id, satisfaccionComp.satisfaccion, time.Now(), satisfaccionComp.factor, satisfaccionComp.mashup_id)
    checkErr(err)
}

func getSatisfaccion() ([]conjunto_compo_sati){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select componentes.id , satisfaccion.avg, componentes.mashup_id "+
                "from (select componente_id, avg(satisfaccion) from satisfaccion_componente group by componente_id) as satisfaccion "+
                "join componentes "+
                "on componentes.id = satisfaccion.componente_id"
    var conjuntos []conjunto_compo_sati
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var componente_idT int
        var promedioT float64
        var mashup_idT int
        err = rows.Scan(&componente_idT, &promedioT, &mashup_idT)
        temporal := conjunto_compo_sati{componente_id: componente_idT, promedio: int(promedioT), mashup_id: mashup_idT}
        conjuntos = append(conjuntos, temporal)
        checkErr(err)
    }
    return conjuntos
}

func setConjuntoCompo(conjuntos []conjunto_compo_sati) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    for _, conjunto := range conjuntos{
        _, err = db.Exec("INSERT INTO conjunto_satisfaccion_compo "+ 
                        "(componente_id, avg, mashup_id, fecha) VALUES($1,$2,$3,$4)",
                        conjunto.componente_id, conjunto.promedio, conjunto.mashup_id, time.Now())
        checkErr(err)
    }
}

func getConjuntoCompo() ([]conjunto_mashup_sati){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select satisfaccion.mashup_id , satisfaccion.avg, mashups.usuario_id "+
                "from (select mashup_id, avg(conjunto_satisfaccion_compo.avg) from conjunto_satisfaccion_compo group by mashup_id) as satisfaccion "+
                "join mashups "+
                "on mashups.id = satisfaccion.mashup_id"
    var conjuntos []conjunto_mashup_sati
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var mashup_idT int
        var promedioT float64
        var usuario_idT int
        err = rows.Scan(&mashup_idT, &promedioT, &usuario_idT)
        temporal := conjunto_mashup_sati{mashup_id: mashup_idT, promedio: int(promedioT), usuario_id: usuario_idT}
        conjuntos = append(conjuntos, temporal)
        checkErr(err)
    }
    return conjuntos
}

func setConjuntoMashup(conjuntosMashup []conjunto_mashup_sati) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    for _, conjunto := range conjuntosMashup{
        _, err = db.Exec("INSERT INTO conjunto_satisfaccion_mashup "+ 
                        "(mashup_id, avg, usuario_id, fecha) VALUES($1,$2,$3,$4)",
                        conjunto.mashup_id, conjunto.promedio, conjunto.usuario_id, time.Now())
        checkErr(err)
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//66 lineas
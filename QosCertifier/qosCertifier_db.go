package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
    // "strconv"
)

//Declaracion de la cuenta de la base de datos
const (
    DB_USER     = "awaresystems"
    DB_PASSWORD = "3ee798d8"
    DB_NAME     = "awaresystems"
)

//Crear como una Clase ApiDataTest
type ApiDataTest struct {
    id int
    rendimiento float64
    latencia float64
    status float64
    tiempoDeRespuesta float64      
    disponibilidad float64      
    confiabilidad float64
    nombre string
    categoria string
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getNombres(qoss []ApiDataTest) ([]string){
    var nombres []string
    for _, q := range qoss{
        nombres = append(nombres,q.nombre)
    }
    return nombres
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getRendimientos(qoss []ApiDataTest) ([]float64){
    var rendimientos []float64
    for _, q := range qoss{
        if(q.rendimiento == -1){
            rendimientos = append(rendimientos,100000)
        }else{
            rendimientos = append(rendimientos,q.rendimiento)
        }
    }
    return rendimientos
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getLatencias(qoss []ApiDataTest) ([]float64){
    var latencias []float64
    for _, q := range qoss{
        if(q.latencia == -1){
            latencias = append(latencias,100000)
        }else{
            latencias = append(latencias,q.latencia)
        }
    }
    return latencias
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getTiemposDeRespuestas(qoss []ApiDataTest) ([]float64){
    var tiemposDeRespuestas []float64
    for _, q := range qoss{
        if(q.tiempoDeRespuesta == -1){
            tiemposDeRespuestas = append(tiemposDeRespuestas,100000)
        }else{
            tiemposDeRespuestas = append(tiemposDeRespuestas,q.tiempoDeRespuesta)
        }
    }
    return tiemposDeRespuestas
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getDisponibilidades(qoss []ApiDataTest) ([]float64){
    var disponibilidades []float64
    for _, q := range qoss{
        disponibilidades = append(disponibilidades,q.disponibilidad)
    }
    return disponibilidades
}

//Este es un metodo de ApiDataTest
func (this ApiDataTest) getConfiabilidades(qoss []ApiDataTest) ([]float64){
    var confiabilidades []float64
    for _, q := range qoss{
        confiabilidades = append(confiabilidades,q.confiabilidad)
    }
    return confiabilidades
}

//Crea una Clase Estadistica_Percentil
type Estadistica_Percentil struct {
    nivel_factor string
    minimo int
    medio int
    maximo int
    factor int
}


//Se obtiene el ultimo test de cada API
func getData() []ApiDataTest{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT id,rendimiento,latencia,status,tiempo_de_respuesta,disponibilidad"+
                            ",confiabilidad,nombre FROM apis_data_test "+
                            "where to_char(fecha, 'DD-MM-YYYY') = to_char("+
                            "(select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY')"+
                            " from apis_data_test order by fecha desc limit 1),'DD-MM-YYYY') ")
    checkErr(err)

    var qoss []ApiDataTest
    for rows.Next() {
        var idT int
        var rendimientoT float64
        var latenciaT float64
        var statusT float64
        var tiempoDeRespuestaT float64
        var disponibilidadT float64
        var confiabilidadT float64
        var nombreT string
        err = rows.Scan(&idT, &rendimientoT, &latenciaT, &statusT, &tiempoDeRespuestaT, &disponibilidadT, &confiabilidadT,&nombreT)
        checkErr(err)
        temporal := ApiDataTest{
            id: idT,
            rendimiento: rendimientoT,
            latencia: latenciaT,
            status: statusT,
            tiempoDeRespuesta: tiempoDeRespuestaT,
            disponibilidad: disponibilidadT,
            confiabilidad: confiabilidadT,
            nombre: nombreT,
        }
        qoss = append(qoss, temporal)
    }
    return qoss
}

//otro metodo en el aire
func setData( ep Estadistica_Percentil, categoria string) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    fmt.Println(ep)
    var lastInsertId int
    err = db.QueryRow("INSERT INTO estadistica_cuartil(nivel_factor, minimo, medio, maximo,"+
                        " factor_id, categoria, fecha) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;",
                         ep.nivel_factor, ep.minimo, ep.medio, ep.maximo, ep.factor, categoria, time.Now()).Scan(&lastInsertId)
    checkErr(err)
}
/*La categoria de la api es dada por la cantidad mas alta de registros de categoria
que hayan echo los componentes que referencien a la api en particular*/
func getCategoriaDeApi(api_nombre string) (string){
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 := "select count(api.nombre), componentes.categoria from (select * from apis where nombre = '"+api_nombre+
                "') as api"+
                " join componentes on componentes.api_id = api.id group by componentes.categoria"+
                " order by count desc limit 1"
    var categoria string
    rows, err := db.Query(select1)
    checkErr(err)
    for rows.Next() {
        var comodinT int
        var categoriaT string
        err = rows.Scan(&comodinT, &categoriaT)
        categoria = categoriaT
        checkErr(err)
    }
    return categoria
}

// es un try catch 
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//66 lineas
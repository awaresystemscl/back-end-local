package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

//Declaracion de la cuenta de la base de datos
const (
    DB_USER     = "workshop"
    DB_PASSWORD = "workshop2017"
    DB_NAME     = "workshop"
)

//Crear como una Clase Qos
type Qos struct {
    id int
    rendimiento float64
    latencia float64
    status float64
    tiempoDeRespuesta float64      
    disponibilidad float64      
    confiabilidad float64
    nombre string     
}

//Este es un metodo de Qos
func (this Qos) getNombres(qoss []Qos) ([]string){
    var nombres []string
    for _, q := range qoss{
        nombres = append(nombres,q.nombre)
    }
    return nombres
}

//Este es un metodo de Qos
func (this Qos) getRendimientos(qoss []Qos) ([]float64){
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

//Este es un metodo de Qos
func (this Qos) getLatencias(qoss []Qos) ([]float64){
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

//Este es un metodo de Qos
func (this Qos) getTiemposDeRespuestas(qoss []Qos) ([]float64){
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

//Este es un metodo de Qos
func (this Qos) getDisponibilidades(qoss []Qos) ([]float64){
    var disponibilidades []float64
    for _, q := range qoss{
        disponibilidades = append(disponibilidades,q.disponibilidad)
    }
    return disponibilidades
}

//Este es un metodo de Qos
func (this Qos) getConfiabilidades(qoss []Qos) ([]float64){
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


//Es un metodo en el aire
func getData() []Qos{
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT id,rendimiento,latencia,status,tiempo_de_respuesta,disponibilidad"+
                            ",confiabilidad,nombre FROM qos")
    checkErr(err)

    var qoss []Qos
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
        temporal := Qos{
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
func setData( ep Estadistica_Percentil) {
    dbinfo := fmt.Sprintf("host=170.239.84.238 user=%s password=%s dbname=%s sslmode=disable",
        DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    var lastInsertId int
    err = db.QueryRow("INSERT INTO estadistica_percentil(nivel_factor, minimo, medio, maximo,"+
                        " factor, categoria) VALUES($1,$2,$3,$4,$5,'Mapping') returning id;",
                         ep.nivel_factor, ep.minimo, ep.medio, ep.maximo, ep.factor).Scan(&lastInsertId)
}

// es un try catch 
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//66 lineas
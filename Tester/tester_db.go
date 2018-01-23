package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
)

type api struct {
    id int
    nombre string
    descripcion string
    url string      
}

func getData() []api{
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        db_config.Host, db_config.User, db_config.Pass, db_config.Name)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("SELECT id, nombre, descripcion, url FROM apis")
    checkErr(err)

    var apis []api
    for rows.Next() {
        var idT int
        var nombreT string
        var descripcionT string
        var urlT string
        err = rows.Scan(&idT, &nombreT, &descripcionT, &urlT)
        checkErr(err)
        temporal := api{id: idT, nombre: nombreT, descripcion: descripcionT, url: urlT}
        apis = append(apis, temporal)
    }
    return apis
}

func getCount(nombreApi string) (float64, float64, float64){
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        db_config.Host, db_config.User, db_config.Pass, db_config.Name)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    select1 :="(select COALESCE( (select count(nombre) from apis_data_test where apis_data_test.nombre = '"+nombreApi+"' and apis_data_test.status = 200 group by nombre),0))"
    select2 :="(select COALESCE( (select count(nombre) from apis_data_test where apis_data_test.nombre = '"+nombreApi+"' and apis_data_test.status = -1 group by nombre),0))"
    select3 :="select count(nombre) from apis_data_test where apis_data_test.nombre = '"+nombreApi+"' group by nombre"
    var cuentaTotal float64
    var cuentaExito float64
    var cuentaStatusFatal float64
    rows1, err := db.Query(select1)
    rows2, err := db.Query(select2)
    rows3, err := db.Query(select3)
    checkErr(err)
    for rows1.Next() {
        err = rows1.Scan(&cuentaExito)
        checkErr(err)
    }
    for rows2.Next() {
        err = rows2.Scan(&cuentaStatusFatal)
        checkErr(err)
    }
    for rows3.Next() {
        err = rows3.Scan(&cuentaTotal)
        checkErr(err)
    }
    return cuentaExito,cuentaStatusFatal,cuentaTotal
}

func setData(apis_data_testInsert apis_data_test) {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        db_config.Host, db_config.User, db_config.Pass, db_config.Name)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
    var lastInsertId int
    err = db.QueryRow("INSERT INTO apis_data_test (rendimiento, latencia, status, tiempo_de_respuesta,"+
                        " disponibilidad, confiabilidad, fecha, nombre) VALUES($1,$2,$3,$4,$5,$6,$7,$8) returning id",
                         apis_data_testInsert.rendimiento, apis_data_testInsert.latencia, apis_data_testInsert.status,
                         apis_data_testInsert.tiempoDeRespuesta, apis_data_testInsert.disponibilidad,
                         apis_data_testInsert.confiabilidad, time.Now(), apis_data_testInsert.nombre).Scan(&lastInsertId)
    checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
//66 lineas
package main

import (
    "context"
    "fmt"
    "net/http"
    "net/http/httptrace"
    "time"
    "encoding/json"
    "io/ioutil"
)

type configDB struct {
    Host string `json:"HOST_IP"` 
    User string `json:"DB_USER"`
    Pass string `json:"DB_PASSWORD"`
    Name string `json:"DB_NAME"`
}

var db_config configDB

type apis_data_test struct {
    rendimiento int
    latencia int
    status int
    tiempoDeRespuesta int      
    disponibilidad int      
    confiabilidad int
    nombre string     
}
var lat float64

func obtenerMetricas(urlApi,nombreApi string) (apis_data_test){
    metodoHttp := "GET"
    cuentaExito,cuentaStatusFatal,cuentaTotal := getCount(nombreApi) //obtener total de status y tipos
    cuentaTotal += 1
    cuentaStatusFatal += 1
    disponibilidadT := cuentaExito/cuentaTotal*100 // calcular disponibilidad
    confibailidadT := (cuentaTotal-cuentaStatusFatal)/cuentaTotal*100 // calcular confiabilidad
    errorQos := apis_data_test{rendimiento: -1, latencia: -1, status: -1, tiempoDeRespuesta: -1, disponibilidad: int(disponibilidadT), confiabilidad: int(confibailidadT), nombre: nombreApi} //si falla, guardar test en -1
    req, err := http.NewRequest(metodoHttp,urlApi, nil)
    if err != nil {
        return errorQos
    }
    var t1 time.Time
    ctx := context.Background()
    trace := &httptrace.ClientTrace{
        GotFirstResponseByte: func() {
            lat = float64(time.Since(t1).Seconds()*1000)
            },
    }
    req = req.WithContext(httptrace.WithClientTrace(ctx, trace))

    client := new(http.Client)
    t1 = time.Now()
    resp, err := client.Do(req)
    if err != nil {
        return errorQos
    }
    defer resp.Body.Close()
    tiempoDR := time.Since(t1).Seconds()*1000
    fmt.Printf("Latencia: %.0f ms\n", lat)
    fmt.Printf("Tiempo de respuesta: %0.f ms\n", float64(time.Since(t1).Seconds()*1000))
    fmt.Println("Status: ", resp.StatusCode)
    rendimiento,statusRendimiento := testRendimiento(10,1,metodoHttp,urlApi)
    cuentaStatusFatal -= 1
    if resp.StatusCode == 200 {
        cuentaExito += 1
    }
    disponibilidadT = cuentaExito/cuentaTotal*100
    confibailidadT = (cuentaTotal-cuentaStatusFatal)/cuentaTotal*100
    fmt.Println("Rendimiento: ",rendimiento," ms")
    fmt.Println("Porcentaje de exito: ",statusRendimiento*100,"%")
    fmt.Println("Disponibilidad: ",int(disponibilidadT),"%")
    fmt.Println("Confiabilidad: ",int(confibailidadT),"%")
    temporalQoS := apis_data_test{rendimiento: int(rendimiento), latencia: int(lat), status: resp.StatusCode, tiempoDeRespuesta: int(tiempoDR), disponibilidad: int(disponibilidadT), confiabilidad: int(confibailidadT), nombre: nombreApi}
    return temporalQoS
}

func main() {
    db_config = configuracion()
    tiempoDeScript := time.Now()
    dataApis := getData() //obtener apis
    fmt.Println("=======================================================================")
    for _, dataApi := range dataApis{ //recorrer apis
        qosTemp := obtenerMetricas(dataApi.url, dataApi.nombre)
        fmt.Println(qosTemp)
        fmt.Println("=======================================================================")
        setData(qosTemp)
    }
    fmt.Println("El test se ha ejecutado en: ",float64(int(time.Since(tiempoDeScript).Seconds() * 1000)) / 1000)
}

func configuracion() configDB{
    jsonFile, err := ioutil.ReadFile("../config.json")
    checkErr(err)
    var config configDB
    json.Unmarshal(jsonFile, &config)
    return config
}
//58
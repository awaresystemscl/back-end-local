package main

import (
    "context"
    "fmt"
    "net/http"
    "net/http/httptrace"
    "time"
)

type qos struct {
    rendimiento int
    latencia int
    status int
    tiempoDeRespuesta int      
    disponibilidad int      
    confiabilidad int
    nombre string     
}
var lat float64

func obtenerMetricas(urlApi,nombreApi string) (qos){
    metodoHttp := "GET"
    cuentaExito,cuentaStatusFatal,cuentaTotal := getCount(nombreApi)
    cuentaTotal += 1
    cuentaStatusFatal += 1
    disponibilidadT := cuentaExito/cuentaTotal*100
    confibailidadT := (cuentaTotal-cuentaStatusFatal)/cuentaTotal*100
    errorQos := qos{rendimiento: -1, latencia: -1, status: -1, tiempoDeRespuesta: -1, disponibilidad: int(disponibilidadT), confiabilidad: int(confibailidadT), nombre: nombreApi}
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
    temporalQoS := qos{rendimiento: int(rendimiento), latencia: int(lat), status: resp.StatusCode, tiempoDeRespuesta: int(tiempoDR), disponibilidad: int(disponibilidadT), confiabilidad: int(confibailidadT), nombre: nombreApi}
    return temporalQoS
}

func main() {
    tiempoDeScript := time.Now()
    dataApis := getData()
    fmt.Println("=======================================================================")
    for _, dataApi := range dataApis{
        qosTemp := obtenerMetricas(dataApi.url, dataApi.nombre)
        fmt.Println()
        fmt.Println("=======================================================================")
        setData(qosTemp)
    }
    fmt.Println("El test se ha ejecutado en: ",float64(int(time.Since(tiempoDeScript).Seconds() * 1000)) / 1000)
}
//58
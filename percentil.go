package main

import (
    "fmt"
    "sort"
)

type Quartile struct {

}

//Estructura de los metodos:
// (clase Nombre_de_la_Clase) metodo(Parametros_de_entrada) (Parametros_de_salida)opcional*
func (this Quartile) quartile(valores []float64, porciento float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    rango := int64((float64(len(copiaQ)) * porciento/100)+0.5)
    return copiaQ[rango]
}

func (this Quartile) minimo(valores []float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    return copiaQ[0]
}

func (this Quartile) maximo(valores []float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    return copiaQ[len(valores)-1]
}


func main() {

    qos := new(Qos)
    pruebas := getData()
    rendimientos := qos.getRendimientos(pruebas)
    latencias := qos.getLatencias(pruebas)
    tiemposDeRespuestas := qos.getTiemposDeRespuestas(pruebas)
    disponibilidades := qos.getDisponibilidades(pruebas)
    confiabilidades := qos.getConfiabilidades(pruebas)
    quartil := new(Quartile)
    // arreglo := []float64{5,6,4,8,7,1,3,9,2}
    //Rendimiento
    minimo := quartil.minimo(rendimientos)
    aceptable := quartil.quartile(rendimientos,25)
    normal := quartil.quartile(rendimientos,50)
    pocoAceptable := quartil.quartile(rendimientos,75)
    maximo := quartil.maximo(rendimientos)
    fmt.Println("========================================================================================")
    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
    //Latencia
    minimo = quartil.minimo(latencias)
    aceptable = quartil.quartile(latencias,25)
    normal = quartil.quartile(latencias,50)
    pocoAceptable = quartil.quartile(latencias,75)
    maximo = quartil.maximo(latencias)
    fmt.Println("========================================================================================")
    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
    //Tiempo de respuesta
    minimo = quartil.minimo(tiemposDeRespuestas)
    aceptable = quartil.quartile(tiemposDeRespuestas,25)
    normal = quartil.quartile(tiemposDeRespuestas,50)
    pocoAceptable = quartil.quartile(tiemposDeRespuestas,75)
    maximo = quartil.maximo(tiemposDeRespuestas)
    fmt.Println("========================================================================================")
    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
    //Disponibilidad
    minimo = quartil.minimo(disponibilidades)
    aceptable = quartil.quartile(disponibilidades,25)
    normal = quartil.quartile(disponibilidades,50)
    pocoAceptable = quartil.quartile(disponibilidades,75)
    maximo = quartil.maximo(disponibilidades)
    fmt.Println("========================================================================================")
    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
    //Confiabilidad
    minimo = quartil.minimo(confiabilidades)
    aceptable = quartil.quartile(confiabilidades,25)
    normal = quartil.quartile(confiabilidades,50)
    pocoAceptable = quartil.quartile(confiabilidades,75)
    maximo = quartil.maximo(confiabilidades)
    fmt.Println("========================================================================================")
    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
    fmt.Println("========================================================================================")

}
//58
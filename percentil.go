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
    quartil := new(Quartile)
    arreglo := []float64{5,6,4,8,7,1,3,9,2}

    minimo := quartil.minimo(arreglo)
    aceptable := quartil.quartile(arreglo,25)
    normal := quartil.quartile(arreglo,50)
    pocoAceptable := quartil.quartile(arreglo,75)
    maximo := quartil.maximo(arreglo)

    fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
}
//58
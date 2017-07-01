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

func estadisticaPercentil(arregloDeFactores []float64, nombreFactor string){
    quartil := new(Quartile)
    q0 := int(quartil.minimo(arregloDeFactores))
    q25 := int(quartil.quartile(arregloDeFactores,25))
    q50 := int(quartil.quartile(arregloDeFactores,50))
    q75 := int(quartil.quartile(arregloDeFactores,75))
    q100 := int(quartil.maximo(arregloDeFactores))
    mAceptable := Estadistica_Percentil{
        nivel_factor: "Muy Aceptable",
        minimo: q0,
        medio: q0,
        maximo: q25,
        factor: nombreFactor,
    }
    aceptable := Estadistica_Percentil{
        nivel_factor: "Aceptable",
        minimo: q0,
        medio: q25,
        maximo: q50,
        factor: nombreFactor,
    }
    normal := Estadistica_Percentil{
        nivel_factor: "Normal",
        minimo: q25,
        medio: q50,
        maximo: q75,
        factor: nombreFactor,
    }
    pAcceptable := Estadistica_Percentil{
        nivel_factor: "Poco Aceptable",
        minimo: q50,
        medio: q75,
        maximo: q100,
        factor: nombreFactor,
    }
    mpAceptable := Estadistica_Percentil{
        nivel_factor: "Muy poco Aceptable",
        minimo: q75,
        medio: q100,
        maximo: q100,
        factor: nombreFactor,
    }
    setData(mAceptable)
    setData(aceptable)
    setData(normal)
    setData(pAcceptable)
    setData(mpAceptable)
    fmt.Println("========================================================================================")
    // fmt.Println("Muy Aceptable - minimo:",minimo,"medio:",minimo,"maximo:",aceptable)
    // fmt.Println("Aceptable - minimo:",minimo,"medio:",aceptable,"maximo:",normal)
    // fmt.Println("Normal - minimo:",aceptable,"medio:",normal,"maximo:",pocoAceptable)
    // fmt.Println("Poco aceptable - minimo:",normal,"medio:",pocoAceptable,"maximo:",maximo)
    // fmt.Println("Muy Poco Aceptable - minimo:",pocoAceptable,"medio:",maximo,"maximo:",maximo)
}

func main() {

    qos := new(Qos)
    pruebas := getData()
    rendimientos := qos.getRendimientos(pruebas)
    latencias := qos.getLatencias(pruebas)
    tiemposDeRespuestas := qos.getTiemposDeRespuestas(pruebas)
    disponibilidades := qos.getDisponibilidades(pruebas)
    confiabilidades := qos.getConfiabilidades(pruebas)
    estadisticaPercentil(rendimientos,"Rendimiento")
    estadisticaPercentil(latencias,"Latencia")
    estadisticaPercentil(tiemposDeRespuestas,"Tiempo de Respuesta")
    estadisticaPercentil(disponibilidades,"Disponibilidad")
    estadisticaPercentil(confiabilidades,"Confiabilidad")
    // arreglo := []float64{5,6,4,8,7,1,3,9,2}
}
//58
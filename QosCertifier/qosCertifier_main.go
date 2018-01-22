package main

import (
    "fmt"
    "sort"
)

//Esto declara como la clase Quartil
type Quartile struct {

}

//Estructura de los metodos:
// (clase Nombre_de_la_Clase) metodo(Parametros_de_entrada) (Parametros_de_salida)opcional*
//Este es un metodo de Quartile
func (this Quartile) quartile(valores []float64, percentil float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    var rango int64
    rango = int64((float64(len(copiaQ)) * percentil/100)-0.5) 
    // fmt.Println(rango)
    // if rango >= int64(len(copiaQ)){
    //     rango = rango-1
    // }
    if rango < 0{
        rango = 0
    }
    return copiaQ[rango]
}

//Este es un metodo de Quartile
func (this Quartile) minimo(valores []float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    return copiaQ[0]
}

//Este es un metodo de Quartile
func (this Quartile) maximo(valores []float64) (float64){
    copiaQ := make([]float64, len(valores))
    copy(copiaQ,valores)
    sort.Float64s(copiaQ)
    return copiaQ[len(valores)-1]
}

type GrupoDeApis struct{
    grupo []ApiDataTest
    nombre string
}

// Este es un metodo que ordena y obtiene los percentiles para ingresar un factor para una categoria
func estadisticaPercentil(arregloDeFactores []float64, nombreFactor int, categoria string){
    quartil := new(Quartile)
    fmt.Println("================= Categoria ===================")
    fmt.Println(categoria)
    fmt.Println(arregloDeFactores)
    q0 := int(quartil.minimo(arregloDeFactores))
    q25 := int(quartil.quartile(arregloDeFactores,25))
    q50 := int(quartil.quartile(arregloDeFactores,50))
    q75 := int(quartil.quartile(arregloDeFactores,75))
    q100 := int(quartil.maximo(arregloDeFactores))
    var mAceptable Estadistica_Percentil
    var aceptable Estadistica_Percentil
    var normal Estadistica_Percentil
    var pAcceptable Estadistica_Percentil
    var mpAceptable Estadistica_Percentil

    // 2(Disponibilidad) y 4(Confiabilidad)
    // son los factores que van de mayor a menor en nivel de calidad
    if nombreFactor == 2 || nombreFactor == 4{ 
        mAceptable = Estadistica_Percentil{
            nivel_factor: "Muy poco Aceptable",
            minimo: q0,
            medio: q0,
            maximo: q25,
            factor: nombreFactor,
        }
        aceptable = Estadistica_Percentil{
            nivel_factor: "Poco Aceptable",
            minimo: q0,
            medio: q25,
            maximo: q50,
            factor: nombreFactor,
        }
        normal = Estadistica_Percentil{
            nivel_factor: "Normal",
            minimo: q25,
            medio: q50,
            maximo: q75,
            factor: nombreFactor,
        }
        pAcceptable = Estadistica_Percentil{
            nivel_factor: "Aceptable",
            minimo: q50,
            medio: q75,
            maximo: q100,
            factor: nombreFactor,
        }
        mpAceptable = Estadistica_Percentil{
            nivel_factor: "Muy Aceptable",
            minimo: q75,
            medio: q100,
            maximo: q100,
            factor: nombreFactor,
        }
    // en caso contrario van de menor a mayor
    }else{
        mAceptable = Estadistica_Percentil{
            nivel_factor: "Muy Aceptable",
            minimo: q0,
            medio: q0,
            maximo: q25,
            factor: nombreFactor,
        }
        aceptable = Estadistica_Percentil{
            nivel_factor: "Aceptable",
            minimo: q0,
            medio: q25,
            maximo: q50,
            factor: nombreFactor,
        }
        normal = Estadistica_Percentil{
            nivel_factor: "Normal",
            minimo: q25,
            medio: q50,
            maximo: q75,
            factor: nombreFactor,
        }
        pAcceptable = Estadistica_Percentil{
            nivel_factor: "Poco Aceptable",
            minimo: q50,
            medio: q75,
            maximo: q100,
            factor: nombreFactor,
        }
        mpAceptable = Estadistica_Percentil{
            nivel_factor: "Muy poco Aceptable",
            minimo: q75,
            medio: q100,
            maximo: q100,
            factor: nombreFactor,
        }
        
    }
    setData(mAceptable, categoria)
    setData(aceptable, categoria)
    setData(normal, categoria)
    setData(pAcceptable, categoria)
    setData(mpAceptable, categoria)
    fmt.Println("============================================================")
    // fmt.Println(mAceptable)
    // fmt.Println(aceptable)
    // fmt.Println(normal)
    // fmt.Println(pAcceptable)
    // fmt.Println(mpAceptable)
}

func agrupar(apis []ApiDataTest)([]GrupoDeApis){
    var categorias []string
    var grupoApisCategoria []GrupoDeApis
    for index, apiTest := range apis{
        apis[index].categoria = getCategoriaDeApi(apiTest.nombre)
    }
    //obtengo un arreglo de todas las categorias existentes
    for _, apiTest := range apis{
        validador := true
        for _, c := range categorias{
            if c == apiTest.categoria{
                validador = false
            }
        }
        if validador{
            categorias = append(categorias, apiTest.categoria)
        }
    }
    // fmt.Println(categorias)

    // Se asosian las apis a su respectivo conjunto de categoria
    for _, c := range categorias{
        var temporal GrupoDeApis
        temporal.nombre = c
        for _, apiTest := range apis{
            if apiTest.categoria == c{
                temporal.grupo = append(temporal.grupo, apiTest)
            }
        }
        grupoApisCategoria = append(grupoApisCategoria, temporal)
    }

    // for _, g := range grupoApisCategoria{
    //     fmt.Println(g)
    // }
    return grupoApisCategoria

}

//Agrega un grupo de categoria al algoritmo de percentil
func agregarGrupo(g GrupoDeApis){
    apiDataTest := new(ApiDataTest)
    rendimientos := apiDataTest.getRendimientos(g.grupo)
    latencias := apiDataTest.getLatencias(g.grupo)
    tiemposDeRespuestas := apiDataTest.getTiemposDeRespuestas(g.grupo)
    disponibilidades := apiDataTest.getDisponibilidades(g.grupo)
    confiabilidades := apiDataTest.getConfiabilidades(g.grupo)
    // fmt.Println(g.grupo)
    estadisticaPercentil(rendimientos,3, g.nombre)
    estadisticaPercentil(latencias,5, g.nombre)
    estadisticaPercentil(tiemposDeRespuestas,1, g.nombre)
    estadisticaPercentil(disponibilidades,2, g.nombre)
    estadisticaPercentil(confiabilidades,4, g.nombre)
}

func main() {

    // apiDataTest := new(ApiDataTest)
    pruebas := getData()
    grupos := agrupar(pruebas)
    for _, g := range grupos{
        agregarGrupo(g)
    }




    
    // arreglo := []float64{5,6,4,8,7,1,3,9,2}
}
//58
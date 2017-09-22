package main

import (
    "fmt"
)
func tipoFactor(factor int, test apiTest) (int){
    var valor int
    switch factor{
    case 1:
        valor = test.tiempo_de_respuesta
    case 2:
        valor =  test.disponibilidad
    case 3:
        valor =  test.rendimiento
    case 4:
        valor =  test.confiabilidad
    case 5:
        valor =  test.latencia
    }
    return valor
}

func monitorearComponente(componenteId int, mashupId int){
    componente := componenteId
    mashup := mashupId
    temporal1 := getApiTest(componente)
    temporal2 := getComponentRules(componente)
    temporal3 := getCatCuartil(componente)

    for _, rule := range temporal2{ //requerimiento
        for _, cuartil := range temporal3{ //cuartil
            if cuartil.factor == rule.nombre && cuartil.nivel_factor == rule.nivel{ //si el cuartil es = al requerimiento
                valor := tipoFactor(rule.nombre,temporal1) // se obtiene el valor del requerimiento
                if rule.tendencia == true { // si la regla dice que es ** A LO MAS **
                    if rule.tipoDeMedida == "menor"{ // si la calidad aumenta a MENOR numero
                        if valor >= cuartil.medio {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:100, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("100 "+strconv.Itoa(rule.nombre))
                        }else if valor <= cuartil.minimo {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:0, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("0 "+strconv.Itoa(rule.nombre))
                        }else{
                            temporal:= int((valor-cuartil.minimo)*100/(cuartil.medio-cuartil.minimo))
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:temporal, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println(strconv.Itoa(temporal)+strconv.Itoa(rule.nombre))
                        }
                    }else{ //si la calidad aumenta a MAYOR numero
                        if valor <= cuartil.medio {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:100, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("100 "+strconv.Itoa(rule.nombre))
                        }else if valor >= cuartil.maximo {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:0, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("0 "+strconv.Itoa(rule.nombre))
                        }else{
                            temporal:= int((cuartil.maximo-valor)*100/(cuartil.maximo-cuartil.medio))
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:temporal, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println(strconv.Itoa(temporal)+strconv.Itoa(rule.nombre))
                        }
                    }
                }else{ // en caso de que sea ** A LO MENOS ** que seria decir false
                    if rule.tipoDeMedida == "menor"{ // si la calidad aumenta a MENOR numero
                        if valor <= cuartil.medio {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:100, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("100 "+strconv.Itoa(rule.nombre))
                        }else if valor >= cuartil.maximo {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:0, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("0 "+strconv.Itoa(rule.nombre))
                        }else{
                            temporal:= int((cuartil.maximo-valor)*100/(cuartil.maximo-cuartil.medio))
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:temporal, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println(strconv.Itoa(temporal)+strconv.Itoa(rule.nombre))
                        }
                    }else{ //si la calidad aumenta a MAYOR numero
                        if valor >= cuartil.medio {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:100, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("100 "+strconv.Itoa(rule.nombre))
                        }else if valor <= cuartil.minimo {
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:0, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println("0 "+strconv.Itoa(rule.nombre))
                        }else{
                            temporal:= int((valor-cuartil.medio)*100/(cuartil.medio-cuartil.minimo))
                            insert := satisfaccion_componente{componente_id:componente, satisfaccion:temporal, factor:rule.nombre, mashup_id:mashup }
                            setSatisfaccionCompo(insert)
                            // fmt.Println(strconv.Itoa(temporal)+strconv.Itoa(rule.nombre))
                        }
                    }
                }
            }
        }        
    }
}

func main() {
    componentes := getComponentes()
    // fmt.Println(componentes)
    // fmt.Println("------------")
    for _, componente := range componentes{
        monitorearComponente(componente.id, componente.mashup_id)
    }
    conjunto_compo := getSatisfaccion()
    setConjuntoCompo(conjunto_compo)
    // fmt.Println(conjunto_compo)
    conjunto_mashup := getConjuntoCompo()
    setConjuntoMashup(conjunto_mashup)
    // fmt.Println(conjunto_mashup)
}
//58

//select to_date(to_char(fecha,'DD-MM-YYYY'),'DD-MM-YYYY') from satisfaccion_componente order by fecha asc limit 1;
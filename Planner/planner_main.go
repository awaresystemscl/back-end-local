package main

import (
    "strconv"
    "fmt"
    "encoding/json"
    "io/ioutil"
)
type mashup struct {
    id int
    limite int
}

type configDB struct {
    Host string `json:"HOST_IP"` 
    User string `json:"DB_USER"`
    Pass string `json:"DB_PASSWORD"`
    Name string `json:"DB_NAME"`
    Tolerancia int `json:"Tolerancia"`
    Ventana int `json:"Ventana"`
}

var db_config configDB

func masMenos(restriccion bool)string{
    if restriccion {
        return "a lo mas"
    }else{
        return "a lo menos"
    }
}

func main() {
    db_config = configuracion()
	alertas := getAlertas()
    for i, alerta := range alertas{
        fmt.Println(alerta)
        componentes := getComponentesMashup(alerta.mashup_id)
        alertas[i].componentes = componentes
        for j, c := range componentes{
            restricciones := getRestricciones(c.id)
            alertas[i].componentes[j].restricciones = restricciones
        }
    }
    for _, alerta := range alertas{
        var mensaje string
        mensaje = "Estimad@ "+alerta.nombre+"\n\nLe informamos que la aplicacion "+alerta.mashup_nombre+
                    " el cual contiene los siguientes componentes: \n"
        for _, c := range alerta.componentes{
            mensaje = mensaje+"- "+c.nombre+"("+c.categoria+")\n"
        }
        mensaje = mensaje+"Tienen una satisfaccion inferior al rango que usted considera como correcto ("+
                    strconv.Itoa(alerta.mashup_limite)+"%) presentando una satisfaccion de "+
                    strconv.Itoa(alerta.promedio)+"%. Donde:\n\n"
        for _, c := range alerta.componentes{
            mensaje = mensaje+"- El componente "+c.nombre+", satisface en un "
            for i, r := range c.restricciones{
                mensaje = mensaje+strconv.Itoa(r.satisfaccion)+"% el requerimiento no funcional de "+r.factor+" "+
                            masMenos(r.tendencia)+" "+r.nivel
                if i+1 < len(c.restricciones){
                    mensaje = mensaje+", "
                }else{
                    mensaje = mensaje+".\n\n"
                }
            }
        }
    mensaje = mensaje+"\n Equipo AwareSystems"
    send(mensaje,alerta.nombre,alerta.email)
    fmt.Println(mensaje)
    }
}

func configuracion() configDB{
    jsonFile, err := ioutil.ReadFile("../config.json")
    checkErr(err)
    var config configDB
    json.Unmarshal(jsonFile, &config)
    return config
}
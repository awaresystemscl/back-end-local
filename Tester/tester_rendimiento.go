package main

import (
  // "fmt"
  "time"
  vegeta "github.com/tsenart/vegeta/lib"
)

func testRendimiento(cantidadDeTestPorSegundo, cantidadDeVecesPorSegundo int, metodoHttp, urlApi string ) (float64, float64){
  rate := uint64(cantidadDeTestPorSegundo) // cantidad por segundo
  duration := time.Duration(cantidadDeVecesPorSegundo) * time.Second //cantidad de segundos que se repetira
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: metodoHttp, //metodo http
    URL:    urlApi, //url
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()
  return float64(metrics.Latencies.P99/time.Millisecond), float64(metrics.Success)
}
//12
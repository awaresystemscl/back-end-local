package main

import (
  "fmt"
  "time"
  "os"

  vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
  rate := uint64(20) // cantidad por segundo
  duration := 1 * time.Second //cantidad de segundos que se repetira
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: "GET", //metodo http
    URL:    "http://api.openweathermap.org/data/2.5/weather?lat=-33.019&lon=-71.49&appid=5d9c00182991935598975f1780bbc808", //url
  })
  attacker := vegeta.NewAttacker()

  // open the out file for writing
  outfile, err := os.Create("./out.json") //se crea un archivo
  if err != nil {
      panic(err)
  }
  defer outfile.Close()
  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  reporter := vegeta.NewJSONReporter(&metrics)
  metrics.Close()
  reporter.Report(outfile)

  fmt.Printf("99th percentile: %v\n", float64(metrics.Latencies.P99/time.Millisecond))
  fmt.Printf("99th percentile: %v\n", metrics.Latencies.P99)
  fmt.Printf("Exitosas: %v\n", metrics.Success)
}
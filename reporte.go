package main

import (
  vegeta "github.com/tsenart/vegeta/lib"
  "time"
  "os"
)

func main() {
  targets := vegeta.NewTargets([]string{"GET http://localhost:9100/"})
  rate := uint64(100) // per second
  duration := 4 * time.Second
  reporter := vegeta.NewTextReporter()

  vegeta.Attack(targets, rate, duration, reporter)

  reporter.Report(os.Stdout)
}

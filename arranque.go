package main

import (
	"log"
	"os/exec"
	"fmt"
)

func main() {
	tester, err := exec.Command("./Tester/tester").Output()
	fmt.Printf("%s\n", tester)
	certifer, err := exec.Command("./QosCertifier/certifer").Output()
	fmt.Printf("%s\n", certifer)
	monitoring, err := exec.Command("./Monitoring/monitoring").Output()
	fmt.Printf("%s\n", monitoring)
	analyzing, err := exec.Command("./Analyzing/analyzing").Output()
	fmt.Printf("%s\n", analyzing)
	planner, err := exec.Command("./Planner/planner").Output()
	fmt.Printf("%s\n", planner)
	if err != nil {
		log.Fatal(err)
	}
}
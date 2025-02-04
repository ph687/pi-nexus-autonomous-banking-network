package main

import (
	"fmt"
	"log"

	"nexarion/internal/nci"
	"nexarion/internal/qc"
	"nexarion/internal/ida"
	"nexarion/internal/nc"
)

func main() {
	// Initialize components
	nci.Init()
	qc.Init()
	ida.Init()
	nc.Init()

	// Create Nexarion Core instance
	core := nc.NewNexarionCore()

	// Perform a simple task
	input := nci.GetInput()
	result := qc.PerformQuantumComputation(input)
	output := ida.VisualizeResult(result)
	nci.DisplayOutput(output)

	fmt.Println("Simplified Nexarion is online!")
}

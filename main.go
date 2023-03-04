package main

import (
	"log"
	"time"

	core "github.com/rromulos/aps-finder/internal"
	"github.com/rromulos/aps-finder/pkg/logger"
	"github.com/rromulos/aps-finder/pkg/report"
)

func main() {
	logger.InitLogs()
	report.InitReports()
	var verboseMode = ""

	for {
		verboseMode = core.EnableVerboseMode()
		if verboseMode == "y" || verboseMode == "n" {
			break
		}
	}

	start := time.Now()
	core.PerformAnalysis("app", ".php", verboseMode)
	elapsed := time.Since(start)
	println("=====================================================================")
	log.Printf("Execution took %s", elapsed)
}

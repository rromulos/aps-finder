package main

import (
	"log"
	"time"

	"github.com/rromulos/aps-finder/core"
	"github.com/rromulos/aps-finder/helpers/logger"
	"github.com/rromulos/aps-finder/helpers/report"
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

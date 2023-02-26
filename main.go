package main

import (
	"log"
	"time"

	"github.com/rromulos/aps-finder/helpers/core"
	"github.com/rromulos/aps-finder/helpers/logger"
	"github.com/rromulos/aps-finder/helpers/report"
)

func main() {
	logger.InitLogs()
	report.InitReports()
	prefix := ""
	for {
		prefix = core.EnableDebugMode()
		if prefix == "y" || prefix == "n" {
			break
		}
	}
	start := time.Now()
	core.PerformAnalysis("app", ".php", prefix)
	elapsed := time.Since(start)
	println("=====================================================================")
	log.Printf("Execution took %s", elapsed)
}

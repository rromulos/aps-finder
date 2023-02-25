package main

import (
	"log"
	"time"

	"github.com/rromulos/aps-finder/helpers/core"
	"github.com/rromulos/aps-finder/helpers/logger"
)

func main() {
	logger.InitLogs()
	var prefix = core.EnableDebugMode()
	start := time.Now()
	core.PerformAnalysis("app", ".php", prefix)
	elapsed := time.Since(start)
	println("=====================================================================")
	log.Printf("Execution took %s", elapsed)
}

package main

import (
	core "github.com/rromulos/aps-finder/internal"
	"github.com/rromulos/aps-finder/pkg/logger"
	"github.com/rromulos/aps-finder/pkg/report"
)

func main() {
	logger.InitLogs()
	report.InitReports()

	core.ShowMainMenu()
}

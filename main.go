package main

import (
	"github.com/rromulos/aps-finder/internal/menus"
	"github.com/rromulos/aps-finder/pkg/logger"
	"github.com/rromulos/aps-finder/pkg/report"
)

func main() {
	logger.InitLogs()
	report.InitReports()

	menus.ShowMainMenu()
}

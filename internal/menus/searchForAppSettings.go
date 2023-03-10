package menus

import (
	"bufio"
	"fmt"
	"os"

	core "github.com/rromulos/aps-finder/internal"
)

//Starts the process to search for app settings.
func MenuSearchForAppSettings() {
	var verboseMode string

	for {
		fmt.Print("Do you want to enable the verbose mode? (y/n) ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		verboseMode = input.Text()

		if verboseMode == "y" || verboseMode == "n" {
			break
		}
	}

	core.PerformAnalysis(verboseMode)
}

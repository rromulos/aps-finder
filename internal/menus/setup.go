package menus

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rromulos/aps-finder/internal/messages"
	dotEnvHelper "github.com/rromulos/aps-finder/pkg/dotenvhelper"
	logger "github.com/rromulos/aps-finder/pkg/logger"
)

//Show the setup menu
//@TODO refactoring would be welcome for this function
func MenuSetup() {
	for {
		displayMenu := checkIfTheSetupMenuCanBeDisplayed()

		if displayMenu {
			fmt.Print(messages.TYPE_APPLICATION_PATH)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			checkPathExists, err := dotEnvHelper.CheckIfPathExists(input.Text())

			if err != nil {
				logger.Log(logger.WARN, messages.ERROR_DURING_PATH_CHECK, logger.APP_FINDER_LOG)
				panic(err)
			}

			if !checkPathExists {
				logger.Log(logger.WARN, messages.APP_PATH_DOES_NOT_EXISTS, logger.APP_FINDER_LOG)
				fmt.Println(messages.APP_PATH_DOES_NOT_EXISTS)
				continue
			}

			if os.Getenv("APP_PATH") == "" {
				dotEnvHelper.SetPath("APP_PATH=" + input.Text())
				ShowMainMenu()
				break
			}
		}
		ShowMainMenu()
		break
	}
}

//Checks if the setup menu can be displayed.
//Return true if so.
//Return false if not.
func checkIfTheSetupMenuCanBeDisplayed() bool {
	var dotEnvCheck, _ = dotEnvHelper.CheckIfDotEnvFileExists()

	if !dotEnvCheck {
		dotEnvHelper.CreateEnv()
	}

	dotEnvCheck = dotEnvHelper.CheckIfDotEnvContentIsValid()

	return dotEnvCheck
}

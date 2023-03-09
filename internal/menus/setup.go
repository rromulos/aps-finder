package menus

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	core "github.com/rromulos/aps-finder/internal"
	dotEnvHelper "github.com/rromulos/aps-finder/pkg/dotenvhelper"
	logger "github.com/rromulos/aps-finder/pkg/logger"
)

//Show the setup menu
//@TODO refactoring would be welcome for this function
func MenuSetup() {
	for {
		err := godotenv.Load()

		if err != nil {
			logger.Log(logger.WARN, core.ERROR_LOADING_DOTENV, logger.APP_FINDER_LOG)
			dotEnvHelper.CreateEnv()
		}

		appPath := dotEnvHelper.GetPath()

		if appPath != "" {
			ShowMainMenu()
			break
		}

		fmt.Print("Type the Path of the RepairQ application (app folder): ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		checkPathExists, err := dotEnvHelper.CheckIfPathExists(input.Text())

		if err != nil {
			logger.Log(logger.WARN, core.ERROR_DURING_PATH_CHECK, logger.APP_FINDER_LOG)
			panic(err)
		}

		if !checkPathExists {
			logger.Log(logger.WARN, core.APP_PATH_DOES_NOT_EXISTS, logger.APP_FINDER_LOG)
			fmt.Println(core.APP_PATH_DOES_NOT_EXISTS)
			continue
		}

		if appPath == "" {
			dotEnvHelper.SetPath("APP_PATH=" + input.Text())
			ShowMainMenu()
			break
		}
	}
}

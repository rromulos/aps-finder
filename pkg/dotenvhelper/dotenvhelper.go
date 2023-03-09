package dotEnvHelper

import (
	"os"

	"github.com/joho/godotenv"
	core "github.com/rromulos/aps-finder/internal"
	"github.com/rromulos/aps-finder/pkg/logger"
)

//checks if the informed path exists
//returns whether the given file or directory exists
func CheckIfPathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

//creates a .env file
func CreateEnv() {
	file, err := os.Create(".env")

	if err != nil {
		panic(err)
	}

	defer file.Close()
}

//write the application path into .env file
func SetPath(content string) {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logger.Log(logger.ERROR, core.ERROR_LOADING_DOTENV, logger.APP_FINDER_LOG)
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		logger.Log(logger.ERROR, core.CANT_WRITE_IN_DOTENV, logger.APP_FINDER_LOG)
		panic(err)
	}
}

//Get the APP_PATH in the .env file
func GetPath() string {
	godotenv.Load()
	return os.Getenv("APP_PATH")
}

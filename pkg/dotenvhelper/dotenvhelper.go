package dotEnvHelper

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rromulos/aps-finder/internal/messages"
	"github.com/rromulos/aps-finder/pkg/logger"
)

//Checks if the path typed by the user exists.
//Returns true and nil if the path does exists.
//Returns false and nil if the path does not exists.
//Returns false and error if something went wrong when checking the path.
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

//Create an empty dotenv file.
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
		logger.Log(logger.ERROR, messages.ERROR_LOADING_DOTENV, logger.APP_FINDER_LOG)
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		logger.Log(logger.ERROR, messages.CANT_WRITE_IN_DOTENV, logger.APP_FINDER_LOG)
		panic(err)
	}
}

//Checks if dotenv file exists.
//Returns false and the error if it does not exist.
//Returns true and nil if the dotenv does exists.
func CheckIfDotEnvFileExists() (bool, error) {
	err := godotenv.Load()

	if err != nil {
		logger.Log(logger.WARN, messages.ERROR_LOADING_DOTENV, logger.APP_FINDER_LOG)
		return false, err
	}

	return true, nil
}

//Checks if the dotenv file contains all the variables needed to run the application.
//Returns true if all variables are found.
//Returns false if some variable is missing.
func CheckIfDotEnvContentIsValid() bool {
	dotEnvCheck, _ := CheckIfDotEnvFileExists()

	if !dotEnvCheck {
		return false
	}

	godotenv.Load()

	appPath := os.Getenv("APP_PATH")

	if appPath != "" {
		return false
	}

	return true
}

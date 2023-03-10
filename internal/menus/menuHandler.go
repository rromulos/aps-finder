package menus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	dotEnvHelper "github.com/rromulos/aps-finder/pkg/dotenvhelper"
	"github.com/rromulos/aps-finder/pkg/terminal"
)

//Display the main menu.
//@TODO the contents of this method needs to be moved to a specific file for the main menu inside the menus folder.
func ShowMainMenu() {

	var checkDotEnvFileExists = false

	checkDotEnvFileExists, _ = dotEnvHelper.CheckIfDotEnvFileExists()
	var checkDotEnvContentIsValid = dotEnvHelper.CheckIfDotEnvContentIsValid()

	if !checkDotEnvFileExists || !checkDotEnvContentIsValid {
		terminal.CleanTerminal()
		fmt.Println("1 - Setup")
	} else {
		terminal.CleanTerminal()
		fmt.Println("2 - Search for App Settings")
		fmt.Println("3 - Compare App Settings")
		fmt.Println("4 - Start Web Service")
	}
	fmt.Println("5 - Close Application")
	fmt.Println("----------------------------")
	fmt.Print("Choose your option :")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	for !validateChosenMenuIsValid(input.Text()) {
		terminal.CleanTerminal()
		break
	}

	redirectToChosenMenu(input.Text())
}

//Checks if the option chosen in the main menu is valid.
//Returns true if so.
//Returns false if not.
func validateChosenMenuIsValid(chosenOption string) bool {
	return contains(chosenOption)
}

//Checks if the chosen options is a valid option, considering the array values.
//Returns true if so.
//Returns false if not.
func contains(str string) bool {
	validMenuOptions := []string{"1", "2", "3", "4", "5"}
	for _, s := range validMenuOptions {
		if s == str {
			return true
		}
	}
	return false
}

//Redirects to the menu referring to the option chosen in the main menu.
//@TODO adds log file to this method (error cases).
//@TODO method is still incomplete.
func redirectToChosenMenu(chosenOption string) {
	num, err := strconv.Atoi(chosenOption)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	switch num {
	case 1:
		MenuSetup()
	case 2:
		MenuSearchForAppSettings()
	case 3:
		fmt.Println("Three")
	default:
		ShowMainMenu()
	}
}

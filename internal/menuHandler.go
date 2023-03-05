package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/rromulos/aps-finder/pkg/terminal"
)

//shows the main menu
func ShowMainMenu() {
	fmt.Println("1 - Setup")
	fmt.Println("2 - Search for App Settings")
	fmt.Println("3 - Compare App Settings")
	fmt.Println("4 - Start Web Service")
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

//checks if the option chosen in the main menu is valid
func validateChosenMenuIsValid(chosenOption string) bool {
	return contains(chosenOption)
}

//checks if the chosen options is a valid option
func contains(str string) bool {
	validMenuOptions := []string{"1", "2", "3", "4", "5"}
	for _, s := range validMenuOptions {
		if s == str {
			return true
		}
	}
	return false
}

//redirects to the chosen menu
func redirectToChosenMenu(chosenOption string) {
	num, err := strconv.Atoi(chosenOption)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		menuSearchForAppSettings()
	case 3:
		fmt.Println("Three")
	default:
		ShowMainMenu()
	}
}

//Starts the process to search for app settings
func menuSearchForAppSettings() {
	for {
		fmt.Print("Do you want to enable the verbose mode? (y/n) ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		verboseMode = input.Text()
		if verboseMode == "y" || verboseMode == "n" {
			break
		}
	}
	PerformAnalysis(verboseMode)
}

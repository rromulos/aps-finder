package core

import (
	"bufio"
	"fmt"
	"os"
)

func EnableVerboseMode() string {
	fmt.Print("Do you want to enable the verbose mode? (y/n) ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

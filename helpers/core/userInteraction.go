package core

import (
	"bufio"
	"fmt"
	"os"
)

func EnableDebugMode() string {
	fmt.Print("Do you want to enable the debug mode? (y/n) ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

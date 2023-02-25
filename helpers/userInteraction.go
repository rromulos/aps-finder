package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func WhatProject() string {
	fmt.Print("Enter the branch prefix (CPR/MK/OTS): ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

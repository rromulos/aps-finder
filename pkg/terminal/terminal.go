package terminal

import (
	"os"
	"os/exec"
	"runtime"
)

//Invokes the terminal cleanup method (runCmd), passing the command according to the operating system.
func CleanTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

//Execute terminal cleanup command.
func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

package terminal

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/everdev/mack"
)

//Notify terminal notifier
func Notify(data []byte) (err error) {
	sys := runtime.GOOS
	d := string(data)
	switch {
	case sys == "darwin":
		err := mack.Notify(fmt.Sprintf("Occurence found! %v", d), "mouscat notifier")
		if err != nil {
			return err
		}
	default:
		exec.Command("tput", "bel").Run()
	}
	return nil
}

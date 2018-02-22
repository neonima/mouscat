package terminal

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/everdev/mack"
)

//Terminal is notifier that does only terminal notifications
type Terminal struct{}

//Notify terminal notifier
func (Terminal) Notify(data []byte, isError bool) error {
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

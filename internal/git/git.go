package git

import (
	"os"
	"os/exec"
)

func LsFunction() error { //func func_name receive return
	cmd := exec.Command("ls", "-alh")
	cmd.Stdout = os.Stdout
	return cmd.Run()

}

package git

import (
	"os"
	"os/exec"
)

func LsFunction(path string) error { //func func_name receive return
	cmd := exec.Command("ls", "-alh", path)
	cmd.Stdout = os.Stdout
	return cmd.Run()

}

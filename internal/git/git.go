package git

import (
	"os"
	"os/exec"
)

func ls(arg string, path string) error { //func func_name receive return
	cmd := exec.Command("ls", arg, path)
	cmd.Stdout = os.Stdout
	return cmd.Run()

}

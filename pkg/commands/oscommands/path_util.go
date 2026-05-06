package oscommands

import (
	"log"
	"os/exec"
)

func WslPathToWin(path string) string {
	pathCmd := exec.Command("wslpath", "-m", path)
	winPathBytes, pathErr := pathCmd.Output()
	if pathErr != nil {
		log.Fatal("Path conversion went bang")
	}
	return string(winPathBytes[:len(winPathBytes)-1])
}

func WinPathToWsl(path string) string {
	pathCmd := exec.Command("wslpath", "-u", path)
	wslPathBytes, pathErr := pathCmd.Output()
	if pathErr != nil {
		log.Fatal("Path conversion went bang")
	}
	return string(wslPathBytes[:len(wslPathBytes)-1])
}

package oscommands

import (
	"log"
	"os/exec"
	"slices"
)

func WslPathToWin(path string) string {
	if path == "" {
		return path
	}
	pathCmd := exec.Command("wslpath", "-m", path)
	winPathBytes, pathErr := pathCmd.Output()
	if pathErr != nil {
		log.Fatal("Path conversion went bang")
	}
	return string(winPathBytes[:len(winPathBytes)-1])
}

func WslPathArrayToWin(paths []string) []string {
	res := make([]string, 0, len(paths))
	for index, path := range paths {
		// slices.Insert(res, index, (WslPathToWin(path)))
		res[index] = WslPathToWin(path)
	}
	return res
}

func WinPathToWsl(path string) string {
	if path == "" {
		return path
	}
	pathCmd := exec.Command("wslpath", "-u", path)
	wslPathBytes, pathErr := pathCmd.Output()
	if pathErr != nil {
		log.Fatal("Path conversion went bang")
	}
	return string(wslPathBytes[:len(wslPathBytes)-1])
}

func WinPathArrayToWsl(paths []string) []string {
	res := make([]string, 0, len(paths))
	for index, path := range paths {
		slices.Insert(res, index, (WinPathToWsl(path)))
	}
	return res
}

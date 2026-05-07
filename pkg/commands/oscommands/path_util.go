package oscommands

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsWinPath(path string) bool {
	// Windows drive letter (C:\ or C:/)
	if len(path) >= 2 && path[1] == ':' {
		return true
	}

	// Windows UNC paths (\\server\share)
	if strings.HasPrefix(path, `\\`) {
		return true
	}

	return false
}

func WslPathToWin(path string) string {
	if path == "" || !filepath.IsAbs(path) {
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
	res := []string{}
	for _, path := range paths {
		res = append(res, WslPathToWin(path))
	}
	return res
}

func WinPathToWsl(path string) string {
	if path == "" || !IsWinPath(path) {
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
	res := []string{}
	for _, path := range paths {
		res = append(res, WinPathToWsl(path))
	}
	return res
}

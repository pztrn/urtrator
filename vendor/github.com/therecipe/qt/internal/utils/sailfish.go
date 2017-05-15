package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func VIRTUALBOX_DIR() string {
	if dir := os.Getenv("VIRTUALBOX_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return "C:\\Program Files\\Oracle\\VirtualBox"
	}
	var path, err = exec.LookPath("vboxmanage")
	if err != nil {
		Log.WithError(err).Error("failed to find vboxmanage in your PATH")
	}
	path = filepath.Dir(path)
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			Log.WithError(err).WithField("path", path).Fatal("can't resolve absolute path")
		}
	}
	return path
}

func SAILFISH_DIR() string {
	if dir := os.Getenv("SAILFISH_DIR"); dir != "" {
		return filepath.Clean(dir)
	}
	if runtime.GOOS == "windows" {
		return "C:\\SailfishOS"
	}
	return filepath.Join(os.Getenv("HOME"), "SailfishOS")
}

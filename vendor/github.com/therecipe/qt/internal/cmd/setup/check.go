package setup

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/utils"
)

func Check(target string, docker bool) {
	utils.Log.Infof("running: 'qtsetup check %v' [docker=%v]", target, docker)
	if docker {
		if _, err := exec.LookPath("docker"); err != nil {
			utils.Log.WithError(err).Fatal("failed to find docker, did you install docker?")
		}
		return
	}

	hash := "please install git"
	if _, err := exec.LookPath("git"); err == nil {
		cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
		cmd.Dir = utils.GoQtPkgPath()
		hash = strings.TrimSpace(utils.RunCmdOptional(cmd, "get git hash"))
	}

	vars := [][]string{
		{"GOOS", runtime.GOOS},
		{"GOARCH", runtime.GOARCH},
		{"GOVERSION", runtime.Version()},
		{"GOROOT", runtime.GOROOT()},
		{"GOPATH", utils.MustGoPath()},
		{"GOBIN", utils.MustGoBin()},
		{"QT_HASH", hash},
		{"QT_VERSION", utils.QT_VERSION()},
		{"QT_VERSION_MAJOR", utils.QT_VERSION_MAJOR()},
		{"QT_DIR", utils.QT_DIR()},
		{"QT_STUB", fmt.Sprint(utils.QT_STUB())},
		{"QT_DEBUG", fmt.Sprint(utils.QT_DEBUG())},
		{"QT_QMAKE_DIR", utils.QT_QMAKE_DIR()},
		{"QT_WEBKIT", fmt.Sprint(utils.QT_WEBKIT())},
	}

	if utils.CI() {
		vars = append(vars, [][]string{
			{"CI", fmt.Sprint(utils.CI())},
		}...)
	}

	switch target {
	case "darwin", "ios", "ios-simulator":
		vars = append(vars, [][]string{
			{"QT_HOMEBREW", fmt.Sprint(utils.QT_HOMEBREW())},
			{"XCODE_DIR", utils.XCODE_DIR()},
			//{"IPHONEOS_SDK_DIR", utils.IPHONEOS_SDK_DIR()},               //TODO: re-add with absolute path
			//{"IPHONESIMULATOR_SDK_DIR", utils.IPHONESIMULATOR_SDK_DIR()}, //TODO: re-add with absolute path
		}...)

		if _, err := exec.LookPath("clang++"); err != nil {
			utils.Log.WithError(err).Panic("failed to find clang++, did you install Xcode?")
		}

	case "linux":
		vars = append(vars, [][]string{
			{"QT_DISTRO", utils.QT_DISTRO()},
			{"QT_PKG_CONFIG", fmt.Sprint(utils.QT_PKG_CONFIG())},
		}...)

		if utils.QT_PKG_CONFIG() {
			vars = append(vars, [][]string{
				{"QT_DOC_DIR", utils.QT_DOC_DIR()},
				{"QT_MISC_DIR", utils.QT_MISC_DIR()},
			}...)
		}

		if _, err := exec.LookPath("g++"); err != nil {
			utils.Log.WithError(err).Panic("failed to find g++, did you install g++?")
		}

	case "windows":
		if runtime.GOOS == target {
			vars = append(vars, [][]string{
				{"QT_MSYS2", fmt.Sprint(utils.QT_MSYS2())},
			}...)

			if utils.QT_MSYS2() {
				vars = append(vars, [][]string{
					{"QT_MSYS2_DIR", utils.QT_MSYS2_DIR()},
					{"QT_MSYS2_ARCH", utils.QT_MSYS2_ARCH()},
					{"QT_MSYS2_STATIC", fmt.Sprint(utils.QT_MSYS2_STATIC())},
				}...)
			}

			if _, err := exec.LookPath("g++"); err != nil && !utils.QT_MSYS2() {
				utils.Log.WithError(err).Panic("failed to find g++, did you add the directory that contains g++ to your PATH?")
			}
		} else {
			vars = append(vars, [][]string{
				{"QT_MXE_DIR", utils.QT_MXE_DIR()},
				{"QT_MXE_ARCH", utils.QT_MXE_ARCH()},
				{"QT_MXE_STATIC", fmt.Sprint(utils.QT_MXE_STATIC())},
			}...)
		}

	case "android", "android-emulator":
		vars = append(vars, [][]string{
			{"JDK_DIR", utils.JDK_DIR()},
			{"ANDROID_SDK_DIR", utils.ANDROID_SDK_DIR()},
			{"ANDROID_NDK_DIR", utils.ANDROID_NDK_DIR()},
		}...)

	case "sailfish", "sailfish-emulator":
		vars = append(vars, [][]string{
			{"VIRTUALBOX_DIR", utils.VIRTUALBOX_DIR()},
			{"SAILFISH_DIR", utils.SAILFISH_DIR()},
		}...)

	case "rpi1", "rpi2", "rpi3":
		vars = append(vars, [][]string{
			{"RPI_TOOLS_DIR", utils.RPI_TOOLS_DIR()},
			{"RPI_COMPILER", utils.RPI_COMPILER()},
		}...)
	}

	for _, v := range vars {
		var set string
		if _, ok := os.LookupEnv(v[0]); ok {
			set = "*"
		}
		utils.Log.Infof("%v:%v%v%v'%v'", v[0], strings.Repeat(" ", 25-len(v[0])), set, strings.Repeat(" ", 3-len(set)), v[1])
		if !strings.HasSuffix(v[0], "_DIR") {
			continue
		}
		if v[0] == "QT_DIR" && (utils.QT_HOMEBREW() || utils.QT_MSYS2() || utils.QT_PKG_CONFIG() || utils.MSYS_DOCKER()) {
			continue
		}
		if _, err := ioutil.ReadDir(v[1]); err != nil && v[1] != "" {
			utils.Log.WithError(err).Panicf("failed to find %v (%v)", v[0], v[1])
		}
	}
}

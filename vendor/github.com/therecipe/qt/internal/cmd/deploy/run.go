package deploy

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/utils"
)

func run(target, name, depPath string) {
	switch target {
	case "android", "android-emulator":
		if utils.ExistsFile(filepath.Join(depPath, "build-debug.apk")) {
			exec.Command(filepath.Join(utils.ANDROID_SDK_DIR(), "platform-tools", "adb"), "install", "-r", filepath.Join(depPath, "build-debug.apk")).Start()
		} else {
			exec.Command(filepath.Join(utils.ANDROID_SDK_DIR(), "platform-tools", "adb"), "install", "-r", filepath.Join(depPath, "build-release-signed.apk")).Start()
		}

	case "ios-simulator":
		//TODO: parse list of available simulators
		utils.RunCmdOptional(exec.Command("xcrun", "instruments", "-w", "iPhone 7 Plus (10.3)#"), "start simulator")
		utils.RunCmdOptional(exec.Command("xcrun", "simctl", "uninstall", "booted", filepath.Join(depPath, "main.app")), "uninstall old app")
		utils.RunCmdOptional(exec.Command("xcrun", "simctl", "install", "booted", filepath.Join(depPath, "main.app")), "install new app")
		utils.RunCmdOptional(exec.Command("xcrun", "simctl", "launch", "booted", fmt.Sprintf("com.identifier.%v", name)), "start app")

	case "darwin":
		exec.Command("open", filepath.Join(depPath, fmt.Sprintf("%v.app", name))).Start()

	case "linux":
		exec.Command(filepath.Join(depPath, fmt.Sprintf("%v.sh", name))).Start()

	case "windows":
		if runtime.GOOS == target {
			exec.Command(filepath.Join(depPath, name+".exe")).Start()
		} else {
			exec.Command("wine", filepath.Join(depPath, name+".exe")).Start()
		}

	case "sailfish-emulator":
		utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "registervm", filepath.Join(utils.SAILFISH_DIR(), "emulator", "SailfishOS Emulator", "SailfishOS Emulator.vbox")), "register vm")
		utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "sharedfolder", "add", "SailfishOS Emulator", "--name", "GOPATH", "--hostpath", utils.MustGoPath(), "--automount"), "mount GOPATH")

		if runtime.GOOS == "windows" {
			utils.RunCmdOptional(exec.Command(filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "SailfishOS Emulator"), "start emulator")
		} else {
			utils.RunCmdOptional(exec.Command("nohup", filepath.Join(utils.VIRTUALBOX_DIR(), "vboxmanage"), "startvm", "SailfishOS Emulator"), "start emulator")
		}

		time.Sleep(10 * time.Second)

		err := sailfish_ssh("2223", "nemo", "sudo", "rpm", "-i", "--force", strings.Replace(strings.Replace(depPath, utils.MustGoPath(), "/media/sf_GOPATH", -1)+"/*.rpm", "\\", "/", -1))
		if err != nil {
			utils.Log.WithError(err).Errorf("failed to install %v for %v", name, target)
		}

		err = sailfish_ssh("2223", "nemo", "nohup", "/usr/bin/harbour-"+name, ">", "/dev/null", "2>&1", "&")
		if err != nil {
			utils.Log.WithError(err).Errorf("failed to run %v for %v", name, target)
		}
	}
}

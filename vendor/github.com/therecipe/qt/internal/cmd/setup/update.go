package setup

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/therecipe/qt/internal/utils"
)

func Update() {
	utils.Log.Info("running: 'qtsetup update'")

	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/cmd/..."), "run \"go clean cmd\"")
	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/internal/..."), "run \"go clean internal\"")

	fetch := exec.Command("git", "fetch", "--all")
	fetch.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(fetch, "run \"git fetch\"")

	checkoutCmd := exec.Command("git", "checkout", "--", utils.GoQtPkgPath("cmd"))
	checkoutCmd.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(checkoutCmd, "run \"git checkout cmd\"")

	checkoutInternal := exec.Command("git", "checkout", "--", utils.GoQtPkgPath("internal"))
	checkoutInternal.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "therecipe", "qt")
	utils.RunCmd(checkoutInternal, "run \"git checkout internal\"")

	utils.RunCmd(exec.Command("go", "install", "-v", fmt.Sprintf("github.com/therecipe/qt/cmd/...")), "run \"go install\"")

	Prep()
}

func Upgrade() {
	utils.Log.Info("running: 'qtsetup upgrade'")

	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/therecipe/qt/..."), "run \"go clean\"")
	utils.RemoveAll(utils.GoQtPkgPath())
	utils.RunCmd(exec.Command("go", "get", "-v", fmt.Sprintf("github.com/therecipe/qt/cmd/...")), "run \"go get\"")
}

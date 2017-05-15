package templater

import (
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(m, target string, mode int) {
	if !parser.ShouldBuildForTarget(m, target) {
		utils.Log.WithField("module", m).Debug("skip generation")
		return
	}
	utils.Log.WithField("module", m).Debug("generating")

	var suffix = func() string {
		switch m {
		case "AndroidExtras":
			{
				return "_android"
			}

		case "Sailfish":
			{
				return "_sailfish"
			}

		default:
			{
				return ""
			}
		}
	}()

	if mode == NONE {
		utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m)))
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
	}

	if mode == MINIMAL {
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), CppTemplate(m, mode, target))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), HTemplate(m, mode))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), GoTemplate(m, false, mode, m, target))

		if !UseStub(false, "Qt"+m, mode) {
			CgoTemplate(m, "", target, mode, m)
		}

		return
	}

	if m == "AndroidExtras" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(m), "utils-androidextras_android.go"), utils.Load(utils.GoQtPkgPath("internal", "binding", "files", "utils-androidextras_android.go")))
	}

	if !UseStub(false, "Qt"+m, mode) {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".cpp"), CppTemplate(m, mode, target))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".h"), HTemplate(m, mode))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".go"), GoTemplate(m, false, mode, m, target))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+".go"), GoTemplate(m, suffix != "", mode, m, target))

	if !UseStub(false, "Qt"+m, mode) {
		CgoTemplate(m, "", target, mode, m)
	}
}

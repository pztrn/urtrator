package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/therecipe/qt/internal/cmd"
	"github.com/therecipe/qt/internal/cmd/setup"

	"github.com/therecipe/qt/internal/utils"
)

func main() {
	flag.Usage = func() {
		println("Usage: qtsetup [-debug] [mode] [target]\n")

		println("Flags:\n")
		flag.PrintDefaults()
		print("\n")

		println("Modes:\n")
		modes := []struct{ name, desc string }{
			{"prep", "try to symlink tooling into the PATH"},
			{"check", "do some basic env checks"},
			{"generate", "generate the code for all packages"},
			{"install", "run go install for all packages"},
			{"test", "build some examples"},
			{"full", "run all of the above"},
			{"help", "print help"},
			{"update", "update 'cmd' and 'internal/cmd'"},
			{"upgrade", "update everything"},
		}
		for _, mode := range modes {
			fmt.Printf("  %v%v%v\n", mode.name, strings.Repeat(" ", 12-len(mode.name)), mode.desc)
		}
		print("\n")

		println("Targets:\n")
		//TODO:
		print("\n")

		os.Exit(0)
	}

	var docker bool
	flag.BoolVar(&docker, "docker", false, "run command inside docker container")

	if cmd.ParseFlags() {
		flag.Usage()
	}

	mode := "full"
	target := runtime.GOOS

	switch flag.NArg() {
	case 0:
	case 1:
		mode = flag.Arg(0)
	case 2:
		mode = flag.Arg(0)
		target = flag.Arg(1)
	default:
		flag.Usage()
	}

	if target == "desktop" {
		target = runtime.GOOS
	}

	utils.CheckBuildTarget(target)

	switch mode {
	case "full":
		setup.Prep()
		setup.Check(target, docker)
		setup.Generate(target, docker)
		setup.Install(target, docker)
		setup.Test(target, docker)
	case "prep":
		setup.Prep()
	case "check":
		setup.Check(target, docker)
	case "generate":
		setup.Generate(target, docker)
	case "install":
		setup.Install(target, docker)
	case "test":
		setup.Test(target, docker)
	case "update":
		setup.Update()
	case "upgrade":
		setup.Upgrade()
	default:
		flag.Usage()
	}
}

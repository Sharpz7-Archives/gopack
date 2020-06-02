package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var fileFlag bool
var devFlag bool

const (
	actionInstall   = "install"
	actionUninstall = "uninstall"
)

// Create Flags needed
func init() {
	flag.BoolVar(&fileFlag, "file", false, "Install from File")
	flag.BoolVar(&devFlag, "dev", false, "Install Developer Packages")
}

func main() {
	// Parses flags and removes them from args
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatal("You can either chose to 'uninstall' or 'install'!")
	} else {
		var act = flag.Args()[0]

		// Load gopack file
		goFile, err := loadFile()
		if err != nil {
			fmt.Println("No gopack.yml was found... generating new one")
			genFile()
			return
		}

		// Check Versions
		err = checkVersions(goFile)
		clientErrCheck(err, "Gofile Has incorrect version")

		// Check Subcommands
		if (act != actionInstall) && (act != actionUninstall) {
			log.Fatal("You can either chose to 'uninstall' or 'install'!")
		}

		if fileFlag {
			actionFile(goFile, act)
		} else {
			action(goFile, act)
		}

	}
	return
}

func clientErrCheck(e error, msg string) {
	if e != nil {
		fmt.Println(e)
		log.Fatal(msg)
	}
}

func checkVersions(goFile goPack) error {
	out, _ := exec.Command("go", "version").Output()
	outString := string(out[13:17])

	if outString != goFile.GoVersion {
		return errors.New("versions do not match")
	}

	return nil
}

func removePackage(targetPkg string, pkgs []string) []string {

	finalGoFile := make([]string, 0)
	for _, pkg := range pkgs{
		if pkg != targetPkg {
			finalGoFile = append(finalGoFile, pkg)
		}
	}

	return finalGoFile
}

func pkgCommand(name string, arg ...string) {
	out, err := exec.Command(name, arg...).CombinedOutput()
	clientErrCheck(err, string(out))
}

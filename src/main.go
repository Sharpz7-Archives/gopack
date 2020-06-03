package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
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

    // Creates Helper Function
	flag.Usage = func() {
		fmt.Println(`
Args of Gopack:

    - install
    - uninstall

You can also manually edit the gopack.yml file and use the file flag
		`)

		flag.PrintDefaults()
	}
}

func main() {
	// Parses flags and removes them from args
	flag.Parse()

    // Checks if an arg has been selected
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
		check(err, "Gofile Has incorrect version")

		// Check Subcommands
		if (act != actionInstall) && (act != actionUninstall) {
			log.Fatal("You can either chose to 'uninstall' or 'install'!")
		}

        // If file flag is on apply file action
		if fileFlag {
			actionFile(goFile, act)
		} else {
			action(goFile, act)
		}

	}
	return
}

func check(e error, msg string) {
    // Try and get SHARPDEV var
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot read enviroment")
	}

	if e != nil {
		if os.Getenv("SHARPDEV") == "TRUE" {
			fmt.Println(e)
		}
		log.Fatal(msg)
	}
}

// Check if version in gopack.yml matches whats installed
func checkVersions(goFile goPack) error {
	out, _ := exec.Command("go", "version").Output()
	outString := string(out[13:17])

	if outString != goFile.GoVersion {
		return errors.New("versions do not match")
	}

	return nil
}

// Remove a package from pkgs slice
func removePackage(targetPkg string, pkgs []string) []string {

	finalGoFile := make([]string, 0)
	for _, pkg := range pkgs {
		if pkg != targetPkg {
			finalGoFile = append(finalGoFile, pkg)
		}
	}

	return finalGoFile
}


// Execute a command to do with pkg install or uninstall
func pkgCommand(name string, arg ...string) {
	out, err := exec.Command(name, arg...).CombinedOutput()
	check(err, string(out))
}

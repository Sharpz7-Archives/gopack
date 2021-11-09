package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func actionFile(goFile goPack, option string) {
	var pkgs []string

	// Add Dev Packages if needed
	if devFlag {
		pkgs = append(goFile.Packages, goFile.DevPackages...)
	} else {
		pkgs = goFile.Packages
	}

	for _, pkg := range pkgs {
		if option == actionInstall {
			// Install
			pkgCommand("go", "get", pkg)
		} else {
			// Uninstall
			pkgCommand("rm", "-r", os.Getenv("GOPATH")+"/src/"+pkg)
		}
	}

	fmt.Println(strings.Title(option)+"ed", "all Packages from File")
}

func action(goFile goPack, option string) {
	pkg := flag.Args()[1]

	if option == actionInstall {
		// Installing
		pkgCommand("go", "get", pkg)

		// If dev enabled add to dev packages
		if devFlag {
			if !stringInSlice(pkg, goFile.DevPackages) {
				goFile.DevPackages = append(goFile.DevPackages, pkg)
			} else {
				fmt.Println("Already Installed " + pkg)
			}

		} else {
			if !stringInSlice(pkg, goFile.Packages) {
				goFile.Packages = append(goFile.Packages, pkg)
			} else {
				fmt.Println("Already Installed " + pkg)
			}

		}

	} else {

		// Remove package from project package
		pkgCommand("rm", "-r", os.Getenv("GOPATH")+"/src/"+pkg)
		if devFlag {
			goFile.DevPackages = removePackage(pkg, goFile.DevPackages)
		} else {
			goFile.Packages = removePackage(pkg, goFile.Packages)
		}

	}

	// Save to disk
	err := saveFile(goFile)
	check(err, "Failed to remove package to gopack.yml")

	fmt.Println(strings.Title(option)+"ed", "Package", pkg)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

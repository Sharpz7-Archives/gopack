package main

import (
	"flag"
	"fmt"
	"os"
)

func actionFile(goFile goPack, option string) {
	var pkgs []string

	if devFlag {
		pkgs = append(goFile.Packages, goFile.DevPackages...)
	} else {
		pkgs = goFile.Packages
	}

	for _, pkg := range pkgs {
		if option == actionInstall {
			pkgCommand("go", "get", pkg)
		} else {
			pkgCommand("rm", "-r", os.Getenv("GOPATH")+"/src/"+pkg)
		}
	}

	if option == actionUninstall {
		goFile.Packages = make([]string, 0)
		if devFlag {
			goFile.DevPackages = make([]string, 0)
		}
		err := saveFile(goFile)
		clientErrCheck(err, "Failed to remove all packages from gopack.yml")
	}

	fmt.Println(option+"ed", "all Packages from File")
}

func action(goFile goPack, option string) {
	pkg := flag.Args()[1]

	if option == actionInstall {
		pkgCommand("go", "get", pkg)

		if devFlag {
			goFile.DevPackages = append(goFile.DevPackages, pkg)
		} else {
			goFile.Packages = append(goFile.Packages, pkg)
		}

	} else {
		pkgCommand("rm", "-r", os.Getenv("GOPATH")+"/src/"+pkg)
		if devFlag {
			goFile.DevPackages = removePackage(pkg, goFile.DevPackages)
		} else {
			goFile.Packages = removePackage(pkg, goFile.Packages)
		}

	}

	err := saveFile(goFile)
	clientErrCheck(err, "Failed to remove package to gopack.yml")

	fmt.Println(option+"ed", "Package")
}

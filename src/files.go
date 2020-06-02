package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func genFile() {
	out, _ := exec.Command("go", "version").Output()
	outString := string(out[13:17])

	testfile := goPack{
		GoVersion: outString}

	err := saveFile(testfile)
	clientErrCheck(err, "Failed to generate gopack.yml")

	fmt.Println("Created gopack.yml")
}

func loadFile() (goPack, error) {
	f, readErr := ioutil.ReadFile("./gopack.yml")
	var goFile goPack
	marshErr := yaml.Unmarshal(f, &goFile)
	if marshErr != nil || readErr != nil {
		return goPack{}, errors.New("failed to load file")
	}

	return goFile, nil
}

func saveFile(goFile goPack) error {
	yamlData, marshErr := yaml.Marshal(goFile)
	clientErrCheck(marshErr, "Failed to Convert to Yaml")
	writeErr := ioutil.WriteFile("./gopack.yml", yamlData, 0644)

	if marshErr != nil || writeErr != nil {
		return errors.New("failed to save file")
	}

	return nil
}

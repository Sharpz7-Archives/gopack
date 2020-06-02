package main

type goPack struct {
	GoVersion   string   `yml:"goversion"`
	Packages    []string `yml:"packages"`
	DevPackages []string `yml:"devpackages"`
}

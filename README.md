[![CircleCI](https://circleci.com/gh/Sharpz7/gopack.svg?style=svg)](https://circleci.com/gh/Sharpz7/gopack)

# GOPACK | Simple GOlang Package Manager

GoPack aims to take some of the features of tools like `pipenv` to golang.

Here are some of its features:

- Version Checking: If the versions of gopack and the go service don't match, it will error.
- Uninstalling: Using GOPATH, packages can be uninstalled.
- Packages you install with GoPack will automatically be added to gopack.yml!

# Example Config
```yml
goversion: "1.12"
packages:
- gopkg.in/yaml.v2
- github.com/joho/godotenv
devpackages: []

```

# Installation
On linux, just run:
```console
╭─adam@box ~/
╰─➤  sudo curl -s -L https://github.com/Sharpz7/gopack/releases/download/0.1.2/install.sh | sudo bash
```

## Command Options

On linux, just run:
```console
╭-adam@box ~/
╰─➤  gopack --help

Args of Gopack:

    - install
    - uninstall

You can also manually edit the gopack.yml file and use the file flag

  -dev
        Install Developer Packages
  -file
        Install from File
```

## Maintainers

- [Adam McArthur](https://adam.mcaq.me)
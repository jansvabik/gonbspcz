# (Czech-only) Non-breakable space auto-adder for Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/jansvabik/gonbspcz)](https://goreportcard.com/report/github.com/jansvabik/gonbspcz)
![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)

Go package for replacing spaces by unbreakable spaces where they should be according to the Czech typographical rules. This package should automatically add a non-breakable space in all of these cases:

* after digits
* after one-letter words
* after ordinals
* after initials

## Installation
You can add this package to your workspace by installing it and then importing it.

1. Run command `go get github.com/jansvabik/gonbspcz`
2. Import package `import "github.com/jansvabik/gonbspcz"`

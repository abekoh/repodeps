package main

import (
	"github.com/abekoh/repodeps"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(repodeps.Analyzer) }

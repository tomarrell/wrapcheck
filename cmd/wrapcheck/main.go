package main

import (
	"github.com/tomarrell/wrapcheck/wrapcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(wrapcheck.Analyzer)
}

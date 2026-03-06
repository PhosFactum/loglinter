// Точка входа в... Log-linter :D
package main

import (
	"github.com/PhosFactum/loglinter/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}

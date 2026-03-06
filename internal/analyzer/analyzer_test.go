package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	// Тестим на файлах прямиком из testdata директории
	analysistest.Run(t, analysistest.TestData(), Analyzer, "example_err", "example_ok")
}

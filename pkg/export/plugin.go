package export

import (
    "github.com/yodaskilledme/errorstyle-analyzer/pkg/analyzer"
    "golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
    return analyzer.ErrStyleAnalyzer
}

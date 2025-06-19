package main

import (
    "github.com/yodaskilledme/errorstyle-analyzer/pkg/analyzer"
    "golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
    return []*analysis.Analyzer{
        analyzer.ErrStyleAnalyzer,
    }
}

var AnalyzerPlugin analyzerPlugin

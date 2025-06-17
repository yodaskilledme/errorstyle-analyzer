package main

import (
    "github.com/yodaskilledme/errorstyle-analyzer/pkg/analyzer"
    "golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
    singlechecker.Main(analyzer.ErrStyleAnalyzer)
}

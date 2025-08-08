package export

import (
    "golang.org/x/tools/go/analysis"
    "github.com/golangci/plugin-module-register/register"
    "github.com/yodaskilledme/errorstyle-analyzer/pkg/analyzer"
)

func NewAnalyzer() *analysis.Analyzer {
    return analyzer.ErrStyleAnalyzer
}

type errStylePlugin struct{}

func (*errStylePlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
    return []*analysis.Analyzer{analyzer.ErrStyleAnalyzer}, nil
}

func (*errStylePlugin) GetLoadMode() string {
    return register.LoadModeTypesInfo
}

func newErrStylePlugin(conf any) (register.LinterPlugin, error) {
    // Применяем настройки из .golangci.yml -> linters.settings.custom.errstyle.settings
    if m, ok := conf.(map[string]any); ok {
        if v, ok := m["op_name"].(string); ok && v != "" {
            _ = analyzer.ErrStyleAnalyzer.Flags.Set("op_name", v)
        }
        if v, ok := m["errType"].(string); ok && v != "" {
            _ = analyzer.ErrStyleAnalyzer.Flags.Set("errType", v)
        }
    }
    return &errStylePlugin{}, nil
}

func init() {
    register.Plugin("errstyle", newErrStylePlugin)
}

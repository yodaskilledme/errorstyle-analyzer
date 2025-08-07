package export

import (
    "golang.org/x/tools/go/analysis"
    "github.com/yodaskilledme/errorstyle-analyzer/pkg/analyzer"
    "github.com/golangci/plugin-module-register/register"
)

// NewAnalyzer оставляем для обратной совместимости с legacy-API.
func NewAnalyzer() *analysis.Analyzer {
    return analyzer.ErrStyleAnalyzer
}

// errStylePlugin реализует интерфейс register.LinterPlugin.
type errStylePlugin struct{}

// Возвращаем единственный анализатор плагина.
func (*errStylePlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
    return []*analysis.Analyzer{analyzer.ErrStyleAnalyzer}, nil
}

// Заявляем режим загрузки: наш анализатор использует типовую информацию.
func (*errStylePlugin) GetLoadMode() string {
    return register.LoadModeTypesInfo
}

// Конструктор плагина: сюда можно передавать настройки из golangci.yml.
func newErrStylePlugin(conf any) (register.LinterPlugin, error) {
    return &errStylePlugin{}, nil
}

func init() {
    // Регистрируем плагин под именем "errstyle".
    register.Plugin("errstyle", newErrStylePlugin)
}

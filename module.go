package core

type ModuleInterface interface {
	Init(app *App)
	Load()
	Export() *Module
}

type Module struct{}

func (m *Module) Init(app *App) {}

func (m *Module) Load() {}

func (m *Module) Export() *Module {
	return m
}

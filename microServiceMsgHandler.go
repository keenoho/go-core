package core

type MicroServiceMsgHandlerInterface interface {
	Init()
	URLMapping()
	Resigter(app *MicroServiceApp)
}

type MicroServiceMsgHandler struct{}

func (m *MicroServiceMsgHandler) URLMapping() {}

func (m *MicroServiceMsgHandler) Init() {}

func (m *MicroServiceMsgHandler) Mapping(path string, method string, fn ControllerFunc) {}

func (m *MicroServiceMsgHandler) Resigter(app *MicroServiceApp) {}

func ResigterMicroAppController(app *MicroServiceApp, execControllers ...MicroServiceMsgHandlerInterface) {
}

func ControllerToMicroServiceMsgHandler(controller ControllerFunc) {

}

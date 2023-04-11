package core

type RegisterCenter struct {
}

func (rc *RegisterCenter) Init() {

}

func (rc *RegisterCenter) GetServiceAddress(targetServiceName string) string {
	return "0.0.0.0:8001"
}

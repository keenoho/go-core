package service

type MicroServiceRegisterInterface interface {
}

type MicroServiceRegister struct{}

func (r *MicroServiceRegister) AddService() {}

func (r *MicroServiceRegister) RemoveService() {}

func (r *MicroServiceRegister) UpdateService() {}

func (r *MicroServiceRegister) GetService() {}

func (r *MicroServiceRegister) GetAllService() {}

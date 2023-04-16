package core

import "time"

type RegisterService struct {
	Logger
	Key           string
	Value         string
	App           string
	Id            string
	Address       string
	Health        bool
	LastCheckTime time.Time
}

func (rs *RegisterService) CheckHealth() bool {
	return true
}

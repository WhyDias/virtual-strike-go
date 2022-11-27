package service

import (
	"virtual-strike-backend-go/pkg/modules"
)

type Time interface {
	TimeLogic(jsonInput modules.TimeRequest) (code int, any modules.Response)
}

type Service struct {
	Time
}

func NewService() *Service {
	return &Service{
		Time: NewTimeService(),
	}
}

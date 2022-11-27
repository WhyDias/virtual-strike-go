package service

import (
	"virtual-strike-backend-go/pkg/modules"
)

type Time interface {
	TimeLogic(jsonInput modules.TimeRequest) (code int, any modules.Response)
}

type Upload interface {
	UploadLogic(jsonInput modules.UploadRequest) (code int, any modules.Response)
}

type Service struct {
	Time
	Upload
}

func NewService() *Service {
	return &Service{
		Time:   NewTimeService(),
		Upload: NewUploadService(),
	}
}

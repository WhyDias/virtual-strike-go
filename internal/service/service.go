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

type WorkDayInfo interface {
	WorkDayInfoLogic(jsonInput modules.WorkDayInfoRequest) (code int, any modules.Response)
}

type Point interface {
	PointLogic(jsonInput modules.PointRequest) (code int, any []modules.PointResponse)
}

type Logging interface {
	LoggingLogic(jsonInput modules.LoggingRequest) (code int, any modules.Response)
}

type Service struct {
	Time
	Upload
	WorkDayInfo
	Point
	Logging
}

func NewService() *Service {
	return &Service{
		Time:        NewTimeService(),
		Upload:      NewUploadService(),
		WorkDayInfo: NewWorkDayInfoService(),
		Point:       NewPointService(),
		Logging:     NewLoggingService(),
	}
}

package service

type Register interface {
}

type Login interface {
}

type Service struct {
	Register
	Login
}

func NewService() *Service {
	return &Service{
		Register: NewRegisterService(),
		Login:    NewLoginService(),
	}
}

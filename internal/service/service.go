package service

type Search interface {
}

type Service struct {
	Search
}

func NewService() *Service {
	return &Service{
		Search: NewSearchService(),
	}
}

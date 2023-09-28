package sales

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	Create(sales *domain.Sales) error
	ReadAll() ([]*domain.Sales, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(sales *domain.Sales) error {
	_, err := s.r.Create(sales)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Sales, error) {
	return s.r.ReadAll()
}

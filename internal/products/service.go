package products

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	Create(product *domain.Product) error
	ReadAll() ([]*domain.Product, error)
	InsertMany(products []*domain.Product) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(product *domain.Product) error {
	_, err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Product, error) {
	return s.r.ReadAll()
}

func (s *service) InsertMany(products []*domain.Product) error {
	return s.r.InsertMany(products)
}

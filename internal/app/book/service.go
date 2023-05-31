package book

import "main/internal/domain"

type IService interface {
	CreateBook(book domain.Book) (*domain.Book, error)
	GetBooks() ([]domain.Book, error)
	GetById(id string) (*domain.Book, error)
	RemoveById(id string) error
	UpdateById(id string, book domain.Book) (*domain.Book, error)
}

type Service struct {
	bookRepository IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		bookRepository: repo,
	}
}

func (service *Service) CreateBook(book domain.Book) (*domain.Book, error) {
	response, err := service.bookRepository.CreateBookRepo(book)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Service) GetBooks() ([]domain.Book, error) {
	response, err := service.bookRepository.GetBooksRepo()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Service) GetById(id string) (*domain.Book, error) {
	response, err := service.bookRepository.GetByIdRepo(id)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Service) RemoveById(id string) error {
	err := service.bookRepository.RemoveByIdRepo(id)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) UpdateById(id string, book domain.Book) (*domain.Book, error) {
	response, err := service.bookRepository.UpdateByIdRepo(id, book)
	if err != nil {
		return nil, err
	}

	return response, nil
}

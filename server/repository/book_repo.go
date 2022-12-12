package repository

import (
	"github.com/AntonioTrupac/socialHabitsTracker/graph/customTypes"
	"github.com/AntonioTrupac/socialHabitsTracker/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(bookInput *customTypes.BookInput) (*models.Book, error)
	UpdateBook(bookInput *customTypes.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*customTypes.Book, error)
}

type BookService struct {
	Db *gorm.DB
}

var _ BookRepository = &BookService{}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		Db: db,
	}
}

func (b BookService) CreateBook(bookInput *customTypes.BookInput) (*models.Book, error) {
	book := &models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}

	err := b.Db.Create(&book).Error

	return book, err
}

func (b BookService) UpdateBook(bookInput *customTypes.BookInput, id int) error {
	book := models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
		ID:        uint(id),
	}

	err := b.Db.Model(&book).Where("id = ?", id).Updates(book).Error
	return err
}

func (b BookService) DeleteBook(id int) error {
	book := &models.Book{}

	err := b.Db.Delete(book, id).Error
	return err
}

func (b BookService) GetOneBook(id int) (*models.Book, error) {
	book := &models.Book{}

	err := b.Db.Where("id = ?", id).First(book).Error
	return book, err
}

func (b BookService) GetAllBooks() ([]*customTypes.Book, error) {
	var books []*customTypes.Book

	err := b.Db.Find(&books).Error
	return books, err
}

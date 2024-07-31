package services

import (
	"errors"
	"main/models"
)

type LibraryService struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (s *LibraryService) AddBook(book models.Book) {
	s.books[book.ID] = book
}

func (s *LibraryService) RemoveBook(bookID int) {
	delete(s.books, bookID)
}

func (s *LibraryService) BorrowBook(bookID int, memberID int) error {
	book, exists := s.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := s.members[memberID]
	if !exists {
		member = models.Member{ID: memberID}
	}
	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	s.books[bookID] = book
	s.members[memberID] = member
	return nil
}

func (s *LibraryService) ReturnBook(bookID int, memberID int) error {
	book, exists := s.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is already available")
	}

	member, exists := s.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	book.Status = "Available"
	s.books[bookID] = book
	s.members[memberID] = member
	return nil
}

func (s *LibraryService) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range s.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (s *LibraryService) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := s.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}

package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		member = &models.Member{
			ID:   memberID,
			Name: "Member",
		}
		l.Members[memberID] = member
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks =
		append(member.BorrowedBooks, book)

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {

			member.BorrowedBooks =
				append(member.BorrowedBooks[:i],
					member.BorrowedBooks[i+1:]...)

			book.Status = "Available"
			l.Books[bookID] = book

			return nil
		}
	}

	return errors.New("book not borrowed by member")
}

func (l *Library) ListAvailableBooks() []models.Book {

	var available []models.Book

	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}

	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {

	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}

	return member.BorrowedBooks
}
package services

import (
	"errors"
	"library_management/models"
	"sync"
)

type LibraryManager interface {
	AddMember(member models.Member)
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberId int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error
}

type Library struct {
	books      map[int]*models.Book
	members    map[int]*models.Member
	mutex      sync.Mutex
	reserveReq chan models.ReservationRequest
}

func (l *Library) Books() map[int]*models.Book {
	return l.books
}
func (l *Library) Members() map[int]*models.Member {
	return l.members
}
func (l *Library) Mutex() *sync.Mutex {
	return &l.mutex
}
func (l *Library) ReservationChannel() chan models.ReservationRequest {
	return l.reserveReq
}

func NewLibrary() *Library {
	return &Library{
		books:      make(map[int]*models.Book),
		members:    make(map[int]*models.Member),
		reserveReq: make(chan models.ReservationRequest, 100),
	}
}

func (l *Library) GetBook(bookID int) (*models.Book, error) {
	book, exists := l.books[bookID]
	if !exists {
		return nil, errors.New("book Not found")
	}
	return book, nil
}
func (l *Library) AddMember(member models.Member) {
	l.members[member.ID] = &member
}

func (l *Library) GetMember(memberID int) (*models.Member, error) {
	member, exists := l.members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member, nil
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.books[book.ID] = &book
}

func (l *Library) RemoveBook(bookID int) error {
	if _, err := l.GetBook(bookID); err != nil {
		return err
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, err := l.GetBook(bookID)
	if err != nil {
		return err
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}
	if book.Status == "Reserved" && book.ReservedBy != memberID {
		return errors.New("book is reserved by another person")
	}

	member, err := l.GetMember(memberID)
	if err != nil {
		return err
	}

	book.Status = "Borrowed"
	book.ReservedBy = 0
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, err := l.GetBook(bookID)
	if err != nil {
		return err
	}
	member, err := l.GetMember(memberID)
	if err != nil {
		return err
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book.Status = "Available"
			return nil
		}
	}
	return errors.New("book is not borrowed by this member")
}

func (l *Library) ListAvailableBooks() []*models.Book {
	var availableBooks []*models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) ([]*models.Book, error) {
	member, err := l.GetMember(memberID)
	if err != nil {
		return nil, err
	}
	return member.BorrowedBooks, nil
}

func (l *Library) ListAllBorrowedBooks() []*models.Book {
	var boorrwedBooks []*models.Book
	for _, book := range l.books {
		if book.Status == "Borrowed" {
			boorrwedBooks = append(boorrwedBooks, book)
		}
	}
	return boorrwedBooks
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	book, exists := l.books[bookID]
	if _, err := l.GetMember(memberID); err != nil {
		return errors.New("member not found")
	}
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != "Available" {
		return errors.New("book is not available for reservation")
	}

	book.Status = "Reserved"
	book.ReservedBy = memberID
	l.reserveReq <- models.ReservationRequest{BookID: bookID, MemberID: memberID}
	return nil
}

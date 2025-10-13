package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// main tests
// main tests

func main() {

	lib := Library{
		Name:  "My Library",
		Books: []BookList{},
		Users: make(map[string]*Users),
	}

	lib.AddBook(BookList{
		Name:    "The Idiot",
		Author:  "Fyodor Dostoevsky",
		Subject: "Russian Literature",
		ISBN:    12345,
		ID:      "B001",
		Year:    1869,
		Status:  Available,
	})
	lib.AddBook(BookList{
		Name:    "The Stranger",
		Author:  "Albert Camus",
		Subject: "French Literature",
		ISBN:    67890,
		ID:      "B002",
		Year:    1942,
		Status:  Available,
	})
	AddUser(lib.Users, "Alice", "U001", "alice@example.com", "alice123")
	AddUser(lib.Users, "Bob", "U002", "bob@example.com", "bob456")

	fmt.Println("All Books:")
	for _, b := range lib.Books {
		fmt.Println(b.GetBookInfo())
	}
	fmt.Println("\nAlice checks out 'The Idiot':")
	lib.CheckoutBook("U001", "The Idiot")
	for _, b := range lib.FindBorrowedBooksByUser("U001") {
		fmt.Println(b.GetBookInfo())
	}

	fmt.Println("\nAlice returns 'The Idiot':")
	lib.ReturnBook("U001", "The Idiot")
	for _, b := range lib.FindBorrowedBooksByUser("U001") {
		fmt.Println(b.GetBookInfo())
	}
}

// Book
type BookList struct {
	Name    string
	Subject string
	ISBN    int
	Author  string
	ID      string
	Year    int
	Status  BookStatus
}

// BookStatus
type BookStatus int

const (
	Available BookStatus = iota
	CheckedOut
	Reserved
	Maintenance
)

func (bs BookStatus) String() string {
	switch bs {
	case Available:
		return "Available"
	case CheckedOut:
		return "Checked Out"
	case Reserved:
		return "Reserved"
	case Maintenance:
		return "Maintenance"
	default:
		return "Unknown"
	}
}

// type Collection[T Item] struct {

type Collection[T Item] struct {
	items map[string]T
}

type Item interface {
	GetBookID() string
	GetBookTitle() string
	GetBookInfo() string
}

func (b *BookList) GetBookID() string {
	return b.ID
}

func (b *BookList) GetBookTitle() string {
	return b.Name
}

func (b *BookList) GetBookInfo() string {
	return fmt.Sprintf("Title: %s | Author: %s | Subject: %s | ISBN: %d | Status: %s",
		b.Name, b.Author, b.Subject, b.ISBN, b.Status)
}

var (
	ErrBookNotFound = errors.New("Book Not Found !")
)

func (l *Collection[T]) GetByID(id string) (T, error) {
	item, exist := l.items[id] //exist boolen
	if !exist {
		var zero T // panic etefag nayofte
		return zero, ErrBookNotFound
	}
	return item, nil
}

// UsersList
type UsersList[T Members] struct {
	Members map[string]T
}
type Members interface {
	GetUserID() string
	GetUserName() string
	GetUserInfo() string
}

func (b *Users) GetUserID() string {
	return b.ID
}

func (b *Users) GetUserName() string {
	return b.Name
}

func (b *Users) GetUserInfo() string {
	return fmt.Sprintf("Name: %s | Email: %s | UserName: %s | CreatedAt: %d ",
		b.Name, b.Email, b.UserName, b.CreatedAt)
}

// libary
type Library struct {
	Name  string
	Books []BookList
	Users map[string]*Users
}

func (l *Library) AddBook(b BookList) {
	l.Books = append(l.Books, b)
}
func (c *Collection[T]) Add(item T) {
	c.items[item.GetBookID()] = item // eleman be map
}

// FindBookBy
func (l *Library) FindBookByName(name string) (int, error) {
	for i, book := range l.Books {
		if book.Name == name {
			return i, nil
		}
	}
	return -1, errors.New("Book not found !")
}

func (l *Library) FindBookByName2(name string) int {
	for i, book := range l.Books {
		if book.Name == name {
			return i
		}
	}
	return -1
}

func FindBookByAuthor(l Library, author string) (int, error) {
	for i, Book := range l.Books {
		if Book.Author == author {
			return i, nil
		}
	}
	return -1, errors.New("No books found for this author")
}
func FindBookByISBN(l Library, isbn int) (int, error) {
	for i, Book := range l.Books {
		if Book.ISBN == isbn {
			return i, nil
		}
	}
	return -1, errors.New("No books found for this ISBN")
}

func (l *Library) UpdateBook(oldName string, uName string, uSubject string, uISBN int, uAuthor string, uYear int) (string, error) {
	idx := l.FindBookByName2(oldName)
	if idx != -1 {
		l.Books[idx].Name = uName
		l.Books[idx].Subject = uSubject
		l.Books[idx].ISBN = uISBN
		l.Books[idx].Author = uAuthor
		l.Books[idx].Year = uYear
		return "Updated successfully", nil
	}
	return "", errors.New("Updating book details failed")
}

func (c *Collection[T]) Remove(id string) {
	delete(c.items, id)
}

// users

type Users struct {
	Name      string
	ID        string
	Email     string
	UserName  string
	CreatedAt time.Time
	Borrowed  []string
}

func AddUser(users map[string]*Users, name, id, email, username string) (string, error) {
	if _, exists := users[id]; exists {
		return "", errors.New("User with this ID already exists")
	}
	users[id] = &Users{
		ID:        id,
		Name:      name,
		Email:     email,
		UserName:  username,
		CreatedAt: time.Now(),
	}
	return "User added successfully", nil

}

func RemoveUser(users map[string]*Users, id string) (string, error) {
	if _, exists := users[id]; !exists {
		return "", errors.New("User not found")
	}

	delete(users, id)
	return "User removed successfully", nil
}

// check out

func (l *Library) CheckoutBook(userID, bookName string) (string, error) {
	user, ok := l.Users[userID]
	if !ok {
		return "", errors.New("User id not found")
	}

	if len(user.Borrowed) >= 5 {
		return "", errors.New("Can't Barrow more then five ")
	}

	idx := l.FindBookByName2(bookName)
	if idx == -1 {
		return "", errors.New("Book not found!")
	}

	book := &l.Books[idx]
	if book.Status != Available {

		return "", errors.New("Book is not available")
	}

	book.Status = CheckedOut
	user.Borrowed = append(user.Borrowed, book.ID)
	return "Book borrowed successfully", nil
}

// ReturnBook
func (l *Library) ReturnBook(userID, bookName string) (string, error) {
	user, ok := l.Users[userID]
	if !ok {
		return "", errors.New("User ID not found")
	}

	idx := l.FindBookByName2(bookName)
	if idx == -1 {
		return "", errors.New("Book not found")
	}

	book := &l.Books[idx]

	found := false
	for i, id := range user.Borrowed {
		if id == book.ID {
			user.Borrowed = append(user.Borrowed[:i], user.Borrowed[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return "", errors.New("This book was not borrowed by the user")
	}

	book.Status = Available
	return "Book returned successfully", nil
}

// FindBookBySubject

func (l *Library) FindBookBySubject(subject string) []BookList {
	var result []BookList

	for _, book := range l.Books {
		if strings.ToLower(book.Subject) == strings.ToLower(subject) {
			result = append(result, book)
		}

	}
	return result

}

// FindBorrowedBooksByUser
func (l *Library) FindBorrowedBooksByUser(userID string) []BookList {
	var result []BookList

	user, ok := l.Users[userID]
	if !ok {
		fmt.Println("User not found")
		return result
	}

	for _, borrowedID := range user.Borrowed {
		for _, book := range l.Books {
			if book.ID == borrowedID {
				result = append(result, book)
				break
			}
		}
	}
	return result
}

//Filter books by author

func (l *Library) SortBookByAuthor(author string) []BookList {
	var result []BookList
	for _, book := range l.Books {
		if strings.ToLower(book.Author) == strings.ToLower(author) {
			result = append(result, book)
		}
	}
	return result
}

// Filter books by Year
func (l *Library) SortBookByYear(year int) []BookList {
	var result []BookList

	for _, book := range l.Books {
		if book.Year == year {
			result = append(result, book)
		}
	}
	return result
}

//err
// Filter books by Books Status

func (l *Library) SortBookByStatus(status string) []BookList {
	var result []BookList

	for _, book := range l.Books {
		if book.Status.String() == status {
			result = append(result, book)
		}
	}
	return result
}

package main

import (
	"fmt"
	"time"
)

// کتاب
type BookList struct {
	Name    string
	Subject string
	ISBN    int
	Author  string
	ID      string
	Status  BookStatus
}

// وضعیت کتاب
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

// ارتباط
type Collection[T Item] struct {
	items map[string]T
}

type Item interface {
	GetID() string
	GetTitle() string
	GetInfo() string
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

// ارتباط یوزر
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

// جستجوی کتاب بر اساس عنوان، نویسنده، یا شابک
func FindBookByName(l Library, Name string) int {
	for i, Book := range l.Books {
		if Book.Name == Name {
			return i
		}
	}
	return -1
}

func FindBookByAuthor(l Library, author string) int {
	for i, Book := range l.Books {
		if Book.Author == author {
			return i
		}
	}
	return -1
}
func FindBookByISBN(l Library, isbn int) int {
	for i, Book := range l.Books {
		if Book.ISBN == isbn {
			return i
		}
	}
	return -1
}

func (l *Library) UpdateBook(oldName string, uName string, uSubject string, uISBN int, uAuthor string) bool {
	idx := FindBookByName(*l, oldName)
	if idx != -1 {
		l.Books[idx].Name = uName
		l.Books[idx].Subject = uSubject
		l.Books[idx].ISBN = uISBN
		l.Books[idx].Author = uAuthor
		return true
	}
	return false
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
}

func AddUser(users map[string]Users, name, id, email, username string) {
	users[id] = Users{
		ID:        id,
		Name:      name,
		Email:     email,
		UserName:  username,
		CreatedAt: time.Now(),
	}
}
func RemoveUser(users map[string]Users, id string) {
	delete(users, id)
}

// سیستم امانت

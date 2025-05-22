package main

import "fmt"

type Book struct {
	title       string
	author      string
	isAvailable bool
}

func (b Book) CheckOut(title string) bool {
	b.isAvailable = false
	return b.isAvailable
}

func (b Book) Return(title string) bool {
	b.isAvailable = true
	return b.isAvailable
}

type Library struct {
	Books ([]Book)
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l Library) ListBooks() {
	for _, book := range l.Books {
		availability := "Да"
		if !book.isAvailable {
			availability = "Нет"
		}
		fmt.Printf("Название: \"%s\", Автор: \"%s\", Доступна: %s\n", book.title, book.author, availability)
	}
}

func (l Library) CheckBookAvailability(title string) bool {
	for _, book := range l.Books {
		if book.title == title {
			return book.isAvailable
		}
	}
	return false
}

func main() {
	lib := Library{}
	lib.AddBook(Book{"Месть орков", "Ричард Кнаак", true})
	lib.AddBook(Book{"Последний страж", "Джэфф Грабб", true})

	lib.ListBooks()

	// Проверка доступности книги
	title := "Месть орков"
	if lib.CheckBookAvailability(title) {
		fmt.Printf("Книга \"%s\" доступна для займа.\n", title)
	} else {
		fmt.Printf("Книга \"%s\" недоступна.\n", title)
	}
}

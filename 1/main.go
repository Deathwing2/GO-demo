package main

import (
	"fmt"
)

// Структура книги
type Book struct {
	Title       string
	Author      string
	IsAvailable bool
}

// Метод для проверки доступности книги
func (l *Library) CheckBookAvailability(title string) bool {
	for _, book := range l.Books {
		if book.Title == title {
			return book.IsAvailable
		}
	}
	return false // Если книга не найдена, считаем, что она недоступна
}

// Структура библиотеки
type Library struct {
	Books []Book
}

// Метод для добавления книги в библиотеку
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// Метод для отображения всех книг
func (l *Library) ListBooks() {
	for _, book := range l.Books {
		availability := "Да"
		if !book.IsAvailable {
			availability = "Нет"
		}
		fmt.Printf("Название: \"%s\", Автор: \"%s\", Доступна: %s\n", book.Title, book.Author, availability)
	}
}

func main() {
	lib := Library{}
	lib.AddBook(Book{"1984", "Джордж Оруэлл", true})
	lib.AddBook(Book{"Мастер и Маргарита", "Михаил Булгаков", true})

	lib.ListBooks()

	// Проверка доступности книги
	title := "1984"
	if lib.CheckBookAvailability(title) {
		fmt.Printf("Книга \"%s\" доступна для займа.\n", title)
	} else {
		fmt.Printf("Книга \"%s\" недоступна.\n", title)
	}
}

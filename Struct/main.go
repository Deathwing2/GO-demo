package main

import "fmt"

type Book struct {
	title       string
	author      string
	isAvailable bool
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
		fmt.Println("Название:", book.title, "Автор:", book.author, "Доступна:", availability)
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

func (l Library) ChangeStatus(title string) bool {
	for i, book := range l.Books {
		if book.title == title {
			if book.isAvailable {
				l.Books[i].isAvailable = false
			} else {
				l.Books[i].isAvailable = true
			}
		}
	}
	return false
}

func main() {
	lib := Library{}
	lib.AddBook(Book{"Месть орков", "Ричард Кнаак", true})
	lib.AddBook(Book{"Последний страж", "Джэфф Грабб", true})

	// Изменяем  статус книги
	title := "Последний страж"
	lib.ChangeStatus(title)

	// Показываем всю библиотеку
	lib.ListBooks()

	// Проверка доступности книги
	title = "Последний страж"
	if lib.CheckBookAvailability(title) {
		fmt.Printf("Книга \"%s\" доступна для займа.\n", title)
	} else {
		fmt.Printf("Книга \"%s\" недоступна.\n", title)
	}

}

package internal

import "errors"

type Author struct {
	Name string `json:"name"`
}

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

var Books []Book
var BookID int

func InitBook() ([]Book, int) {
	books := []Book{
		{
			ID:    1,
			Title: "The Great Gatsby",
			Genre: "Fiction",
		},
		{
			ID:    4,
			Title: "To Kill a Mockingbird",
			Genre: "Fiction",
		},
		{
			ID:    7,
			Title: "1984",
			Genre: "Dystopian",
		},
	}
	return books, len(books)
}

func getBooks() ([]Book, error) {
	return Books, nil
}

func (t *Book) createBook() error {
	BookID++
	t.ID = BookID
	Books = append(Books, *t)
	return nil
}

func (t *Book) getBook() error {
	id := t.ID
	for _, book := range Books {
		if book.ID == id {
			t.Title = book.Title
			t.Genre = book.Genre
			return nil
		}
	}
	return errors.New("book not found")
}

func (t *Book) updateBook() error {
	id := t.ID
	for i, book := range Books {
		if book.ID == id {
			Books[i].Title = t.Title
			Books[i].Genre = t.Genre
			return nil
		}
	}
	return errors.New("book not found")
}

func (t *Book) deleteBook() error {
	id := t.ID
	indexToBeDeleted := -1
	for i, book := range Books {
		if book.ID == id {
			indexToBeDeleted = i
			break
		}
	}
	if indexToBeDeleted == -1 {
		return errors.New("book not found")
	}
	Books = append(Books[:indexToBeDeleted], Books[indexToBeDeleted+1:]...)
	return nil
}

package internal

import "errors"

type Author struct {
	Name string `json:"name"`
	Home string `json:"home"`
}

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorBooks struct {
	Author `json:"author"`
	BookDB []int `json:"books"`
}

var BookDB []Book
var BookID int
var CredentialsDB map[string]Credentials

func InitBook() ([]Book, int) {
	books := []Book{
		{
			ID:    1,
			Title: "The Great Gatsby",
			Genre: "Fiction",
		},
		{
			ID:    2,
			Title: "To Kill a Mockingbird",
			Genre: "Fiction",
		},
		{
			ID:    3,
			Title: "1984",
			Genre: "Dystopian",
		},
	}
	return books, len(books)
}

func getBooks() ([]Book, error) {
	return BookDB, nil
}

func (t *Book) createBook() error {
	BookID++
	t.ID = BookID
	BookDB = append(BookDB, *t)
	return nil
}

func (t *Book) getBook() error {
	id := t.ID
	for _, book := range BookDB {
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
	for i, book := range BookDB {
		if book.ID == id {
			BookDB[i].Title = t.Title
			BookDB[i].Genre = t.Genre
			return nil
		}
	}
	return errors.New("book not found")
}

func (t *Book) deleteBook() error {
	id := t.ID
	indexToBeDeleted := -1
	for i, book := range BookDB {
		if book.ID == id {
			indexToBeDeleted = i
			break
		}
	}
	if indexToBeDeleted == -1 {
		return errors.New("book not found")
	}
	BookDB = append(BookDB[:indexToBeDeleted], BookDB[indexToBeDeleted+1:]...)
	return nil
}

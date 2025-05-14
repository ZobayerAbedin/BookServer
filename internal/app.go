package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) handleRoutes() {
	a.Router.HandleFunc("/login", a.login).Methods("POST")
	a.Router.HandleFunc("/logout", a.logout).Methods("POST")

	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/books/{id}", a.readBook).Methods("GET")

	secured := a.Router.PathPrefix("/secured").Subrouter()
	secured.Use(JWTMiddleware)

	secured.HandleFunc("/books", a.createBook).Methods("POST")
	secured.HandleFunc("/books/{id:[0-9]+}", a.updateBook).Methods("PUT")
	secured.HandleFunc("/books/{id:[0-9]+}", a.deleteBook).Methods("DELETE")
}

func (a *App) Initialise(initialBooks []Book, id int) {
	BookDB = initialBooks
	BookID = id
	a.Router = mux.NewRouter().StrictSlash(true)
	a.handleRoutes()
}

func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, books)
}

func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = book.createBook()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusCreated, book)
}

func (a *App) readBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book := Book{ID: id}
	err = book.getBook()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, book)
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = id
	err = book.updateBook()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, book)
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book := Book{ID: id}
	err = book.deleteBook()
	if err != nil {
		sendError(w, http.StatusNotFound, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username != "user" || creds.Password != "pass" {
		sendError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := GenerateJWT(creds.Username)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(15 * time.Minute),
	})
	sendResponse(w, http.StatusOK, map[string]string{"token": token})
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now().Add(-time.Hour),
	})
	sendResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, message string) {
	errorMessage := map[string]string{"error": message}
	sendResponse(w, statusCode, errorMessage)
}

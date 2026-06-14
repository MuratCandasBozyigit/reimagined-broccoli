package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/models"
	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var newBooks models.Book // küçük harfe çekildi

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookObj := &models.Book{} // küçük harfe çekildi
	utils.ParseBody(r, bookObj)

	// FIX: Model dosyasındaki fonksiyonun adı CreateBook() olduğu için burası düzeltildi!
	b := bookObj.CreateBook()

	res, err := json.Marshal(b)
	if err != nil {
		fmt.Println("error while creating")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() // küçük harfe çekildi
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooksById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID) // küçük harfe çekildi
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publisher != "" {
		bookDetails.Author = updateBook.Author
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

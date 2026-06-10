package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/Model"
	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook Model.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookObj := &Model.Book{} // İsim çakışması olmasın diye bookObj yaptık
	utils.ParseBody(r, bookObj)
	b := bookObj.Create()
	res, err := json.Marshal(b) // err tanımlandı
	if err != nil {
		fmt.Println("error while creating")
	}
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := Model.GetAllBooks()                    // () eklendi, fonksiyon tetiklendi
	res, _ := json.Marshal(newBooks)                   // newBook değil, yukarıdaki newBooks verildi
	w.Header().Set("Content-Type", "application/json") // Yazım hatası düzeltildi
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
	bookDetails, _ := Model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)                // Tırnaklar kaldırıldı, değişkenin kendisi verildi
	w.Header().Set("Content-Type", "application/json") // Yazım hatası düzeltildi
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// main.go'da hata vermemesi için parametreleri eklendi
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}

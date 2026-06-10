package models

import (
	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"` // gorm:"" şeklindeki hatalı tırnak temizlendi
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	// CRITICAL FIX: config.GetDB().Close() satırı silindi! Burası kapanırsa API'ler çalışmaz.
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book) // Pointer (&) eklendi
	return book
}

// func GetId(Id int64)  {
// 	db := db.Where("ID=?", Id).Find(&id)
// 	return  &id,db
// }

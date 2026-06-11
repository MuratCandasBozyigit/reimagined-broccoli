package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// 1. Port bilgisini ekledik: @tcp(127.0.0.1:3306)
	// 2. Karakter setini düzelttik: charset=utf8mb4
	// 3. Şifre ve host arasına @ işareti koyduk
	dsn := "murat:9212arda@tcp(127.0.0.1:3306)/rahle?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Veritabanına bağlanılamadı kanka: %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

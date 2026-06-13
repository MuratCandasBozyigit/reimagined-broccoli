package main

import (
	"fmt"
	"net/http"
	"strings"
)

func replacer(input string) {
	input = strings.TrimSpace(input)
	inputTL := strings.ToLower(input)
	dash := strings.ReplaceAll(inputTL, " ", "-")
	replacer := strings.NewReplacer(
		"ç", "c",
		"ğ", "g",
		"ı", "i",
		"ö", "o",
		"ş", "s",
		"ü", "u",
	)
	done := replacer.Replace(dash)
	fmt.Println(done)
}

func slugApi(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	gelenMetin := queryParams.Get("input")

	if gelenMetin == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400 Hata Kodu
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Naber %s kanka!", gelenMetin)
}

func main() {
	http.HandleFunc("/api/v1/slugify", slugApi)
	http.ListenAndServe(":9010", nil)
}

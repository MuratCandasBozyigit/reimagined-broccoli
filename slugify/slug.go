package main

import (
	"fmt"
	"net/http"
	"strings"
)

func replacer(input string) string {
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
	return done
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
	res := replacer(gelenMetin)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"slug": "%s"}`, res)
}
func analyze() {

}

func main() {
	http.HandleFunc("/api/v1/slugify", slugApi)
	http.HandleFunc("/api/v1/analyze", analyze)
	http.ListenAndServe(":9010", nil)
}

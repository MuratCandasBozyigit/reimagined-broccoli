package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

var urlVeritabani = make(map[string]string)
var tiklamaSayilari = make(map[string]int)
var bodyguard sync.Mutex

func replacer(input string) string {
	inputTrim := strings.TrimSpace(input)
	inputLower := strings.ToLower(inputTrim)
	replacer := strings.NewReplacer("ç", "c",
		"ğ", "g",
		"ı", "i",
		"ö", "o",
		"ş", "s",
		"ü", "u",
		" ", "-")
	done := replacer.Replace(inputLower)
	return done
}
func randomCode() string {
	havuz := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	kod := ""
	maxL := 0
	for maxL < 6 {
		randomIndex := rand.Intn(len(havuz))
		kod += string(havuz[randomIndex])
		maxL++
	}
	return kod
}
func slug(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryZig := queryParams.Get("input")

	if queryZig == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400 Hata Kodu
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}

	res := replacer(queryZig)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"slug": "%s"}`, res)
}
func shortener(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryStr := queryParams.Get("input")
	if queryStr == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}

	kod := randomCode()
	bodyguard.Lock()
	defer bodyguard.Unlock()

	urlVeritabani[kod] = queryStr
	tiklamaSayilari[kod] = 0

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"shortener":"%s"}`, kod)
}
func redirect(w http.ResponseWriter, r *http.Request) {
	queryCode := r.URL.Query().Get("code")

	if queryCode == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}
	bodyguard.Lock()
	defer bodyguard.Unlock()

	orijinalUrl, varMi := urlVeritabani[queryCode]

	if !varMi { // yani varMi == false ise (kod bulunamadıysa)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound) // 404 Hatası
		fmt.Fprintf(w, `{"error": "Böyle bir kod yok kanka!"}`)
		return // Fonksiyonu burada kes, aşağıya devam etmesin
	}
	tiklamaSayilari[queryCode]++

	http.Redirect(w, r, orijinalUrl, http.StatusMovedPermanently)

}
func main() {
	http.HandleFunc("/api/slug", slug)
	http.HandleFunc("/api/shorten", shortener)
	http.HandleFunc("/api/redirect", redirect)
	http.ListenAndServe(":9010", nil)
}

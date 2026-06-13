package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AnalizSonucu struct {
	TotalChars int `json:"total_chars"`
	WordCount  int `json:"word_count"`
}

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
func analyze(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryStr := queryParams.Get("input")
	splittedStr := strings.Split(queryStr, "")
	uz := len(splittedStr)
	kelimeler := strings.Split(queryStr, " ")
	kelimeSayisi := len(kelimeler)
	rapor := AnalizSonucu{
		TotalChars: uz,
		WordCount:  kelimeSayisi,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rapor)

}

func base64Api(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryStr := queryParams.Get("input")
	action := queryParams.Get("action")
	if action == "encode" {
		encodedStr := base64.StdEncoding.EncodeToString([]byte(queryStr))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result": "%s"}`, encodedStr)
	} else if action == "decode" {
		decodedStr, err := base64.StdEncoding.DecodeString(queryStr)
		if err != nil {
			fmt.Println("naber")
		}
		temizMetin := string(decodedStr)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"result": "%s"}`, temizMetin)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest) // 400 Hata Kodu
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}
}

func main() {
	http.HandleFunc("/api/v1/slugify", slugApi)
	http.HandleFunc("/api/v1/analyze", analyze)
	http.HandleFunc("/api/v1/base64", base64Api)
	http.ListenAndServe(":9010", nil)
}

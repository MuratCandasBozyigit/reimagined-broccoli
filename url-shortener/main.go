package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

var urlVeritabani = make(map[string]string)
var tiklamaSayilari = make(map[string]int)
var bodyguard sync.Mutex
var istekSayaclari = make(map[string]int)

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

func limitKontrol(r *http.Request) bool {
	reqIp, _, _ := net.SplitHostPort(r.RemoteAddr)

	bodyguard.Lock()
	defer bodyguard.Unlock()

	istekSayaclari[reqIp]++

	if istekSayaclari[reqIp] > 3 {
		return false // Sınırı aştı, geçiş yasak!
	}
	return true // Sınırı aşmadı, geçebilir kanka
}

func slug(w http.ResponseWriter, r *http.Request) {
	// 1. Önce Limit Kontrolü (Kilit açılmadan önce!)
	if !limitKontrol(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, `{"error": "Çok hızlısın kanka, yavaşla!"}`)
		return
	}

	queryParams := r.URL.Query()
	queryZig := queryParams.Get("input")

	if queryZig == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}

	res := replacer(queryZig)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"slug": "%s"}`, res)
}

func shortener(w http.ResponseWriter, r *http.Request) {
	// 1. Önce Limit Kontrolü!
	if !limitKontrol(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, `{"error": "Çok hızlısın kanka, yavaşla!"}`)
		return
	}

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
	urlVeritabani[kod] = queryStr
	tiklamaSayilari[kod] = 0
	bodyguard.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"shortener":"%s"}`, kod)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	// 1. Önce Limit Kontrolü!
	if !limitKontrol(r) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, `{"error": "Çok hızlısın kanka, yavaşla!"}`)
		return
	}

	queryCode := r.URL.Query().Get("code")
	if queryCode == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "input parametresi boş olamaz kanka!"}`)
		return
	}

	bodyguard.Lock()
	orijinalUrl, varMi := urlVeritabani[queryCode]
	if varMi {
		tiklamaSayilari[queryCode]++
	}
	bodyguard.Unlock()

	if !varMi {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "Böyle bir kod yok kanka!"}`)
		return
	}

	http.Redirect(w, r, orijinalUrl, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/api/slug", slug)
	http.HandleFunc("/api/shorten", shortener)
	http.HandleFunc("/api/redirect", redirect)

	// Temizlik robotunu ListenAndServe'den önceye aldık, artık hayatta!
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			bodyguard.Lock()
			istekSayaclari = make(map[string]int)
			bodyguard.Unlock()
		}
	}()

	fmt.Println("🚀 Komuta Merkezi Port :9010 üzerinde ayağa kalktı kanka!")
	http.ListenAndServe(":9010", nil)
}

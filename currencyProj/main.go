// GO-2 [ORTA] - Eşzamanlı Döviz & Altın Fiyat Derleyici
//         - İsterler:
//             * Arka planda 3 farklı sahte finans kaynağını simüle et (time.Sleep ile).
//             * Ana endpoint çağrıldığında 3 fonksiyonu aynı anda (go keyword) ateşle.
//             * Context & Timeout: 1 saniyede dönmeyen kaynağı iptal et, eldekileri bas.
//         - Kural: sync.WaitGroup, chan (kanallar) ve context.WithTimeout zorunlu.

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type DovizRaporu struct {
	BankaAdi string
	Fiyat    float64
}

func merkezBankasi() DovizRaporu {
	time.Sleep(200 * time.Millisecond)
	return DovizRaporu{
		BankaAdi: "Merkez Bankası",
		Fiyat:    34.50,
	}

}

func bankA() DovizRaporu {
	time.Sleep((500 * time.Millisecond))
	return DovizRaporu{
		BankaAdi: "a",
		Fiyat:    36.50,
	}

}

func bankB() DovizRaporu {
	time.Sleep((1500 * time.Millisecond))
	return DovizRaporu{
		BankaAdi: "b",
		Fiyat:    34.40,
	}
}

func finansApi(w http.ResponseWriter, r *http.Request) {
	raporKanali := make(chan DovizRaporu, 3)
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	var gelenSonuclar []DovizRaporu
Havuz:
	for {
		select {
		// Kapı A: Kanala bir veri düştüğünde burası çalışır
		case rapor, acikMi := <-raporKanali:
			if !acikMi {
				// Arkadaki bekçi işini bitirip kanalı kapattıysa döngüyü kır!
				break Havuz
			}
			// Kanal hâlâ açıksa gelen raporu sepete ekle
			gelenSonuclar = append(gelenSonuclar, rapor)

		// Kapı B: 1 saniyelik süre dolduğu an bu kapı tetiklenir
		case <-ctx.Done():
			// Süre bitti kanka! Hantal bankayı (Banka B) beklemeden döngüyü kır!
			break Havuz
		}
	}

	// 3. Döngüden çıkınca eldeki sepeti JSON olarak internete fırlatma
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gelenSonuclar)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		raporKanali <- merkezBankasi()
	}()
	go func() {
		defer wg.Done()
		raporKanali <- bankA()
	}()
	go func() {
		defer wg.Done()
		raporKanali <- bankB()
	}()
	go func() {
		wg.Wait()
		close(raporKanali)
	}()
}

func main() {
	http.HandleFunc("/api/v2/finans", finansApi)
}

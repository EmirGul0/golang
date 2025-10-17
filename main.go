package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"fmt"
)

// TimeResponse API'den dönecek JSON yapısını tanımlar
type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
	Timestamp   int64  `json:"timestamp"`
}

// timeHandler, saat verisini JSON olarak döner
func timeHandler(w http.ResponseWriter, r *http.Request) {
	// CORS başlıkları ekle (Frontend'in farklı bir adresten erişmesi için)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()

	response := TimeResponse{
		CurrentTime: now.Format("15:04:05"), // Saat:Dakika:Saniye formatı
		Timestamp:   now.Unix(),             // Unix zaman damgası (JS için kolay)
	}

	json.NewEncoder(w).Encode(response)
}

// Varsayılan handler (hala "Merhaba" mesajı verebilir)
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Golang Saat API'si Yayında. Veri almak için /api/time adresini kullanın.")
}

func main() {
	// Endpoint atamaları
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/time", timeHandler)

	// Port ayarı (Render'dan gelen PORT değişkenini kullanır)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Sunucu %s portunda başlıyor...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
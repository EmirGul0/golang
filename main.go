package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// TimeResponse API'den dönecek JSON yapısı
type TimeResponse struct {
	CurrentTime string `json:"currentTime"`
	Timestamp   int64  `json:"timestamp"` // Milisaniye cinsinden
	Hours       int    `json:"hours"`
	Minutes     int    `json:"minutes"`
	Seconds     int    `json:"seconds"`
}

// timeHandler, saat verisini JSON olarak döner
func timeHandler(w http.ResponseWriter, r *http.Request) {
	// CORS başlıkları
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// İstanbul saat dilimini kullan
	location, _ := time.LoadLocation("Europe/Istanbul")
	now := time.Now().In(location)

	response := TimeResponse{
		CurrentTime: now.Format("15:04:05"),
		Timestamp:   now.UnixMilli(), // Milisaniye cinsinden
		Hours:       now.Hour(),
		Minutes:     now.Minute(),
		Seconds:     now.Second(),
	}

	json.NewEncoder(w).Encode(response)
}

// Varsayılan handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <title>Golang Saat API</title>
    <style>
        body { font-family: 'Segoe UI', sans-serif; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; text-align: center; padding: 50px; }
        .info { background: rgba(255,255,255,0.1); padding: 30px; border-radius: 15px; max-width: 600px; margin: 0 auto; backdrop-filter: blur(10px); }
        a { color: #ffd700; text-decoration: none; font-weight: bold; }
    </style>
</head>
<body>
    <div class="info">
        <h1>🕐 Golang Saat API'si Aktif</h1>
        <p>Gerçek zamanlı saat verisi almak için:</p>
        <p><a href="/api/time">/api/time</a> endpoint'ini kullanın</p>
    </div>
</body>
</html>`
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/time", timeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Sunucu %s portunda başlıyor...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
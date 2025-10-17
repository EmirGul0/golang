package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// Ana sunucu fonksiyonu. Tüm gelen isteklere bu fonksiyon yanıt verir.
func handler(w http.ResponseWriter, r *http.Request) {
    // Tarayıcıya veya API isteğine gönderilecek yanıt.
    // Erişim yolu (r.URL.Path) genellikle "/" olacaktır.
    fmt.Fprintf(w, "Golang Web Servisimden Merhaba! Erişim Yolu: %s\n", r.URL.Path)
}

func main() {
    // 1. Gelen tüm isteklere (/) handler fonksiyonunu ata
    http.HandleFunc("/", handler)

    // 2. Sunucunun dinleyeceği portu belirle
    // Cloud/Render gibi platformlar portu bir Ortam Değişkeni (PORT) olarak ayarlar.
    port := os.Getenv("PORT")
    if port == "" {
        // Eğer PORT ortam değişkeni yoksa (yerelde test ederken) varsayılan olarak 8080 kullan
        port = "8080" 
    }

    // 3. Sunucuyu başlat
    log.Printf("Sunucu %s portunda başlıyor...\n", port)
    
    // http.ListenAndServe ile sunucuyu dinlemeye başla
    err := http.ListenAndServe(":"+port, nil)
    
    // Eğer bir hata olursa (örneğin port meşgulse) programı durdur
    if err != nil {
        log.Fatal("Sunucuyu başlatırken hata oluştu: ", err)
    }
}
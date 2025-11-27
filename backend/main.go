package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    	"depixen/backend/database" 
	"depixen/backend/repository" 
	"depixen/backend/service" 
	"depixen/backend/endpoint"
	"depixen/backend/transport"
    
    // CORS desteği için (Eğer tarayıcıdan test edecekseniz şarttır)
    "github.com/rs/cors"
)

func main() {
    // 1. Veritabanı Bağlantısını Başlat
    db := database.InitDB() 
    repo := repository.NewGormCaseStudyRepository(db)

    // 2. Cloudinary Uploader'ı Başlat (Hata Kontrolü Önemli!)
    uploader, err := service.NewUploader()
    if err != nil {
        log.Fatalf("Cloudinary başlatma hatası: %v. Lütfen ortam değişkenlerini kontrol edin.", err)
    }

    // 3. Service (İş Mantığı) Katmanını Başlat
    svc := service.NewCaseStudyService(repo, uploader) // Uploader enjekte edildi

    // 4. Go-Kit Endpoint katmanını başlat
    endpoints := endpoint.New(svc)

    // 5. Gorilla Mux Router'ı başlat
	router := mux.NewRouter()

    // 6. Go-Kit Transport katmanını Router'a bağla
    handler := transport.NewHTTPHandler(endpoints, router)
    
    // 7. CORS Ayarları (3000 portu için)
    // Frontend (React) 3000'den çalıştığı için CORS zorunludur.
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // React projesinin çalıştığı port
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    // CORS handler'ı ile Router'ı sarmala
    corsHandler := c.Handler(handler)


    // Sağlık Kontrolü Rotası
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API is fully functional and ready for transactions!")
	}).Methods("GET")


	port := ":8080"
	fmt.Printf("Backend API %s portunda dinleniyor...\n", port)

	// 8. HTTP Sunucusunu başlat (CORS Handler kullanıldı)
	err = http.ListenAndServe(port, corsHandler) // Handler yerine corsHandler kullanıldı
	
	if err != nil {
		log.Fatal("Server başlatılırken hata oluştu: ", err)
	}
}
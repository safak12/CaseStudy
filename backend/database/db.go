package database

import (
	"fmt"
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type CaseStudy struct {
	ID          uint `gorm:"primarykey"`
	Title       string
	Description string
	ImageURI    string 
	CreatedDate time.Time
}

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=depixen-pass dbname=postgres port=5439 sslmode=disable TimeZone=Europe/Istanbul"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	err = db.Table("tb_casestudy").AutoMigrate(&CaseStudy{})
	if err != nil {
		log.Fatalf("Tablo oluşturulamadı (AutoMigrate hatası): %v", err)
	}

	fmt.Println("Veritabanı bağlantısı başarılı ve tb_casestudy tablosu hazır.")
	DB = db
	return db
}
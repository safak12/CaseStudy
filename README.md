\# Case Study: Depixen Go-Kit \& React Uygulaması



Bu proje, bir tam yığın (Full-Stack)  çalışmasıdır. Proje; Go-Kit tabanlı bir mikroservis mimarisi kullanarak bir kart yönetim API'si (CRUD işlemleri) oluşturmayı ve bu API'yi tüketen bir React arayüzü geliştirmeyi amaçlamaktadır.



\## 🚀 Proje Mimarisi ve Teknolojileri



| Katman | Teknoloji | Amaç |

| :--- | :--- | :--- |

| \*\*Backend (API)\*\* | Go, Go-Kit, GORM, Gorilla Mux | API Endpointleri, Servis Katmanı, İş Mantığı, Routing. |

| \*\*Veritabanı\*\* | PostgreSQL (Docker ile) | Güvenilir ve kalıcı veri depolama. |

| \*\*Görsel Depolama\*\* | Cloudinary | Base64 görselleri işleme, bulutta güvenli depolama ve CDN. |

| \*\*Frontend (UI)\*\* | React, Axios/Fetch API | Kullanıcı arayüzü ve Backend API'yi tüketen istemci. |



\## 📦 Teslimat Paketi İçeriği



Teslimat paketininizin içinde (bu ZIP/RAR dosyasının içinde) şunlar bulunmaktadır:



1\.  `backend/`: Go-Kit tabanlı Go projesi ve API kaynak kodları.

2\.  `frontend/`: React tabanlı arayüz projesi kaynak kodları (`node\_modules` \*\*silinmiştir\*\*).

3\.  `postgres\_data\_for\_delivery.zip`: Veritabanı verilerinin (kaydedilen kartlar) bulunduğu sıkıştırılmış PostgreSQL volume içeriği.



\## ▶️ Projeyi Çalıştırma Talimatları



Projeyi yerel ortamınızda başlatmak için sırasıyla aşağıdaki adımları izleyin.



\### Ön Gereksinimler



\* Go (v1.18+)

\* Node.js ve npm

\* Docker Desktop

\* Cloudinary Hesabı (API Key, Secret ve Cloud Name gereklidir)



\### Adım 1: Veritabanını (PostgreSQL) Başlatma



Veritabanını başlatmak ve kaydedilmiş test verilerini yüklemek için:



1\.  Terminalinizde Docker'ın çalıştığından emin olun.

2\.  Ana dizinde bulunan `postgres\_data\_for\_delivery.zip` dosyasını açın.

3\.  Zip dosyasının içindeki veri klasörünü (muhtemelen `data` veya benzeri bir isimle) kullanarak bir Docker Volume oluşturup konteyneri başlatın. Veya daha basitçe, aşağıdaki komutu çalıştırarak konteyneri başlatın:

&nbsp;   ```bash

&nbsp;   # Eğer konteyner daha önce oluşturulduysa:

&nbsp;   docker start case-study-postgres



&nbsp;   # Eğer konteyner silindiyse, yeniden oluşturulmalıdır (İlk Kurulum komutu):

&nbsp;   # docker run --name case-study-postgres -e POSTGRES\_USER=gorm -e POSTGRES\_PASSWORD=gorm -e POSTGRES\_DB=casestudy -v depixen-volume:/var/lib/postgresql/data -p 5439:5432 -d postgres:latest

&nbsp;   ```

4\.  Bağlantı başarılı olduğunda, Go Backend'e geçin.



\### Adım 2: Backend API'yi Başlatma (Cloudinary Ayarları)



1\.  \*\*Backend\*\* klasörüne gidin: `cd backend`

2\.  \*\*Cloudinary Ortam Değişkenlerini Tanımlayın:\*\*

&nbsp;   \* \*\*UYARI:\*\* Aşağıdaki komutları kendi Cloudinary bilgilerinizle doldurun!



&nbsp;   ```powershell

&nbsp;   # PowerShell veya CMD ortamında

&nbsp;   $env:CLOUDINARY\_CLOUD\_NAME = "\[CLOUD\_NAME\_BİLGİNİZ]"

&nbsp;   $env:CLOUDINARY\_API\_KEY = "\[API\_KEY\_BİLGİNİZ]"

&nbsp;   $env:CLOUDINARY\_API\_SECRET = "\[API\_SECRET\_BİLGİNİZ]"

&nbsp;   ```

3\.  Backend'i çalıştırın:

&nbsp;   ```bash

&nbsp;   go run main.go

&nbsp;   ```

&nbsp;   \*\*Beklenti:\*\* "Backend API :8080 portunda dinleniyor..." çıktısı alınmalıdır.



\### Adım 3: Frontend'i Başlatma (React)



Yeni bir terminal penceresi açın.



1\.  \*\*Frontend\*\* klasörüne gidin: `cd frontend`

2\.  \*\*Bağımlılıkları Yükleyin:\*\* (`node\_modules` klasörü silindiği için bu zorunludur)

&nbsp;   ```bash

&nbsp;   npm install

&nbsp;   ```

3\.  Frontend'i çalıştırın:

&nbsp;   ```bash

&nbsp;   npm start

&nbsp;   ```

&nbsp;   \*\*Beklenti:\*\* Tarayıcı otomatik olarak `http://localhost:3000` adresinde açılacaktır.



\## ✅ Test Edilen Temel Özellikler



\* \*\*Veri Kaydı (POST /cards):\*\* Form aracılığıyla yeni bir kart oluşturma.

\* \*\*Veri Listeleme (GET /cards):\*\* Tüm kartları PostgreSQL'den çekme ve arayüzde gösterme.

\* \*\*Görsel İşleme:\*\* React tarafından Base64'e çevrilen görselin, Go Backend tarafından Cloudinary'ye yüklenmesi ve URL'in veritabanına kaydedilmesi.

\* \*\*Mikroservis Mimarisi:\*\* Go-Kit'in \*\*Service\*\*, \*\*Endpoint\*\* ve \*\*Transport\*\* katmanlarının başarılı bir şekilde ayrılması.



---



\## 💡 Değerlendirme ve Öğrenim Çıkarımları



Bu çalışması sırasında kazanılan başlıca deneyim ve çıkarımlar aşağıdadır:



1\.  \*\*Go-Kit Mimarisi Anlayışı:\*\* Monolitik bir yapı yerine, Go-Kit'in Transport, Endpoint ve Service katmanlarını kullanarak daha ölçeklenebilir ve test edilebilir bir API yapısını uygulama pratiği kazanılmıştır. Özellikle \*middleware\* ve \*decoder/encoder\* yapılarının çalışma mantığı  öğrenilmiştir.

2\.  \*\*GORM ve PostgreSQL Entegrasyonu:\*\* Go'da ORM (GORM) kullanımıyla PostgreSQL'e veri yazma, okuma ve `CreatedDate` gibi alanların otomatik yönetimi teyit edilmiştir.

3\.  \*\*Harici Servis Entegrasyonu (Cloudinary):\*\* Bir harici API'nin Go dilinde, Base64 formatıyla başlayıp HTTP POST isteğiyle başarılı bir şekilde entegrasyonu sağlanmıştır.

4\.  \*\*Hata Ayıklama (Debug) Pratiği:\*\* Özellikle Cloudinary ortam değişkenlerinin sıfırlanması, React/Go-Kit arasındaki JSON büyük/küçük harf uyuşmazlığı gibi zorlu hataların Network sekmesi ve konsol çıktıları ile çözümlenmesi konusunda güçlü bir pratik kazanılmıştır.




# Notification System

Bu repo, Go ile gelistirilmis basit bir bildirim sistemini icerir. Proje, **hexagonal (ports and adapters)** yaklasimini temel alir ve yeni kanallarin veya veri kaynaklarinin kolayca eklenebilmesi amaciyla moduler bir yapi sunar.

## Hexagonal Mimarinin Avantajlari

- **Bagimsiz Katmanlar**: Is kurallari ile dis dunyayi ayristiran port ve adapter yapisi sayesinde, HTTP servisi, veritabani ya da baska bir servis degistiginde uygulama cekirdegi minimum etkilenir.
- **Test Edilebilirlik**: Uygulama cekirdegi (use-case'ler ve servisler) arayuzler uzerinden calistigi icin, adapter'lar yerine sahte implementasyonlar kullanarak birim testler yazmak kolaydir.
- **Genisletilebilirlik**: Yeni bir bildirim kanali eklemek veya farkli bir veri tabani kullanmak icin mevcut portlara yeni adapter'lar yazmak yeterlidir. Mevcut is kurallari ayni sekilde calisir.

## Kod Dizini

```
├── cmd/api          # Uygulamayi baslatan main paketi
├── internal
│   ├── adapters     # HTTP, mail, sms, push ve repository gibi dis bagimliliklar
│   ├── application
│   │   ├── core     # Use-case'ler ve domain modelleri
│   │   └── notification  # Bildirim servisi ve ilintili kodlar
│   └── ports        # Uygulamanin haberlesecegi arayuzler
└── pkg              # Konfigurasyon ve genel araclar
```

### `internal/application/notification` Neden Var?

Bu klasor, uygulamanin bildirim gonderme ile ilgili tum is kurallarini barindirir. Yapisi su sekildedir:

- **service.go** – `NotificationService`, bildirimleri bir veya birden fazla kanaldan (mail, sms, push) gonderir. Kanallar paralel calistirilir.
- **dispatcher/** – Farkli olay tiplerine ait handler'larin kaydedildigi ve calistirildigi mekanizma. `RegisterHandler` fonksiyonu ile bir olay adina karsilik gelen handler eklenir.
- **notifications/** – Her olay icin olusturulan bildirim turleri. Ornegin `UserFollowedNotification` yalnizca e-posta kanali uzerinden mesaj uretir.
- **messages/** – Mail, SMS ve push icin kullanilan basit mesaj modelleri.

Bu ayrim sayesinde bildirim akisi ile ilgili tum kodlar tek yerde toplanir ve port/adapter katmanlarindan bagimsiz olarak calisir.

### Is Kurallari

`internal/application/core` klasoru altinda uygulamanin temel use-case'leri bulunur. `api` paketindeki `Application` yapisi, dis dunyadan (HTTP vb.) gelen istekleri alir ve gerekli servisleri kullanarak is akisini yurutur. `TriggerNotification` use-case'i gelen olay tipini `dispatcher` uzerinden uygun handler'a yonlendirir ve uretilen bildirimleri `NotificationService` araciligiyla gonderir.

Domain modelleri (`domain` paketi) yalnizca gerekli alanlari barindirir ve dis katmanlara bagimlilik icermez.

## Calistirma

1. `.env.example` dosyasini kopyalayip gerekli degiskenleri doldurun.
2. Gerekli bagimliliklari almak ve uygulamayi baslatmak icin:
   ```bash
   go run ./cmd/api
   ```
3. HTTP servisi varsayilan olarak Fiber uzerinde belirttiginiz portta calisir ve `/api/notifications/trigger` rotasi ile bildirim tetikler.
4. Kullanici islemleri icin `/api/users` (POST), `/api/users/:id` (PUT, DELETE) rotalari kullanilabilir.

## Docker ile Çalıştırma

Projenin tum bagimliliklarini Docker uzerinden calistirmak icin asagidaki komutu kullanabilirsiniz. Bu komut MySQL, MailPit ve uygulamayi bir arada baslatir:

```bash
docker compose up --build
```

Uygulama `localhost:4000` uzerinden erisilebilir. MailPit arayuzune `localhost:8025` adresinden ulasabilirsiniz.

## Sonuc

Bu proje, hexagonal mimariyi kullanarak sade ama genisletilebilir bir bildirim altyapisi sunar. Port ve adapter yapisi sayesinde kodun test edilebilirligi artar, dis bagimliliklarin degismesi kolaylasir ve gelisime acik bir temel olusur.

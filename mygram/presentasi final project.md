# Tahapan pembuatan final project:

1. Membuat folder mygram
2. go mod init mygram
3. menginstall package yg dibutuhkan

   - github.com/gin-contrib/location v0.0.2
   - gorm.io/driver/postgres v1.4.4
   - gorm.io/gorm v1.24.0
   - github.com/gin-gonic/gin v1.8.1
   - github.com/swaggo/gin-swagger v1.5.3
   - github.com/swaggo/http-swagger v1.3.3
   - github.com/swaggo/swag v1.8.7
   - github.com/swaggo/files
   - github.com/dgrijalva/jwt-go v3.2.0+incompatible
   - github.com/dranikpg/dto-mapper v0.1.1
   - github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
   - golang.org/x/crypto v0.1.0

4. setup database mygram_db
5. membuat struktur folder
   - controllers -> berisi controllers yang dibutuhkan
   - database -> berisi fungsi setup koneksi ke database
   - docs -> berisi hasil generate swagger API Documentation
   - helpers -> berisi fungsi-fungsi pendukung
   - middlewares -> berisi fungsi middlewares yang digunakan
   - models -> berisi model yang terhubung dengan tiap table database
   - routers -> berisi fungsi setup routers, dan route endpoint yang dihubungkan dengan controller masing-masing
   - structs -> berisi struct yang dibutuhkan pada program
   - uploaded -> folder static pada server, berisi foto-foto yang akan di upload
6. membuat file main.go
7. membuat file untuk konfigurasi koneksi ke database
8. mengkoneksikan fungsi main dengan fungsi koneksi database
9. membuat file controllers
   - user.go
   - photo.go
   - comment.go
   - socialMedia.go
10. membuat file setup router dan setiap router yang terhubung dengan controllers
11. mengkoneksikan fungsi main dengan setup router, mengirim variable koneksi database ke router.
12. membuat setup router untuk swagger
13. membuat fungsi untuk menjalankan server beserta PORTnya.
14. men-generate swagger docs
15. menjalankan main.go sekaligus melakukan migration
16. membuat dokumentasi postman, dan panduan.
17. build program mygram.exe

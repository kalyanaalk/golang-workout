- [Go Modules](#go-modules)
- [GOPATH dan Workspace](#gopath-dan-workspace)
- [Instalasi Editor](#instalasi-editor)
- [Command](#command)
  - [go mod init](#go-mod-init)
  - [go run](#go-run)
  - [go test](#go-test)
  - [go build](#go-build)
  - [go get](#go-get)
  - [go mod download](#go-mod-download)
  - [go mod tidy](#go-mod-tidy)
  - [go mod vendor](#go-mod-vendor)
- [Program Pertama: Hello World](#program-pertama-hello-world)
  - [Penggunaan Keyword Package](#penggunaan-keyword-package)
  - [Penggunaan Keyword import](#penggunaan-keyword-import)
  - [Penggunaan Fungsi main()](#penggunaan-fungsi-main)
  - [Penggunaan Fungsi fmt.Println()](#penggunaan-fungsi-fmtprintln)
- [Komentar](#komentar)
- [Variabel](#variabel)
- [Tipe Data](#tipe-data)
- [Konstanta](#konstanta)
- [Operator](#operator)
- [Seleksi Kondisi](#seleksi-kondisi)
- [Perulangan](#perulangan)
- [Array](#array)
- [Slice](#slice)
- [Map](#map)
- [Fungsi](#fungsi)
- [Fungsi Multiple Return](#fungsi-multiple-return)
- [Fungsi Variadic](#fungsi-variadic)
- [Fungsi Closure](#fungsi-closure)
- [Fungsi Sebagai Parameter](#fungsi-sebagai-parameter)
- [Pointer](#pointer)
- [Struct](#struct)
- [Method](#method)
- [Properti Public dan Private](#properti-public-dan-private)
- [Interface](#interface)
- [Any / Interface{} / Interface Kosong](#any--interface--interface-kosong)
- [Reflect](#reflect)
- [Goroutine](#goroutine)
- [Channel](#channel)
- [Buffered Channel](#buffered-channel)
- [Channel - Select](#channel---select)
- [Channel - Range & Close](#channel---range--close)
- [Channel - Timeout](#channel---timeout)
- [Defer & Exit](#defer--exit)
- [Error, Panic, & Recover](#error-panic--recover)
- [Layout Format String](#layout-format-string)
- [Random](#random)
- [Time, Parsing Time, & Format Time](#time-parsing-time--format-time)
- [Timer, Ticker, & Scheduler](#timer-ticker--scheduler)
- [Time Duration](#time-duration)
- [Konversi Antar Tipe Data](#konversi-antar-tipe-data)
- [Fungsi String](#fungsi-string)
- [Regexp](#regexp)
- [Encode - Decode Base64](#encode---decode-base64)
- [Hash Sha1](#hash-sha1)
- [Arguments & Flag](#arguments--flag)
- [Exec](#exec)
- [File](#file)
- [Web Server](#web-server)
- [URL Parsing](#url-parsing)
- [JSON Data](#json-data)
- [Web Service API Server](#web-service-api-server)
- [Simple Client HTTP Request](#simple-client-http-request)
- [SQL](#sql)
- [NoSQL MongoDB](#nosql-mongodb)
- [Unit Test](#unit-test)
- [sync.WaitGroup](#syncwaitgroup)
- [sync.Mutex](#syncmutex)
- [Go Vendoring](#go-vendoring)
- [Concurrency Pattern: Pipeline](#concurrency-pattern-pipeline)
- [Concurrency Pattern: Simplified Fan-in Fan-out Pipeline](#concurrency-pattern-simplified-fan-in-fan-out-pipeline)
- [Concurrency Pattern: Context Cancellation Pipeline](#concurrency-pattern-context-cancellation-pipeline)
- [Go Generics](#go-generics)

## Go Modules

Go modules merupakan tools untuk manajemen dependensi resmi milik Go. Modules digunakan untuk menginisialisasi sebuah project, sekaligus melakukan manajemen terhadap 3rd party atau library atau dependency yang digunakan dalam project.

Gunakan command ini untuk menginisialisasi project baru di folder project.

```shell
go mod init <nama-project>
go mod init project-pertama
```

Eksekusi command ini menghasilkan satu buah file baru bernama `go.mod`. File ini digunakan oleh Go toolchain untuk menandai bahwa folder di mana file tersebut berada adalah folder project. 

## GOPATH dan Workspace

GOPATH adalah variabel yang digunakan oleh Go sebagai rujukan lokasi di mana semua folder project disimpan (kecuali untuk yg diinisialisasi menggunakan Go Modules). GOPATH berisikan 3 buah sub-folder: `src`, `bin`, dan `pkg`.

Project di Go bisa ditempatkan dalam `$GOPATH/src`. Sebagai contoh anda ingin membuat project dengan nama `belajar`, maka harus dibuatkan sebuah folder dengan nama 
`belajar`, ditempatkan dalam `src` ( `$GOPATH/src/belajar` ).

Lokasi folder yang dapat dijadikan sebagai workspace adalah bebas selain di GOROOT dan harus didaftarkan dalam path variable dengan nama `GOPATH`.

Setelah `GOPATH` berhasil dikenali, perlu disiapkan 3 buah sub folder di dalamnya, dengan kriteria sebagai berikut:
- Folder `src`, adalah path di mana project Go disimpan 
- Folder `pkg`, berisi file hasil kompilasi
- Folder `bin`, berisi file executable hasil build

## Instalasi Editor

Dapat menggunakan editor Visual Studio Code (VSCode) dengan menginstal extensi Go. Banyak benefit yang didapat dari ekstensi ini, beberapa di antaranya adalah integrasi dengan kompiler Go, auto lint on save, testing with coverage, fasilitas debugging with breakpoints, dan lainnya. 

Editorconfig juga dapat diinstal di VSCode untuk membantu coding style menjadi konsisten untuk dibaca oleh banyak developer, dan juga ketika dimuat pada berbagai macam IDE. Editorconfig pada sebuah proyek (biasanya berada di root direktori proyek tersebut) berupa konfigurasi format file 
.editorconfig  yang berisi definisi style penulisan yang menyesuaikan dengan standar penulisan masing-masing bahasa pemrograman. Misalnya untuk style guide GO kita bisa mulai dengan menggunakan konfigurasi sederhana sebagai berikut:

```
root = true
[*]
insert_final_newline = true
charset = utf-8
trim_trailing_whitespace = true
indent_style = space
indent_size = 2
[{Makefile,go.mod,go.sum,*.go}]
indent_style = tab
indent_size = 8
```

## Command

### go mod init

Untuk inisialisasi project pada Go yang menggunakan Go Modules. Untuk nama project bisa menggunakan apapun, tapi umumnya disamakan dengan nama direktori/folder. Nama project ini penting karena nantinya berpengaruh pada import path sub packages yang ada dalam project tersebut.

### go run

Untuk eksekusi file program, yaitu file yang berekstensi `.go`. Cara penggunaannya dengan menuliskan command tersebut diikuti argumen nama file di folder berisi program tersebut. Berikut adalah contoh penerapan `go run`  untuk eksekusi file main.go yang path project-nya sudah diinisialisasi menggunakan `go mod init`.

```shell
go run <nama-file>
go run main.go
```

### go test

Go menyediakan package testing , berguna untuk keperluan pembuatan file test. Pada penerapannya, ada aturan yang wajib diikuti yaitu nama file test harus berakhiran `_test.go`.Berikut adalah contoh penggunaan command.

```shell
go test <nama-file-testing>
go test main_test.go
```

### go build

Command ini digunakan untuk mengkompilasi file program.

Perbedaannya dengan `go run` adalah, file hasil kompilasi menggunakan `go run` disimpan pada folder temporary untuk selanjutnya langsung dieksekusi, sedangkan `go build` menghasilkan file executable atau binary pada folder yang sedang aktif. 

```shell
go build -o <nama-executable>
go build -o program.exe
```

### go get

Digunakan untuk men-download package atau dependency.

Contohnya, jika ingin men-download package Kafka driver untuk Go pada project `project-pertama`, maka command-nya kurang lebih seperti berikut:

```
cd project-pertama
go get github.com/segmentio/kafka-g
```

Lalu jika ingin mengunduh package/dependency versi terbaru, gunakan flag -u pada command `go get`, contohnya:
 
```
go get -u github.com/segmentio/kafka-go
```

Command ini harus dijalankan dalam folder project. Jika dijalankan di luar path project maka dependency yang terunduh akan ter-link dengan GOPATH, bukan dengan project

### go mod download

Command `go mod download` digunakan untuk men-download dependency.

### go mod tidy

Command `go mod tidy`  digunakan untuk memvalidasi dependency sekaligus men-download-nya jika memang belum ter-download.

### go mod vendor

Command ini digunakan untuk vendoring.

## Program Pertama: Hello World

1. Buat direktori bernama hello-world lalu inisialisasi project.

```shell
 mkdir hello-world
 cd hello-world
 go mod init hello-world
 ```

2. Buka project di VSCode.

3. Buat file program berekstensi `.go`, misal bernama `main.go`.

4. Salin program di bawah ke `main.go`.

```go
 package main
 import "fmt"
 func main() {
    fmt.Println("Hello world")
 }
```

5. Jalankan dengan command `go run main.go`.

### Penjelasan Program Hello World

#### Penggunaan Keyword Package

Setiap file program harus memiliki package. Setiap project harus ada minimal satu file dengan nama package `main`. File yang ber-package main, akan dieksekusi pertama kali ketika program di jalankan. 

Cara menentukan package adalah dengan menggunakan keyword `package`, berikut adalah contoh penulisannya.

```go
 package <nama-package>
 package main
```

#### Penggunaan Keyword import

Keyword `import` digunakan untuk meng-import atau memasukan package lain ke dalam file program, agar isi dari package yang di-import bisa dimanfaatkan. 

Package `fmt` merupakan salah satu package bawaan yang disediakan oleh Go, isinya banyak fungsi untuk keperluan I/O yang berhubungan dengan text.

```go
 import "<nama-package>"
 import "fmt"
```

#### Penggunaan Fungsi main()

Dalam sebuah proyek harus ada file program yang di dalamnya berisi sebuah fungsi bernama `main()`. Fungsi tersebut harus berada di file yang package-nya bernama main.

Fungsi `main()` adalah yang dipanggil pertama kali pada saat eksekusi program. Contoh penulisan fungsi main:

```go
 func main() {
    // program
 }
```

#### Penggunaan Fungsi fmt.Println()

Fungsi `fmt.Println()` digunakan untuk memunculkan output teks. Di project hello-world, fungsi ini memunculkan tulisan Hello world.
 
Skema penulisan keyword `fmt.Println()` bisa dilihat pada contoh berikut:

```go
 fmt.Println("<isi-pesan>")
 fmt.Println("Hello world")
```

Fungsi fmt.Println()  berada dalam package `fmt`, maka untuk menggunakannya perlu dilakukan import package tersebut.

Fungsi fmt.Println()  dapat menampung parameter yang tidak terbatas jumlahnya. Semua data parameter akan dimunculkan dengan pemisah tanda spasi. Misal, `fmt.Println("Hello", "world!", "how", "are", "you")` akan menghasilkan output "Hello world! how are you".

## Komentar

Komentar adalah bagian program yang diabaikan saat kompilasi, biasanya dimanfaatkan pengembang untuk menyisipkan catatan atau menulis penjelasan blok kode, atau menonaktifkan potongan kode tanpa menghapusnya.

Pertama adalah komentar inline yang menggunakan `//`

```go
 package main
 import "fmt"
 func main() {
    // komentar kode
    // menampilkan pesan hello world
    fmt.Println("hello world")
    
    // fmt.Println("baris ini tidak akan dieksekusi")
 }
 ```

 Kedua adalah komentar multiline yang menggunakan `/* */`

```go
 package main
 import "fmt"
 func main() {
    /*
    komentar kode
    menampilkan pesan hello world
    */
    fmt.Println("hello world")
 }
 ```

## Variabel

### Deklarasi dengan Tipe Data

Istilah dari metode deklarasi variabel ini adalah manifest typing. Di kode di bawah, keyword `var` digunakan untuk deklarasi variabel, contohnya bisa dilihat pada `firstName`  dan `lastName`.
 
Nilai variabel `firstName` diisi langsung ketika deklarasi, berbeda dibanding `lastName` yang nilainya diisi setelah baris kode deklarasi, hal seperti ini diperbolehkan di Go.

```go
 package main

 import "fmt"
 func main() {
   var firstName string = "john"

   var lastName string
   lastName = "wick"

   fmt.Printf("halo %s %s!\n", firstName, lastName)
 }
```

### Deklarasi dengan Keyword var

Pada kode di atas bisa dilihat bagaimana sebuah variabel dideklarasikan dan diisi nilainya. Keyword `var` digunakan untuk membuat variabel baru.

Skema penggunaan keyword var:

``` go
 var <nama-variabel> <tipe-data>
 var <nama-variabel> <tipe-data> = <nilai>
// Contoh:
 var lastName string
 var firstName string = "john"
```

#### Penggunaan Fungsi fmt.Printf()

Fungsi ini digunakan untuk menampilkan output dalam bentuk tertentu. Kegunaannya sama seperti fungsi `fmt.Println()`, hanya saja struktur outputnya didefinisikan di awal.  Fungsi `fmt.Printf()` tidak menghasilkan baris baru di akhir text, oleh karena itu digunakanlah literal newline yaitu `\n`, untuk memunculkan baris baru di akhir. Hal ini sangat berbeda jika dibandingkan dengan fungsi `fmt.Println()` yang secara otomatis menghasilkan new line (baris baru) di akhir.

Ketiga baris kode berikut ini akan menghasilkan output yang sama, meskipun cara penulisannya berbeda.

```go
 fmt.Printf("halo john wick!\n")
 fmt.Printf("halo %s %s!\n", firstName, lastName)
 fmt.Println("halo", firstName, lastName + "!")
```

### Deklarasi tanpa Tipe Data

Selain manifest typing, Go juga mengadopsi konsep type inference, yaitu metode deklarasi variabel yang tipe data-nya diketahui secara otomatis dari data/nilai variabel. Dengan metode jenis ini, keyword `var` dan tipe data tidak perlu ditulis, lalu operand `=` harus diganti dengan `:=`. 

 ```go
 var firstName string = "john"
 lastName := "wick"
 ```

Tipe data `lastName`  secara otomatis akan ditentukan menyesuaikan value-nya. Jika nilainya adalah berupa string  maka tipe data variabel adalah string. Pada contoh di atas, nilainya adalah string "wick".

Diperbolehkan untuk tetap menggunakan keyword `var` pada saat deklarasi meskipun tanpa menuliskan tipe data, dengan ketentuan tidak menggunakan
 tanda `:=`, melainkan tetap menggunakan `=`.

Tanda `:=` hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda `=`.

 ```go
 // menggunakan var, tanpa tipe data, menggunakan perantara "="
 var firstName = "john"
 // tanpa var, tanpa tipe data, menggunakan perantara ":="
 lastName := "wick"
 lastName = "ethan"
 lastName = "bourne"
 ```

### Deklarasi Multi Variabel

Go mendukung metode deklarasi banyak variabel secara bersamaan, caranya dengan menuliskan variabel-variabel-nya dengan pembatas tanda koma (,). Untuk pengisian nilainya-pun diperbolehkan secara bersamaan. Pengisian nilai juga bisa dilakukan bersamaan pada saat deklarasi.

Dengan menggunakan teknik type inference, deklarasi multi variabel bisa dilakukan untuk variabel-variabel yang tipe data satu sama lainnya berbeda.


```go
 var first, second, third string
 first, second, third = "satu", "dua", "tiga"

 var fourth, fifth, sixth string = "empat", "lima", "enam"

 seventh, eight, ninth := "tujuh", "delapan", "sembilan"

 one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
```

### Variabel Underscore _

Dalam Go, tidak boleh ada satupun variabel yang menganggur. Artinya, semua variabel yang dideklarasikan harus digunakan. Jika ada variabel yang tidak digunakan tapi dideklarasikan, error akan muncul pada saat kompilasi dan program tidak akan bisa di-run.

Underscore ( _ ) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.

Variabel underscore adalah predefined, jadi tidak perlu menggunakan `:=` untuk pengisian nilai, cukup dengan `=`  saja. Namun khusus untuk pengisian nilai multi variabel yang dilakukan dengan metode type inference, boleh di dalamnya terdapat variabel underscore.

Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan.

```go
 _ = "belajar Golang"
 _ = "Golang itu mudah"
 name, _ := "john", "wick"
 ```

Pada contoh di atas, variabel `name` akan berisikan text john , sedang nilai `wick` ditampung oleh variabel underscore, menandakan bahwa nilai tersebut tidak akan digunakan.

Perlu diketahui, bahwa isi variabel underscore tidak dapat ditampilkan. Data yang sudah masuk variabel tersebut akan hilang.

### Deklarasi dengan Keyword new

Fungsi `new()` digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya.

```go
 name := new(string)
 fmt.Println(name)   // 0x20818a220
 fmt.Println(*name)  // ""
```

Variabel `name` menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal). Untuk menampilkan nilai aslinya, variabel tersebut perlu didereference terlebih dahulu, caranya dengan menuliskan tanda asterisk (*) sebelum nama variabel.

### Deklarasi dengan Keyword make

Fungsi `make()` ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu channel, slice, dan  map.

## Tipe Data

Go mengenal beberapa jenis tipe data, di antaranya adalah tipe data numerik (desimal & non-desimal), string, dan boolean.

### Tipe Data Numerik Non-Desimal

Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui.
 
- uint , tipe data untuk bilangan cacah (bilangan positif).
- int , tipe data untuk bilangan bulat (bilangan negatif dan positif).

```go
 var positiveNumber uint8 = 89
 var negativeNumber = -1243423644
 fmt.Printf("bilangan positif: %d\n", positiveNumber)
 fmt.Printf("bilangan negatif: %d\n", negativeNumber)
```

String format `%d` pada  fmt.Printf()  digunakan untuk memformat data numerik non-desimal.

### Tipe Data Numerik Desimal

Tipe data numerik desimal yang perlu diketahui ada 2, `float32` dan `float64`. Perbedaan kedua tipe data tersebut berada di lebar cakupan nilai desimal yang bisa ditampung.

```go
 var decimalNumber = 2.62
 fmt.Printf("bilangan desimal: %f\n", decimalNumber)
 fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
```

String format `%f` digunakan untuk memformat data numerik desimal menjadi string. 

Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel `decimalNumber` adalah 2.620000. Jumlah digit yang muncul bisa dikontrol menggunakan `%.nf`, tinggal ganti `n` dengan angka yang diinginkan. Contoh: `%.3f` maka akan menghasilkan 3 digit desimal, `%.10f` maka akan menghasilkan 10 digit desimal.

### Tipe Data  bool  (Boolean)

 Tipe data  bool  berisikan hanya 2 variansi nilai, `true` dan `false`. Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan.

```go
 var exist bool = true
 fmt.Printf("exist? %t \n", exist)
```

Gunakan `%t` untuk memformat data `bool` menggunakan fungsi `fmt.Printf()`.

### Tipe Data string 

Ciri khas dari tipe data string adalah nilainya di apit oleh tanda quote atau petik dua ( " ). Contoh penerapannya:
 
```go
 var message string = "Halo"
 fmt.Printf("message: %s \n", message)
```

Selain menggunakan tanda quote, deklarasi string juga bisa dengan tanda grave accent/backticks ( ` ), tanda ini terletak di sebelah kiri tombol 1. Keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter di dalamnya tidak di escape, termasuk 
`\n`, tanda petik dua dan tanda petik satu, baris baru, dan lainnya. Semua akan terdeteksi sebagai string.

```go
 var message = `Nama saya "John Wick".
 Salam kenal.
 Mari belajar "Golang".`

 fmt.Println(message)
```

Ketika dijalankan, output akan muncul sama persis sesuai nilai variabel message di atas. Tanda petik dua akan muncul, baris baru juga muncul, sama persis.

### Nilai nil  & Zero Value

`nil`  bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya `nil` berarti memiliki nilai kosong.

Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya. 

- Zero value dari `string` adalah `""` (string kosong).
- Zero value dari `bool` adalah `false`.
- Zero value dari tipe numerik non-desimal adalah 0.
- Zero value dari tipe numerik desimal adalah 0.0.

Selain tipe data yang disebutkan di atas, ada juga tipe data lain yang zero value-nya adalah `nil`. `Nil`  merepresentasikan nilai kosong, benar-benar kosong. `nil`  tidak bisa digunakan pada tipe data yang sudah dibahas di atas. Beberapa tipe data yang bisa di-set nilainya dengan 
`nil`, di antaranya:

- pointer
- tipe data fungsi
- slice
- map 
- channel 
- interface kosong atau any (yang merupakan alias dari 
interface{} )

## Konstanta

Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah dideklarasikan. Inisialisasi nilai konstanta hanya dilakukan sekali saja di awal, setelah itu variabel tidak bisa diubah nilainya. Teknik type inference bisa diterapkan pada konstanta, caranya cukup dengan menghilangkan tipe data pada saat deklarasi.

```go
 const firstName string = "john"
 fmt.Print("halo ", firstName, "!\n")
 const lastName = "wick"
 fmt.Print("nice to meet you ", lastName, "!\n")
```

### Penggunaan Fungsi fmt.Print()

Fungsi ini memiliki peran yang sama seperti fungsi 
fungsi `fmt.Println()`, perbedaannya `fmt.Print()` tidak menghasilkan baris baru di akhir output-nya.

Perbedaan lainnya: nilai argument parameter yang ditulis saat pemanggilan fungsi akan di-print tanpa pemisah. Tidak seperti pada fungsi `fmt.Println()` yang nilai argument paremeternya dipisah menggunakan karakter spasi.

```go
 fmt.Println("john wick")
 fmt.Println("john", "wick")
 fmt.Print("john wick\n")
 fmt.Print("john ", "wick\n")
 fmt.Print("john", " ", "wick\n")
```

Bila menggunakan `fmt.Println()`, maka tidak perlu menambahkan spasi di tiap kata, karena fungsi tersebut akan secara otomatis menambahkannya di sela-sela text. Berbeda dengan `fmt.Print() yang perlu ditambahkan spasi, karena fungsi ini tidak menambahkan spasi secara otomatis di sela-sela nilai text yang digabungkan.

### Deklarasi Multi Konstanta

Sama seperti variabel, konstanta juga dapat dideklarasikan secara bersamaan. Berikut adalah contoh deklarasi konstanta dengan tipe data dan nilai yang berbeda.

```go
 const (
    square        = "kotak"
    isToday bool  = true
    numeric uint8 = 1
    floatNum      = 2.2
 )
```
- square , dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "kotak"
- isToday , dideklarasikan dengan metode manifest typing dengan tipe data bool dan nilainya true
- numeric , dideklarasikan dengan metode manifest typing dengan tipe data uint8 dan nilainya 1
- floatNum , dideklarasikan dengan metode type inference dengan tipe data float dan nilainya 2.2

Contoh deklarasi konstanta dengan tipe data dan nilai yang sama:
 
```go
 const (
    a = "konstanta"
    b
 )
```

- a  dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"
- b  dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"

 Berikut contoh gabungan dari keduanya

```go
 const (
    today string = "senin"
    sekarang
    isToday2 = true
 )
```

- today  dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
- sekarang  dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
- isToday2  dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true

Berikut contoh deklrasi multiple konstanta dalam satu baris:

```go
 const satu, dua = 1, 2
 const three, four string = "tiga", "empat"
```

- satu , dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 1
- dua , dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 2
- three , dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "tiga"
- four , dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "empat"


## Operator

### Operator Aritmatika

Operator aritmatika adalah operator yang digunakan untuk operasi yang sifatnya perhitungan.

- `+` penjumlahan
- `-` pengurangan
- `*` perkalian
- `/` pembagian
- `%` modulus / sisa hasil pembagian

Contoh penggunaan:

``` go
 var value = (((2 + 6) % 3) * 4 - 2) / 3
```

### Operator Perbandingan

- `==` apakah nilai kiri sama dengan nilai kanan
- `!=` apakah nilai kiri tidak sama dengan nilai kanan
- `<` apakah nilai kiri lebih kecil daripada nilai kanan
- `<=` apakah nilai kiri lebih kecil atau sama dengan nilai kanan
- `>` apakah nilai kiri lebih besar dari nilai kanan
- `>=` apakah nilai kiri lebih besar atau sama dengan nilai kanan

Contoh penggunaan:

```go
 var isEqual = (value == 2)
```

### Operator Logika

Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool  (bisa berupa variabel bertipe perbandingan).

- `&&`  kiri dan kanan
- `||` kiri atau kanan
- `!` negasi / nilai kebalikan

Contoh penggunaan:

```go
 var left = false
 var right = true
 
 var leftAndRight = left && right
 fmt.Printf("left && right \t(%t) \n", leftAndRight)

 var leftOrRight = left || right
 fmt.Printf("left || right \t(%t) \n", leftOrRight)

 var leftReverse = !left
 fmt.Printf("!left \t\t(%t) \n", leftReverse)
```

Hasil dari operator logika sama dengan hasil dari operator perbandingan, yaitu berupa boolean.

Berikut penjelasan statemen operator logika pada kode di atas.

- `leftAndRight` bernilai `false`, karena hasil dari `false` dan `true` adalah `false`.
- `leftOrRight` bernilai `true`, karena hasil dari `false` atau `true` adalah `true`.
- `leftReverse` bernilai `true`, karena negasi `false` adalah `true`.

Template `\t` digunakan untuk menambahkan indent tabulasi. Biasa dimanfaatkan untuk merapikan tampilan output pada console.

## Seleksi Kondisi

Yang dijadikan acuan oleh seleksi kondisi adalah nilai bertipe bool, bisa berasal dari variabel, ataupun hasil operasi perbandingan. Nilai tersebut menentukan blok kode mana yang akan dieksekusi.


### if, else if, else

Cara penerapan if-else di Go sama seperti pada bahasa pemrograman lain. Yang membedakan hanya tanda kurungnya (parentheses), di Go tidak perlu ditulis. Kode berikut merupakan contoh penerapan seleksi kondisi if else, dengan jumlah kondisi 4 buah.

```go
 var point = 8

 if point == 10 {
    fmt.Println("lulus dengan nilai sempurna")
 } else if point > 5 {
    fmt.Println("lulus")
 } else if point == 4 {
    fmt.Println("hampir lulus")
 } else {
    fmt.Printf("tidak lulus. nilai anda %d\n", point)
 }
```

#### Variabel Temporary pada if-else

Variabel temporary adalah variabel yang hanya bisa digunakan pada deretan blok seleksi kondisi di mana ia ditempatkan.

```go
point = 8840.0
 if percent := point / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
 } else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
 } else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
 }
 ```

### switch case

Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel, lalu kemudian di-cek nilainya. Contoh sederhananya seperti penentuan apakah nilai variabel x adalah 1,  2, 3, atau lainnya.

```go
point = 6
 switch point {
 case 8:
    fmt.Println("perfect")
 case 7:
    fmt.Println("awesome")
 default:
    fmt.Println("not bad")
 }
```

Untuk banyak kondisi, dapat dipisah dengan ( , ).

```go
var po = 6
switch po {
case 8:
	fmt.Println("perfect")
case 7, 6, 5, 4:
	fmt.Println("awesome")
default:
   fmt.Println("not bad")
}
```

Kondisi case 7, 6, 5, 4:  akan terpenuhi ketika nilai variabel atau 6 atau 5 atau 4.

Penggunaan switch case juga bisa menggunakan gaya if else.

```go
 var pointy = 6
 switch {
 case pointy == 8:
    fmt.Println("perfect")
 case (pointy < 8) && (pointy > 3):
    fmt.Println("awesome")
 default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
 }
 ```

Terdapat penggunaan fallthrough  dalam  switch. Ketika sebuah  case  terpenuhi, pengecekan kondisi tidak akan diteruskan ke case-case setelahnya. Keyword  fallthrough  digunakan untuk memaksa proses pengecekan tetap diteruskan ke  case  selanjutnya dengan tanpa menghiraukan nilai kondisinya, efeknya adalah case di pengecekan selanjutnya selalu dianggap  true (meskipun aslinya bisa saja kondisi tersebut tidak terpenuhi, akan tetap dianggap true).

```go
switch {
 case point == 8:
    fmt.Println("perfect")
 case (point < 8) && (point > 3):
    fmt.Println("awesome")
    fallthrough
 case point < 5:
    fmt.Println("you need to learn more")
 default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
 }
```

Di contoh, setelah pengecekan case (point < 8) && (point > 3)  selesai, dilanjut ke pengecekan case point < 5 , karena ada fallthrough  di situ. Dan kondisi case < 5  tersebut dianggap true  meskipun secara logika harusnya tidak terpenuhi.  Pada case  dalam sebuah switch  diperbolehkan terdapat lebih dari satu fallthrough .

### Seleksi Kondisi Bersarang

Seleksi kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi, yang mungkin juga berada dalam seleksi kondisi, dan seterusnya. Seleksi kondisi bersarang bisa dilakukan pada if-else, switch, ataupun kombinasi keduanya.

```go
point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
```

## Perulangan

Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti, selama kondisi yang dijadikan acuan terpenuhi. Biasanya disiapkan variabel untuk iterasi atau variabel penanda kapan perulangan akan diberhentikan. Di Go keyword perulangan hanya for saja, tetapi meski demikian, kemampuannya merupakan gabungan for, foreach, dan while ibarat bahasa pemrograman lain.

###  Perulangan Menggunakan Keyword for

```go
 for i := 0; i < 5; i++ {
    fmt.Println("Angka", i)
 }
```

Berikut jika argumen hanya kondisi.

```go
i = 0
for i < 5 {
    fmt.Println("Angka", i)
    i++
 }
```

Berikut jika  tanpa kondisi. Dengan ini akan dihasilkan perulangan tanpa henti (sama dengan `for true`). Pemberhentian perulangan dilakukan dengan menggunakan keyword `break`.

```go
 var i = 0
 for {
    fmt.Println("Angka", i)
    i++
 if i == 5 {
 break
    }
 }
 ```

 Berikut untuk perulangan for:range, biasa digunakan untuk me-loop-ing data gabungan (misal string, array, slice, map)

 ```go
 var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}
	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}
	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}
	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}
	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}
	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}
```

Keyword `break` digunakan untuk menghentikan secara paksa sebuah perulangan, sedangkan `continue` dipakai untuk memaksa maju ke perulangan berikutnya. Berikut merupakan contoh penerapan `continue` dan `break`.

```go
for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
	}
   ```

#### Perulangan Bersarang

```go
for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
 ```

 Di perulangan bersarang, break  dan continue  akan berlaku pada blok perulangan di mana ia digunakan saja. Ada cara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkan teknik pemberian label. Program untuk memunculkan matriks berikut merupakan contoh penerapan label perulangan.

```go
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
```
## Array

Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.

Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika int  maka tiap element zero value-nya adalah 0, jika bool  maka false, dan seterusnya. Setiap elemen array memiliki indeks berupa angka yang merepresentasikan posisi urutan elemen tersebut. Indeks array dimulai dari 0. Jika pengisian elemen array pada indeks yang tidak sesuai dengan jumlah alokasi menghasilkan error.

```go
 var names [4]string
 names[0] = "trafalgar"
 names[1] = "d"
 names[2] = "water"
 names[3] = "law"
 names[4] = "ez" // baris kode ini menghasilkan error
 fmt.Println(names[0], names[1], names[2], names[3])
 ```

Pengisian elemen array bisa dilakukan pada saat deklarasi variabel. Caranya dengan menuliskan data elemen dalam kurung kurawal setelah tipe data, dengan pembatas antar elemen adalah tanda koma ( , ). Khusus untuk deklarasi array dengan cara vertikal, tanda koma wajib dituliskan setelah setiap elemen (termasuk elemen terakhir), agar tidak memunculkan syntax error.

```go
 var fruits = [4]string{"apple", "grape", "banana", "melon"}
  fruits  = [4]string{
 "apple",
 "grape",
 "banana",
 "melon",
 }
 fmt.Println("Jumlah element \t\t", len(fruits))
 fmt.Println("Isi semua element \t", fruits)
 ```

 Deklarasi array yang nilainya diset di awal, boleh tidak dituliskan jumlah lebar array-nya, cukup ganti dengan tanda 3 titik ( ... ). Metode penulisan ini membuat kapasitas array otomatis dihitung dari jumlah elemen array yang ditulis.

 ```go
  var numbers = [...]int{2, 3, 2, 4, 3}
 fmt.Println("data array \t:", numbers)
 fmt.Println("jumlah elemen \t:", len(numbers))
 ```

 ### Array Multidimensi

Array multidimensi adalah array yang tiap elemennya juga berupa array.
 
Level kedalaman array multidimensi adalah tidak terbatas, bisa saja suatu array berisi elemen array yang setiap elemennya juga adalah nilai array, dst.

Cara deklarasi array multidimensi secara umum sama dengan array biasa, bedanya adalah pada array biasa, setiap elemen berisi satu nilai, sedangkan pada array multidimensi setiap elemen berisi array.
 
Khusus penulisan array yang merupakan subdimensi/elemen, boleh tidak dituliskan jumlah datanya. Contohnya bisa dilihat pada deklarasi variabel numbers2  di kode berikut.

```go
 var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
 var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
 fmt.Println("numbers1", numbers1)
 fmt.Println("numbers2", numbers2)
 ```

 ### Array dan Perulangan 

 Contoh menggunakan perulangan for untuk array:

 ```go
	var fruit = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruit); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
 ```

 Contoh menggunakan perulangan for-range untuk array:

 ```go
  var fruits = [4]string{"apple", "grape", "banana", "melon"}
 for i, fruit := range fruits {
    fmt.Printf("elemen %d : %s\n", i, fruit)
 }
 ```

 Contoh menggunakan perulangan for-range dan variabel underscore untuk array:

 ```go
 var fruits = [4]string{"apple", "grape", "banana", "melon"}
 for _, fruit := range fruits {
    fmt.Printf("nama buah : %s\n", fruit)
 }
 for i, _ := range fruits { }
 // atau
 for i := range fruits { }
 ```

 Contoh menggunakan make untuk array:

```go
 var fruits = make([]string, 2)
 fruits[0] = "apple"
 fruits[1] = "manggo"
 fmt.Println(fruits) 
 ```

## Slice

Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena slice merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

Contoh inisialisasi slice:

```go
 var fruits = []string{"apple", "grape", "banana", "melon"}
 fmt.Println(fruits[0]) // "apple"
 ```

Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabelnya, jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.

```go
 var fruitsA = []string{"apple", "grape"} // slice
 var fruitsB = [2]string{"banana", "melon"} // array
 var fruitsC = [...]string{"papaya", "grape"} // array
 ```

Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.

Slice bisa dibentuk dari array yang sudah didefinisikan, caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya. Contoh bisa dilihat pada kode berikut.

```go
 var fruits = []string{"apple", "grape", "banana", "melon"}
 var newFruits = fruits[0:2]

 fmt.Println(newFruits) // ["apple", "grape"]
 ```

Kode fruits[0:2]  maksudnya adalah pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2. Elemen yang memenuhi kriteria tersebut akan didapat, untuk kemudian disimpan pada  variabel lain sebagai slice baru. Pada contoh di atas, 
newFruits  adalah slice baru yang tercetak dari slice  fruits, dengan isi 2 elemen, yaitu "apple" dan "grape".

Ketika mengakses elemen array menggunakan satu buah indeks (seperti data[2]), nilai yang didapat merupakan hasil copy dari referensi aslinya. Berbeda dengan pengaksesan elemen menggunakan 2 indeks (seperti data[0:2]), nilai yang didapat adalah reference elemen atau slice.

Tabel berikut berisi contoh macam-macam operasi slice (atau slicing) menggunakan teknik 2 indeks yang bisa dilakukan di Go.

```go
 var fruits = []string{"apple", "grape", "banana", "melon"}
```

- fruits[0:2] berisi [apple, grape] (semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2).  
-  fruits[0:4] berisi [apple, grape, banana, melon] (semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4)
- fruits[0:0] berisi [] (menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0)
 fruits[4:4] berisi [] (menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4)
 fruits[4:0] berisi [] (error, pada penulisan fruits[a:b] nilai  a harus lebih kecil atau sama dengan  b) 
- fruits[:] berisi [apple, grape, banana, melon] (semua elemen). 
- fruits[2:] berisi [banana, melon] 
(semua elemen mulai indeks ke-2)
- fruits[:2] berisi [apple, grape] (semua elemen hingga sebelum indeks ke-2)

 Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.

 ```go
 var fruits = []string{"apple", "grape", "banana", "melon"}
 var aFruits = fruits[0:3]
 var bFruits = fruits[1:4]
 var aaFruits = aFruits[1:2]
 var baFruits = bFruits[0:1]
 fmt.Println(fruits)   // [apple grape banana melon]
 fmt.Println(aFruits)  // [apple grape banana]
 fmt.Println(bFruits)  // [grape banana melon]
 fmt.Println(aaFruits) // [grape]
 fmt.Println(baFruits) // [grape]
 // Buah "grape" diubah menjadi "pinnaple"
 baFruits[0] = "pinnaple"
 fmt.Println(fruits)   // [apple pinnaple banana melon]
 fmt.Println(aFruits)  // [apple pinnaple banana]
 fmt.Println(bFruits)  // [pinnaple banana melon]
 fmt.Println(aaFruits) // [pinnaple]
 fmt.Println(baFruits) // [pinnaple]
 ```

 ### Fungsi len()

 Fungsi `len()` digunakan untuk menghitung jumlah elemen slice yang ada. Sebagai contoh jika sebuah variabel adalah slice dengan data 4 buah, maka fungsi ini pada variabel tersebut akan mengembalikan angka 4.

 ```go
 var fruits = []string{"apple", "grape", "banana", "melon"}
 fmt.Println(len(fruits)) // 4
```

### Fungsi cap() 

Fungsi cap()  digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len , tapi bisa berubah seiring operasi slice yang dilakukan.

```go
fmt.Println(len(fruits))  // len: 4
 fmt.Println(cap(fruits))  // cap: 4
 var aFruits = fruits[0:3]
 fmt.Println(len(aFruits)) // len: 3
 fmt.Println(cap(aFruits)) // cap: 4
 var bFruits = fruits[1:4]
 fmt.Println(len(bFruits)) // len: 3
 fmt.Println(cap(bFruits)) // cap: 3
 ```

### Fungsi append()

Fungsi append() digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks paling akhir. Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya. Contoh penggunaannya bisa dilihat di kode berikut.

```go
var fruits = []string{"apple", "grape", "banana"}
 var cFruits = append(fruits, "papaya")
 fmt.Println(fruits)  // ["apple", "grape", "banana"]
 fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
 ```

 ### Fungsi copy()

 Fungsi  copy()  digunakan untuk men-copy elements slice pada  src  (parameter ke-2), ke  dst  (parameter pertama).
 
 ```go
  copy(dst, src)
  ```

 Jumlah element yang di-copy dari  src  adalah sejumlah lebar slice  dst  (atau len(dst) ). Jika jumlah slice pada  src  lebih kecil dari  dst , maka akan ter-copy semua. Lebih jelasnya silakan perhatikan contoh berikut.

```go
 dst := make([]string, 3)
 src := []string{"watermelon", "pinnaple", "apple", "orange"}
 n := copy(dst, src)

 fmt.Println(dst) // watermelon pinnaple apple
 fmt.Println(src) // watermelon pinnaple apple orange
 fmt.Println(n)   // 3
 ```

 ### Pengaksesan Elemen Slice dengan 3 Indeks

3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus menentukan kapasitasnya. Cara penggunaannya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1] . Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.

## Map

Map adalah tipe data asosiatif yang ada di Go yang berbentuk key-value pair. Data/value yang disimpan di map selalu disertai dengan key. Key sendiri harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang disimpan di map. 

Kalau dilihat, map  mirip seperti slice, hanya saja identifier yang digunakan untuk pengaksesan bukanlah index numerik, melainkan bisa dalam tipe data apapun sesuai dengan yang diinginkan.

```go
   var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0
```
## Fungsi
## Fungsi Multiple Return
## Fungsi Variadic
## Fungsi Closure
## Fungsi Sebagai Parameter
## Pointer
## Struct
## Method
## Properti Public dan Private
## Interface
## Any / Interface{} / Interface Kosong
## Reflect
## Goroutine
## Channel
## Buffered Channel
## Channel - Select
## Channel - Range & Close
## Channel - Timeout
## Defer & Exit
## Error, Panic, & Recover
## Layout Format String
## Random
## Time, Parsing Time, & Format Time
## Timer, Ticker, & Scheduler
## Time Duration
## Konversi Antar Tipe Data
## Fungsi String
## Regexp
## Encode - Decode Base64
## Hash Sha1
## Arguments & Flag
## Exec
## File
## Web Server
## URL Parsing
## JSON Data
## Web Service API Server
## Simple Client HTTP Request
## SQL
## NoSQL MongoDB
## Unit Test
## sync.WaitGroup
## sync.Mutex
## Go Vendoring
## Concurrency Pattern: Pipeline
## Concurrency Pattern: Simplified Fan-in Fan-out Pipeline
## Concurrency Pattern: Context Cancellation Pipeline
## Go Generics
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

# Instalasi Editor

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
## Tipe Data
## Konstanta
## Tipe Data
## Operator
## Seleksi Kondisi
## Perulangan
## Array
## Slice
## Map
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
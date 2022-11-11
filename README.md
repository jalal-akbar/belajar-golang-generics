# belajar-golang-generics
## Contents
- [Contents](#Contents)
- [Pengenalan Generics](#Pengenalan-Generics)
- [Membuat Projek](#Membuat-Project)
- [Type Paramater](#TypeParamater)
- [Multiple Type Paramater](#Multiple-Type-Paramater)
- [Comparable](#Comparable)
- [Type Parameter Inheritance](#Type-Parameter-Inheritance)
- [Type Sets](#Type-Sets)
- [Type Approximation](#Type-Approximation)
- [Type Inference](#Type-Inference)
- [Generic Type](#Generic-Type)
- [Generic Struct](#Generic-Struct)
- [Generic Interface](#Generic-Interface)
- [In Line Type Constaint](#In-Line-Type-Constaint)
- [Experimental Package](#Experimental-Package)
# Pengenalan-Generics
(1). Generic
- Generic adalah kemampuan menambahkan parameter type saat membuat function
- Berbeda dengan tipe data yang biasa kita gunakan di function, generic memungkinkan kita bisa mengubah-ubah bentuk tipe data sesuai dengan yang - kita mau.
- Fitur generics baru ada sejak Go-Lang versi 1.18
(2). Manfaat Generic
- Pengecekan ketika proses kompilasi
- Tidak perlu manual menggunakan pengecekan tipe data dan konversi tipe data
- Memudahkan programmer membuat kode program yang generic sehingga bisa digunakan oleh berbagai tipe data_

# Membuat-Project
```go
go mod init nama_project
go get github.com/stretchr/testify
```

# Type-Parameter
- [Type Parameter](#Type-Parameter)
(1). Tipe Paramater
- Untuk menandai sebuah function merupakan tipe generic, kita perlu menambahkan Type Parameter pada function tersebut
- Pembuatan Type Parameter menggunakan tanda [] (kurung kotak), dimana di dalam kurung kotak 
tersebut, kita tentukan nama Type Parameter nya
- Hampir sama dengan di bahasa pemrograman lain seperti Java, C# dan lain-lain, biasanya nama -
- Type Parameter hanya menggunakan satu huruf, misal T, K, V dan lain-lain. Walaupun bisa saja lebih dari satu huruf
(2). Tipe Constraint
- Di bahasa pemrograman seperti Java, C# dan lain-lain,
- Type Parameter biasanya tidak perlu kita tentukan tipe datanya, berbeda dengan di Go-Lang.
- Dari pengalaman yang dilakukan para pengembang Go-Lang, akhirnya di Go-Lang, Type Parameter -
wajib memiliki constraint
- Type Constraint merupakan aturan yang digunakan untuk menentukan tipe data yang diperbolehkan pada Type Parameter
- Contoh, jika kita ingin Type Parameter bisa digunakan untuk semua tipe data, kita bisa gunakan 
 interface{} (kosong) sebagai constraint nya
- Type Constraint yang lebih detail akan kita bahas di materi Type Sets
(3). Tipe Data Any
- Di Go-Lang 1.18, diperkenalkan alias baru bernama any untuk interface{} (kosong), ini bisa 
mempermudah kita ketika membuat Type Parameter dengan constraint interface{}, jadi kita cukup gunakan constraint any
(4). Menggunakan Tipe Parameter
- Setelah kita buat Type Parameter di function, selanjutnya kita bisa menggunakan Type Parameter 
tersebut sebagai tipe data di dalam function tersebut
- Misal nya digunakan untuk return type atau function parameter
- Kita cukup gunakan nama Type Parameter nya saja
- Type Parameter hanya bisa digunakan di functionnya saja, tidak bisa digunakan di luar function

# Multiple-Type-Parameter
- Penggunakan Type Parameter bisa lebih dari satu, jika kita ingin menambahkan multiple Type Parameter, 
kita cukup gunakan tanda , (koma) sebagai pemisah antar Type Parameter
- Nama Type Parameter harus berbeda, tidak boleh sama jika kita menambah Type Parameter lebih dari satu

# Comparable
- Selain any, di Go-Lang versi 1.18 juga terdapat tipe data bernama comparable
- Comparable merupakan interface yang diimplementasikan oleh tipe data yang bisa dibandingkan (menggunakan operator != dan ==), seperti booleans, numbers, strings, pointers, channels, interfaces, array yang isinya ada comparable type, atau structs yang fields nya adalah comparable type

# Type-Parameter-Inheritance
- Go-Lang sendiri sebenarnya tidak memiliki pewarisan, namun seperti kita ketahui, jika kita membuat sebuah type yang sesuai dengan kontrak interface, maka dianggap sebagai implementasi interface tersebut
- Type Parameter juga mendukung hal serupa, kita bisa gunakan constraint dengan menggunakan interface, maka secara otomatis semua interface yang compatible dengan type constraint tersebut bisa kita gunakan
- Diagram
```go
https://go.dev/blog/intro-generics
```

# Type-Sets
- Salah satu fitur yang menarik di Go-Lang Generic adalah Type Sets
- Dengan fitur ini, kita bisa menentukan lebih dari satu tipe constraint yang diperbolehkan pada type parameter
()Membuat Type Set
- Type Set adalah sebuah interface
- Cara membuat Type Set :
type NamaTypeSet interface {
     P | Q | R
}
- TypeSet hanya bisa digunakan pada type parameter, tidak bisa digunakan sebagai tipe data field atau variable
- Jika operator bisa digunakan di semua tipe data di dalam type set, maka operator tersebut bisa digunakan dalam kode generic

# Type-Approximation
- (1) Type Declaration
    - Kadang, kita sering membuat Type Declaration di Golang untuk tipe data lain, misal kita membuat tipe data Age untuk tipe data int
    - Secara default, jika kita gunakan Age sebagai type declaration untuk int, lalu kita membuat Type 
      Set yang berisi constraint int, maka tipe  data Age dianggap tidak compatible dengan Type Set yang 
      kita buat
- (2) Type Approximation
    - Untungnya, Go-Lang memiliki feature bernama Type Approximation, dimana kita bisa 
      menyebutkan bahwa semua constraint dengan tipe tersebut dan juga yang memiliki tipe dasarnya 
      adalah tipe tersebut, maka bisa digunakan
    - Untuk menggunakan Type Approximation, kita bisa gunakan tanda ~ (tilde)

# Type-Inference
- Type Inference merupakan fitur dimana kita tidak perlu menyebutkan Type Parameter ketika memanggil kode generic
- Tipe data Type Parameter bisa dibaca secara otomatis misal dari parameter yang kita kirim
- Namun perlu diingat, pada beberapa kasus, jika terjadi error karena Type Inference, kita bisa 
dengan mudah memperbaikinya dengan cara menyebutkan Type Parameter nya saja

# Generic-Type
- Sebelumnya kita sudah bahas tentang generic di function
- Generic juga bisa digunakan ketika membuat type

# Generic-Struct
- (1) Struct
    - Struct juga mendukung generic
    - Dengan menggunakan generic, kita bisa membuat Field dengan tipe data yang sesuai dengan Type Parameter
- (2) Method
    - Selain di function, kita juga bisa tambahkan generic di method (function di struct)
    - Namun, generic di method merupakan generic yang terdapat di struct nya. 
    - Kita wajib menyebutkan semua type parameter yang terdapat di struct, walaupun tidak kita
     gunakan misalnya, atau jika tidak ingin kita gunakan, kita bisa gunakan _ (garis bawah) sebagai 
     pengganti type parameter nya
    - Method tidak bisa memiliki type parameter yang mirip dengan di function

# Generic-Interface
- Generic juga bisa kita gunakan di Interface
- Secara otomatis, semua struct yang ingin mengikuti kontrak interface tersebut harus menggunakan generic juga

# In-Line-Type-Constraint
- (1) In Line Type Constraint
- Sebelum-sebelumnya, kita selalu menggunakan type declaration atau type set ketika membuat 
type constraint di type parameter
- Sebenarnya tidak ada kewajiban kita harus membuat type declaration atau type set jika kita ingin
 membuat type parameter, kita bisa gunakan -  secara langsung (in line) pada type constraint, 
 misalnya di awal kita sudah bahas tentang interface {} (kosong), tapi kita selalu gunakan type 
 declaration any
- Jika kita mau, kita juga bisa langsung gunakan interface { int | float32 | float64} dibanding membuat type set Number misalnya
- (2) Generic di Type Parameter 
- Pada kasus tertentu, kadang ada kebutuhan kita menggunakan type parameter yang ternyata type 
  tersebut juga generic atau memiliki type parameter
- Kita juga bisa menggunakan in line type constraint agar lebih mudah, dengan cara menambahkan
  type parameter selanjutnya, misal
- [S interface{[]E}, E interface{}], artinya S harus slice element E, dimana E boleh tipe apapun
- [S []E, E any], artinya S harus slice element E, dimana E boleh tipe apapun

# Experimental-Package
- (1) Experimental Package
- Saat versi Go-Lang 1.18, terdapat experimental package yang banyak menggunakan fitur Generic, namun belum resmi masuk ke Go-Lang Standard Library
- Kedepannya, karena ini masih experimental (percobaan), bisa jadi package ini akan berubah atau bahkan mungkin akan dihapus
```go
https://pkg.go.dev/golang.org/x/exp
```
- Silahkan install sebagai dependency di Go Modules menggunakan perintah 
go get golang.org/x/exp
- (2) Constraints Package
- Constraints Package berisi type declaration yang bisa kita gunakan untuk tipe data bawaan Go-Lang, misal Number, Complex, Ordered, dan lain-lain
```go
https://pkg.go.dev/golang.org/x/exp/constraints
```
- (3) Maps & Slice Package
- Terdapat juga package maps dan slices, yang berisi function untuk mengelola data Map dan Slice, namun sudah menggunakan fitur Generic
```go
https://pkg.go.dev/golang.org/x/exp/maps
```
```go
https://pkg.go.dev/golang.org/x/exp/slices
```
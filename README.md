# belajar-golang-generics

# Pengenalan-Generics
(1). Generic
- Generic adalah kemampuan menambahkan parameter type saat membuat function
- Berbeda dengan tipe data yang biasa kita gunakan di function, generic memungkinkan kita bisa mengubah-ubah bentuk tipe data sesuai dengan yang - kita mau.
- Fitur generics baru ada sejak Go-Lang versi 1.18
- (2). Manfaat Generic
- Pengecekan ketika proses kompilasi
- Tidak perlu manual menggunakan pengecekan tipe data dan konversi tipe data
- Memudahkan programmer membuat kode program yang generic sehingga bisa digunakan oleh berbagai tipe data_

# Membuat Project
```go
go mod init nama_project
go get github.com/stretchr/testify
```

# Type Parameter
- (1). Tipe Paramater
- Untuk menandai sebuah function merupakan tipe generic, kita perlu menambahkan Type Parameter pada function tersebut
- Pembuatan Type Parameter menggunakan tanda [] (kurung kotak), dimana di dalam kurung kotak 
tersebut, kita tentukan nama Type Parameter nya
- Hampir sama dengan di bahasa pemrograman lain seperti Java, C# dan lain-lain, biasanya nama -
- Type Parameter hanya menggunakan satu huruf, misal T, K, V dan lain-lain. Walaupun bisa saja lebih dari satu huruf
- (2). Tipe Constraint
- Di bahasa pemrograman seperti Java, C# dan lain-lain,
- Type Parameter biasanya tidak perlu kita tentukan tipe datanya, berbeda dengan di Go-Lang.
- Dari pengalaman yang dilakukan para pengembang Go-Lang, akhirnya di Go-Lang, Type Parameter -
wajib memiliki constraint
- Type Constraint merupakan aturan yang digunakan untuk menentukan tipe data yang diperbolehkan pada Type Parameter
- Contoh, jika kita ingin Type Parameter bisa digunakan untuk semua tipe data, kita bisa gunakan 
 interface{} (kosong) sebagai constraint nya
- Type Constraint yang lebih detail akan kita bahas di materi Type Sets
- (3). Tipe Data Any
- Di Go-Lang 1.18, diperkenalkan alias baru bernama any untuk interface{} (kosong), ini bisa 
mempermudah kita ketika membuat Type Parameter dengan constraint interface{}, jadi kita cukup gunakan constraint any
- (4). Menggunakan Tipe Parameter
- Setelah kita buat Type Parameter di function, selanjutnya kita bisa menggunakan Type Parameter 
tersebut sebagai tipe data di dalam function tersebut
- Misal nya digunakan untuk return type atau function parameter
- Kita cukup gunakan nama Type Parameter nya saja
- Type Parameter hanya bisa digunakan di functionnya saja, tidak bisa digunakan di luar function

# Multiple Type Parameter
# 30 Static Analysis

## Resume

`Static Code Analysis` biasanya dilakukan sebagai bagian dari Code Review dan dilakukan pada fase Implementasi dari Security Development Lifecycle (SDL). Static Code Analysisbiasanya mengacu pada menjalankan Static Code Analysis tools yang mencoba mencari kemungkinan kerentanan dalam source code dengan menggunakan teknik seperti Taint Analysis dan Data Flow Analysis.

`Static Code Analysis` adalah analisis perangkat lunak komputer yang dilakukan tanpa menjalankan program apa pun, berbeda dengan analisis dinamis, yang dilakukan pada program selama eksekusi. Istilah ini biasanya diterapkan pada analisis yang dilakukan oleh alat otomatis, dengan analisis manusia yang biasanya disebut "pemahaman program", pemahaman program, atau tinjauan kode. Terakhir, inspeksi perangkat lunak dan penelusuran perangkat lunak juga digunakan. Dalam kebanyakan kasus, analisis dilakukan pada beberapa versi kode sumber program, dan, dalam kasus lain, pada beberapa bentuk kode objeknya.

`Data flow analysis` digunakan untuk mengumpulkan informasi run-time (dinamis) tentang data dalam perangkat lunak saat dalam keadaan statis (Wögerer, 2005). Ada tiga istilah umum yang digunakan dalam data flow analysis, basic block (code), Control Flow Analysis (aliran data) dan Control Flow Path (jalur yang diambil data)

`Taint Analysis` adalah fitur dalam beberapa bahasa pemrograman komputer, seperti Perl, Ruby atau Ballerina yang dirancang untuk meningkatkan keamanan dengan mencegah pengguna jahat menjalankan perintah di komputer host. Taint checks mencari risiko keamanan khusus yang terutama terkait dengan situs web yang diserang menggunakan teknik seperti pendekatan SQL injection atau buffer overflow.

## Task

### Problem 1

Kode Awal

```go
func passed(score int) bool {
    switch {
    case score > 80:
        return true
    case score > 70:
        return true
    default:
        return true
    }
}
```

Hasil Perubahan

```go
func passed(score int) bool {
    switch {
    case score > 70:
        return true
    default:
        return false
    }
}
```

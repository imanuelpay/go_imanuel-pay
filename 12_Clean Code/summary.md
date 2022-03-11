# (12) Clean Code

## Resume

`Clean Code` adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah oleh programmer. Clean Code digunakan unutk kemudahan Work Collabiration, Feature Development, dan Fast Development.

Karakteristik Clean Code:

* Penamaan mudah dipahami.
* Mudah dieja dan dicari.
* Singkat namun mendeskripsikan konteks.
* Konsisten.
* Hindari penamaan konteks yang tidak perlu.
* Komentar yang baik,
* Fungsi yang mudah dipahami.
* Gunakan konvensi.
* Menggunakan formatting.

Prinsip Clean Code:

* `KISS (Keep It So Simple)`
  * Fungsi dan class harus kecil.
  * Fungsi dibuat untuk melakukan satu tugas saja.
  * Jangan gunakan terlalu banyak argumen dalam fungsi.
  * Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil, dan jumlahnya minimal.
* `DRY (Don't Repeat Yourself)`, Buatlah fungsi yang menghindari duplikasi, yang dapat digunakan secara berulang-ulang.
* `Refactoring`, proses refakturasi code yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal.

## Task

> NOTE: Tidak ada screenshot untuk section ini.

### Problem 1 - Analisis

Code yang harus di analisis:

```dart
class user {
    var id;
    var username;
    var password;
}

class userservice {
    user[] users = [];

    user[] getallusers() {
        return users;
    }

    user getuserbyid(userid) {
        return user.fillter(userid);
    }
}
```

Kekurangan dalam code tersebut adalah penulisan class dan function. Untuk penulisan agar lebih bisa membedakan class, function dan variable, biasanya:

* Class menggunakan gaya `Pascal Case`.
* Function menggunakan gaya `Camel Case`.
* Variable menggunakan gaya `Snake Case`.
* Untuk Variable di parameter sebuah Function
  biasanya menggunakan gaya `Camel Case`.
* Untuk Property dalam sebuah Class
  bisa menggunakan `Camel Case`.

> NOTE: Karena saya biasa menggunakan bahasa pemrograman PHP dan Dart.

Perbaikan penulisan :

```dart
/**
* Untuk class User dan UserService: Pascal Case
* Untuk function getAllUsers dan getUserById: Camel Case
* Untuk property id, username dan password: Camel Case
* Untuk variable users: Snake Case atau Camel Case
* Untuk variable di parameter sebuah function seperti userId: Camel Case
*/
class User {
    var id;
    var username;
    var password;
}

class UserService {
    User[] users = [];

    User[] getAllUsers() {
        return users;
    }

    User getUserById(userId) {
        return users.fillter(userId);
    }
}
```

Pada Golang:

```go
/**
* Memperjelas nama struct dan properties dalam struct, serta meperbaiki nama varibel
*/
package main

type User struct {
    id       int
    username string
    password string
}

type UserService struct {
    ListUser []User
}

func (userService UserService) getAllUsers() []User {
    return userService.ListUser
}

func (userService UserService) getUserById(id int) User {
    for _, user := range userService.ListUser {
        if id == user.id {
            return user
        }
    }

    return User{}
}
```

Berikut source code dari Problem 1 - Analisis:

[problem_1.go](praktikum/problem_1.go)

[problem_1.txt](praktikum/problem_1.txt)

### Problem 2 - Rewrite

Code yang harus di tulis ulang:

```dart
class kendaraan {
    var totalroda = 0;
    var kecepatanperjam = 0;
}

class mobil extends kendaraan {
    void berjalan() {
        tambahkecepatan(10);
    }

    tambahkecepatan(var kecepatanbaru) {
        kecepatanperjam += kecepatanbaru;
    }
}

void main() {
    mobilcepat = new mobil();
    mobilcepat.berjalan();
    mobilcepat.berjalan();
    mobilcepat.berjalan();

    mobillamban = new mobil();
    mobillamban.berjalan();
}
```

Code yang sudah di tulis ulang:

```dart
class Kendaraan {
    var totalRoda = 0;
    var kecepatanPerJam = 0;
}

class Mobil extends Kendaraan {
    void berjalan() {
        tambahKecepatan(10);
    }

    tambahKecepatan(var kecepatanBaru) {
        kecepatanPerJam += kecepatanBaru;
    }
}

void main() {
    mobilCepat = new Mobil();
    mobilCepat.berjalan();
    mobilCepat.berjalan();
    mobilCepat.berjalan();

    mobilLamban = new Mobil();
    mobilLamban.berjalan();
}
```

Pada Golang:

```go
/**
* Memperjelas nama struct dan properties dalam struct, serta meperbaiki nama varibel
*/
package main

type Kendaraan struct {
    totalRoda       int
    kecepatanPerJam float32
}

type Mobil struct {
    Kendaraan
}

func (mobil *Mobil) berjalan() {
    mobil.tambahKecepatan(10)
}

func (mobil *Mobil) tambahKecepatan(kecepatanBaru float32) {
    mobil.kecepatanPerJam += kecepatanBaru
}

func main() {
    mobilCepat := Mobil{}
    mobilCepat.berjalan()
    mobilCepat.berjalan()
    mobilCepat.berjalan()

    mobilLamban := Mobil{}
    mobilLamban.berjalan()
}
```

Berikut source code dari Problem 2 - Rewrite:

[problem_2.go](praktikum/problem_2.go)

[problem_2.txt](praktikum/problem_2.txt)

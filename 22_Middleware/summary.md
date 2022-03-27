# (22) Middleware

## Resume

`Middleware` merupakan sebuah sebuah layer tambahan sebelum `http request` dikerjakan oleh `action handler` atau `request handler` yang dituju. Middleware yang baik adalah yang hanya melakukan satu tugas, atau menganut pilosofi `do one` `thing well`. Sehingga untuk menangani banyak kasus di level middleware, mereka harus mampu bersifat `compose-able`.

Middleware juga merupakan sebuah komponen yang menjadi penenengah dalam koneksi data, baik dalam lingkup sendiri maupun integrasi antar sistem.Dengan adanya middleware maka sebuah akses URL dapat dibatasi. Sebagai contoh kita ingin membuat Log aktifitas oleh user, Nah kita dapat mencatat menggunakan middleware.

Middleware sendiri dibagi menjadi 2 yaitu `before` (awal) dan `after` (akhir), `Before` merupakan proses pemgoperasian pada fungsi middleware, jadi middleware yang kita buat bisa disebuut dengan before middleware. Sedangkan `after` middleware merupakan proses setelah melewati middleware, ketika middleware menginjinkan melanjutkan maka akan jadi middleware after.

Contoh kasus yang bisa ditangani oleh middleware antara lain:

- debugging information
- rate limiting
- pengecekan hak akses
- pengaturan CORS, dan masih banyak lagi

## Task

### Problem 1 - Logging & JWT Authentication

Source code: [Folder mvc-auth](praktikum/mvc-auth/)

Code Structure:

![problem_1TF.png](screenshots/problem_1TF.png)

Logging:

![problem_1LD.png](screenshots/problem_1LD.png)

- Base URL = `http://localhost:8080`
- Header :
  - Content-Type: `application/json`

#### Login

- Method : `POST`
- Endpoint : `/login`
- Request:
  | Field | Type |
  | :------ | :-------- |
  | email | string |
  | password | string |

![problem_1UL.png](screenshots/problem_1UL.png)

#### Get All Users

- Method : `GET`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/users`
![problem_1UGAF.png](screenshots/problem_1UGAF.png)
![problem_1UGAS.png](screenshots/problem_1UGAS.png)

#### Get User

- Method : `GET`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/users/:id`
![problem_1UGIF.png](screenshots/problem_1UGIF.png)
![problem_1UGIS.png](screenshots/problem_1UGIS.png)

#### Create User

- Method : `POST`
- Endpoint : `/users`
- Request:
  | Field | Type |
  | :------ | :-------- |
  | name | string |
  | email | string |
  | password | string |

![problem_1UC.png](screenshots/problem_1UC.png)

#### Update User

- Method : `PUT`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/users/:id`
- Request:
  | Field | Type |
  | :------ | :-------- |
  | name | string |
  | email | string |
  | password | string |

![problem_1UUIF.png](screenshots/problem_1UUIF.png)
![problem_1UUIS.png](screenshots/problem_1UUIS.png)
![problem_1UUISD.png](screenshots/problem_1UUISD.png)

#### Delete User

- Method : `DELETE`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/users/:id`
![problem_1UDIF.png](screenshots/problem_1UDIF.png)
![problem_1UDIS.png](screenshots/problem_1UDIS.png)
![problem_1UDISD.png](screenshots/problem_1UDISD.png)

#### Get All Books

- Method : `GET`
- Endpoint : `/books`
![problem_1BGA.png](screenshots/problem_1BGA.png)

#### Get Book

- Method : `GET`
- Endpoint : `/books/:id`
![problem_1BGI.png](screenshots/problem_1BGI.png)

#### Create Book

- Method : `POST`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/books`
- Request:
  | Field | Type |
  | :------ | :-------- |
  | name | string |
  | author | string |
  | publisher | string |

![problem_1BCF.png](screenshots/problem_1BCF.png)
![problem_1BCS.png](screenshots/problem_1BCS.png)

#### Update Book

- Method : `PUT`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/books/:id`
- Request:
  | Field | Type |
  | :------ | :-------- |
  | name | string |
  | author | string |
  | publisher | string |

![problem_1BUIF.png](screenshots/problem_1BUIF.png)
![problem_1BUIS.png](screenshots/problem_1BUIS.png)
![problem_1BUISD.png](screenshots/problem_1BUISD.png)

#### Delete Book

- Method : `DELETE`
- Header :
  - Authorization: `Bearer {token}`
- Endpoint : `/books/:id`
![problem_1BDIF.png](screenshots/problem_1BDIF.png)
![problem_1BDIS.png](screenshots/problem_1BDIS.png)
![problem_1BDISD.png](screenshots/problem_1BDISD.png)

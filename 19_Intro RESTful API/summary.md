# (19) Intro RESTful API

## Resume

`API` adalah singkatan dari `Application Programming Interface`. API sendiri merupakan interface yang dapat menghubungkan satu aplikasi dengan aplikasi lainnya. Dengan kata lain, peran API adalah sebagai perantara antar berbagai aplikasi berbeda, baik dalam satu platform yang sama atau pun lintas platform.

`REST`, singkatan bahasa Inggris dari `REpresentational State Transfer` atau transfer keadaan representasi, adalah suatu gaya arsitektur perangkat lunak untuk untuk pendistribusian sistem hipermedia seperti WWW. Sedangkan `RESTful` merupakan sebutan untuk web services yang menerapkan arsitektur REST.

`JSON`, singkatan dari JavaScript Object Notation, adalah suatu format ringkas pertukaran data komputer. Formatnya berbasis teks dan terbaca-manusia serta digunakan untuk merepresentasikan struktur data sederhana dan larik asosiatif.

Beberapa HTTP Method yang umum digunakan:

- GET
- POST
- PUT
- HEAD
- DELETE
- PATCH
- OPTIONS

Untuk HTTP Response Code dapat di lihat [disini](https://id.wikipedia.org/wiki/Daftar_kode_status_HTTP)

## Task

### Problem - Postman

File export collection postman [export_collection.json](praktikum/export_collection.json)

Mengatur Environment:
![env.png](screenshots/env.png)

#### A. Target API 1

- Base URL = `https://newsapi.org/v2`
- Header :
  - Content-Type: application/json
  - Authorization: {api.key}

\
**Everything**

Method : GET

- Endpoint : `/everything?q=minyak goreng`
  ![newsapi1.png](screenshots/newsapi1.png)
- Endpoint : `/everything?q=rusia&from=2022-03-10&to=2022-03-10&sortBy=popularity&language=id`
  ![newsapi2.png](screenshots/newsapi2.png)
- Endpoint : `/everything?domains=detik.com,kompas.com`
  ![newsapi3.png](screenshots/newsapi3.png)

\
**Top Headlines**

Method : GET

- Endpoint : `/top-headlines?q=ura`
  ![newsapi4.png](screenshots/newsapi4.png)
- Endpoint : `/top-headlines?country=id`
  ![newsapi5.png](screenshots/newsapi5.png)
- Endpoint : `/top-headlines?country=id&category=business`
  ![newsapi6.png](screenshots/newsapi6.png)
- Endpoint : `/top-headlines?sources=cnn`
  ![newsapi7.png](screenshots/newsapi7.png)

#### B. Target API 2

- Base URL = `https://swapi.dev/api`
- Header :
  - Content-Type: application/json
- Method : GET
- Endpoint : `/people`
  ![swapi1.png](screenshots/swapi1.png)
- Endpoint : `/films`
  ![swapi2.png](screenshots/swapi2.png)
- Endpoint : `/films/3`
  ![swapi3.png](screenshots/swapi3.png)
- Endpoint : `/planets`
  ![swapi4.png](screenshots/swapi4.png)
- Endpoint : `/species`
  ![swapi5.png](screenshots/swapi5.png)
- Endpoint : `/starships`
  ![swapi6.png](screenshots/swapi6.png)
- Endpoint : `/vehicles`
  ![swapi7.png](screenshots/swapi7.png)

#### C. Target API 3

- Base URL = `https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0`
- Header :
  - Content-Type: application/json
- Endpoint : `/user`
  - Method : GET
  ![sepulsa1.png](screenshots/sepulsa1.png)
- Endpoint : `/user/1`
  - Method : GET
  ![sepulsa2.png](screenshots/sepulsa2.png)
- Endpoint : `/user`
  - Method : POST
  ![sepulsa3.png](screenshots/sepulsa3.png)
- Endpoint : `/user/1`
  - Method : PUT
  ![sepulsa4.png](screenshots/sepulsa4.png)
- Endpoint : `/user/1`
  - Method : DELETE
  ![sepulsa5.png](screenshots/sepulsa5.png)

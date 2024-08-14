## Endpoints

### Ping Endpoint

Bu endpoint, servisin çalıştığını doğrulamak için kullanılır.

- **URL:** `http://3.121.230.112/ping`
- **Method:** `GET`
- **Response:**
  - HTTP Status Code: `200 OK`

### Verify Endpoint

Bu endpoint, verilen bilgilerin geçerliliğini doğrulamak için kullanılır.

- **URL:** `http://3.121.230.112/verify`
- **Method:** `POST`
- **Request Body:**
  ```json
  { 
      "TCKimlikNo": 12345678901, 
      "Ad": "Ali", 
      "Soyad": "Veli", 
      "DogumYili": 1990 
  } 
- **Request Body:**
  - HTTP Status Code: `200 OK`
  -  ```json
     {
      "isValid": false
     }

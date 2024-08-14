## Amazon Route 53 Public IP Redirection

I used Amazon Route 53 to redirect the EC2 public IP address to the domain `yusuftalhaklc.tech`.

## Endpoints

### Ping Endpoint

Bu endpoint, servisin çalıştığını doğrulamak için kullanılır.

- **URL:** `http://yusuftalhaklc.tech/ping`
- **Method:** `GET`
- **Response:**
  - HTTP Status Code: `200 OK`

### Verify Endpoint

Bu endpoint, verilen bilgilerin geçerliliğini doğrulamak için kullanılır.

- **URL:** `http://yusuftalhaklc.tech/verify`
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

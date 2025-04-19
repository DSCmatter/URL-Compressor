# URL Shortener

A fast, lightweight URL Shortener built using **Go**, **Gin**, and **Redis**. Easily shorten long URLs and get quick redirects using a simple REST API.

---

## Features

- Clean RESTful API
- Fast URL shortening powered by Redis
- Redirection from short URLs to long URLs
- Per-user scoped links (via user ID)
- Built with [Gin](https://github.com/gin-gonic/gin) and [Redis](https://redis.io)

---

## Tech Stack

| Layer    | Tool          |
|----------|---------------|
| Language | Go            |
| Web      | Gin Framework |
| Storage  | Redis         |

---

##  Prerequisites

- Go (Latest)
- Redis installed & running locally (default port `6379`)

---

##  Usage

### 1. Clone the repository

```
git clone https://github.com/your-username/go-url-shortener.git

cd go-url-shortener
```
### 2. Install dependencies

```
go mod tidy
```

### 3. Run the server

```
go run main.go

Server runs at: http://localhost:9808
```
### Using curl to request

```bash
curl --request POST \
--data '{
    "long_url": "https://www.amazon.in/Redragon-K617-Keyboard-Mechanical-Supported/dp/B09BVCVTBC?pd_rd_w=D4ne3&content-id=amzn1.sym.b5387062-d66f-4264-a2b3-7498b12700ed&pf_rd_p=b5387062-d66f-4264-a2b3-7498b12700ed&pf_rd_r=ZP7S4CFDM9J2YEGENZ5H&pd_rd_wg=T2cES&pd_rd_r=b4570cfa-44a0-4b56-8b76-600787cb8cb3&pd_rd_i=B09BVCVTBC&th=1",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}' \
  http://localhost:9808/create-short-url
```
The response will look like: 

```json

{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/FTNyH3Lh"
}
```

## 🤝 Contributing

Contributions are welcome and appreciated! ✨

If you have ideas for improvements, bug fixes, or want to add a new feature — feel free to open an issue or submit a pull request.
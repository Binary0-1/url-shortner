
# URL Shortener Service

A lightweight and secure URL shortener built with Go.  
Easily shorten URLs with unique, random shortcodes to prevent guessability, update existing URLs, and retrieve usage statistics.

---

## 🚀 Features

- Generate **random, unique shortcodes** to prevent URL scraping and brute-force enumeration.
- Update the destination URL for an existing shortcode.
- Retrieve URL statistics including **access count**, creation and update timestamps.
- Clean API design with proper HTTP status codes and JSON responses.
- Simple, idiomatic Go codebase using GORM and standard libraries.

---

## 🛠️ Tech Stack

- **Go (Golang)**
- **GORM** - ORM for database operations
- **PostgreSQL** (or any SQL database supported by GORM)
- Standard HTTP library for routing and JSON handling

---

## ⚙️ Installation & Setup

### Prerequisites

- Go 1.18+
- PostgreSQL (or compatible SQL database)
- Git

### Steps to run

```bash
git clone https://github.com/yourusername/urlshort.git
cd urlshort

# Install dependencies
go mod tidy

# Configure your DB connection in `db` package or via env vars

# (Optional) Run migrations if available
go run cmd/migrate.go

# Start the server
go run main.go
```

---

## 📡 API Endpoints

### 1. Shorten a URL

`POST /shorten`

**Request Body:**

```json
{
  "url": "https://example.com"
}
```

**Response:**

```json
{
  "id": 1,
  "shortCode": "a1B2c3",
  "url": "https://example.com",
  "createdAt": "2025-05-25T12:00:00Z",
  "updatedAt": "2025-05-25T12:00:00Z"
}
```

---

### 2. Update a Shortened URL

`PUT /shorten/{shortcode}`

**Request Body:**

```json
{
  "url": "https://new-url.com"
}
```

**Response:**

Returns the updated URL object.

---

### 3. Get URL Statistics

`GET /shorten/{shortcode}/stats`

**Response:**

```json
{
  "id": 1,
  "shortCode": "a1B2c3",
  "url": "https://example.com",
  "accessCount": 10,
  "createdAt": "2025-05-25T12:00:00Z",
  "updatedAt": "2025-05-25T12:00:00Z"
}
```

---

## 🗂 Project Structure

```
urlshort/
├── cmd/               # Entry points & migrations
├── db/                # Database connection & setup
├── models/            # GORM models
├── services/          # Business logic & HTTP handlers
├── utils/             # Utility functions (responses, errors)
├── main.go            # Application bootstrap
├── go.mod             # Dependencies
└── README.md          # This file
```

---

## 💡 Notes

- Shortcodes are 6-character alphanumeric strings, ensuring **unpredictability and professionalism**.
- The system tries up to 5 times to generate a unique shortcode before returning an error.
- URL statistics help monitor usage and prevent abuse.

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to open a Pull Request or issue.



---

Thanks for checking out the URL Shortener! 🚀

# 📎 Go URL Shortener

A simple URL shortening service built with:

- 🔥 [Gin](https://github.com/gin-gonic/gin) - Web Framework
- 🧠 [GORM](https://gorm.io) - ORM for database operations
- 💾 SQLite - Lightweight embedded database
- 🆔 [UUID](https://github.com/google/uuid) - For generating unique codes

---

## 🚀 Features

- Shorten long URLs
- Auto-generated unique short codes (UUID-based)
- Track creation timestamps
- Easy to extend with click count, custom codes, etc.

---

## 📦 Tech Stack

| Layer      | Tool           |
|------------|----------------|
| Language   | Go             |
| Web Server | Gin            |
| Database   | SQLite         |
| ORM        | GORM           |
| ID Gen     | UUID (Google)  |

---

## 🛠 Installation

1. **Clone the repo**

```bash
git clone https://github.com/longkhan786/url-shortener.git

cd url-shortener
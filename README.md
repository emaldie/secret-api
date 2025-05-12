# 🕵️ Shh...! - One-Time Secret Sharing App (Built in Go)

A simple, secure, and ephemeral text-sharing web app written in pure Go — no frameworks, no fluff.

Share secrets with confidence: the message is viewable **only once**, then it's gone forever. 🔐

---

## 🚀 Features

- 🧪 **One-time use** secret links (auto-deleted after read)
- ⏱️ Optional expiration: destroy after X minutes
- 📜 JSON RESTful API
- 🔒 Middleware for logging, rate-limiting, and error recovery
- 🎯 Pure `net/http`, no third-party frameworks

---

## 🌐 API Endpoints

### `GET /`

* Returns HTML form to submit a new secret

### `POST /create`

* Accepts form or JSON body: `{ "message": "my secret", "expire_minutes": 10 }`
* Returns: Secret link like `/s/2AjxZ3`

### `GET /s/{id}`

* Views the secret (once) and then deletes it

---

## 🧼 Planned Features

* [ ] QR code sharing
* [ ] View counter
* [ ] Client-side enhancements (auto-copy, countdown)

---

## 📖 Learning Goals

This project was built to learn and practice:

* ✅ Go’s `net/http` server
* ✅ Writing clean middleware (logging, recovery, rate limiting)
* ✅ Time-based expiration / onetime view and goroutines

---

## 📄 License

MIT – do whatever you want ✌️

---

## 🧠 Credits

Project idea and realization by emaldie XD

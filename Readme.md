# 🟢 Go Backend Learning Progress

**Tanggal:** 15 September 2025
**Fokus Hari Ini:** Routing & Handler di Go (Gin Framework)

---

## 🔹 Topik yang Dipelajari

* Membuat router menggunakan **Gin**
* Membuat **handler** untuk endpoint sederhana
* Struktur project untuk backend Go yang clean

---

## 🔹 Hal yang Dicoba

```go
func NewRouter(h *handler.UserHandler, log *logger.LoggerType) *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/users/:id", h.GetUser)
    r.POST("/users", h.CreateUser)

    return r
}
```

* Endpoint GET `/users/:id` untuk fetch user
* Endpoint POST `/users` untuk create user baru

---

## 🔹 Insight

* Struktur project yang rapi bikin maintenance lebih mudah
* Gin bikin routing lebih cepat dibanding `net/http` biasa
* Handler bisa di-inject dependency (contoh logger, service) → memudahkan testing

---

## 🔹 Next Step

* Pelajari **middleware** di Gin
* Menambahkan **validation & error handling** di handler
* Coba implement **database connection** (PostgreSQL/Redis)

---

Kalau mau, aku bisa bikinin **versi markdown yang lebih “visual”** kayak diary coding harian, lengkap sama **emoji progress tracker** biar gampang dilihat progres tiap hari di GitHub.

Mau aku bikinin versi itu juga?

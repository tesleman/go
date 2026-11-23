# System Design Урок 4: Мережеві протоколи та Мікросервіси

## 1. REST vs gRPC — аналог Express REST vs native gRPC
- **REST (HTTP/1.1):** Текстовий формат (JSON), людиночитаний, добре підходить для публічних API. Накладні витрати на парсинг та заголовки.
- **gRPC (HTTP/2):** Бінарний формат (Protocol Buffers), мультиплексування з'єднань через одне TCP-з'єднання, вбудована генерація клієнтського/серверного коду. Ідеально для міжсервісної комунікації.

> **Від Node.js:** 
> - REST ≈ Express.js (`app.get('/users/:id')` → JSON). Знаома модель.
> - gRPC ≈ `grpc-js` пакет у Node.js — обидва генерують код з `.proto` файлів.
> - **Головна різниця:** gRPC швидший (~5-10x) через бінарний protobuf замість JSON. Але складніший у налаштуванні.
>
> **Коли краще REST (Node.js/Go):** публічні API, проста інтеграція, браузерна сумісність. **Коли краще gRPC:** внутрішня комунікація між мікросервісами, високе навантаження, потрібна швидкість.

## 2. HTTP-сервер: Go net/http vs Node.js Express
```go
// Go — стандартна бібліотека, нічого встановлювати не треба
http.HandleFunc("/users/:id", handler)
http.ListenAndServe(":8080", nil)
```
```javascript
// Node.js — Express (потрібно npm install express)
app.get('/users/:id', (req, res) => {
    res.json({ id: req.params.id });
});
app.listen(8080);
```

> **Від Node.js:**
> - Go `net/http` ≈ Express.js (core routing). Go бібліотеки: `chi`, `gin` ≈ Express з middleware.
> - Go middleware: `func(next http.Handler) http.Handler` ≈ Express `app.use((req, res, next) => {})`.
> - Go JSON: `json.NewDecoder(r.Body).Decode(&x)` ≈ `express.json()` middleware.
> - **Головна різниця:** Go компілюється в нативний бінарник — не потрібен Node.js runtime. Швидший startup, менше RAM.
>
> **Коли краще Go:** high-load сервіси, мікросервіси, де важлива швидкість. **Коли краще Node.js:** швидка розробка, багато готових npm-пакетів, frontend-команда вже знає JS.

## 3. API Gateway та Service Discovery
- **API Gateway:** Єдина точка входу для клієнтів. Маршрутизація, аутентифікація, rate limiting, SSL termination.
- **Service Discovery (Consul, Etcd):** Динамічне знаходження IP-адрес мікросервісів у середовищі (Kubernetes).

> **Від Node.js:** У Node.js-проєктах часто використовують Express Gateway, Kong, або AWS API Gateway. Service Discovery — зазвичай через Consul або DNS. У обох стеках (Node.js/Go) архітектура однакова — різниця лише у мові реалізації конкретного мікросервісу.
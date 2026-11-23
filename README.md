# Go & System Design Learning Roadmap

Ця структура папок створена для вивчення та повторення мови програмування Go (Golang) та архітектури розподілених систем (System Design).

## 📂 Структура репозиторію

### 1. `01-go-review/` — Повторення та поглиблене вивчення Go
- **`01-basics/`** — Синтаксис, типи даних, змінні, функції, масиви/зрізи (slices), мапи (maps), вказівники.
- **`02-concurrency/`** — Горутини (Goroutines), канали (Channels), `sync` пакет (Mutex, WaitGroup, Once), паттерни конварентності (Worker Pools, Fan-in/Fan-out).
- **`03-interfaces-and-structs/`** — Структури, вбудовування (embedding), інтерфейси та їх підводні камені (type assertions, nil interfaces).
- **`04-advanced-features/`** — Дженерики (Generics), рефлексія (Reflection), робота з пам'яттю (Escape Analysis, Garbage Collector).
- **`05-testing-and-profiling/`** — Юніт-тести, Бенчмарки, профилювання (pprof), логування та обробка помилок (`errors.Is`/`As`).
- **`06-exercises/`** — Практичні задачі, алгоритми та кодинг-інтерв'ю на Go.

### 2. `02-system-design/` — Архітектура та System Design
- **`01-fundamentals/`** — Масштабованість, доступність, латентність vs пропускна здатність, CAP-теорема, балансування навантаження, кешування.
- **`02-databases-and-storage/`** — Реляційні (SQL) vs нереляційні (NoSQL), шардинг, реплікація, індексація (B-Tree, LSM-Tree), транзакції (ACID).
- **`03-messaging-and-queues/`** — Асинхронна архітектура, брокери повідомлень (Kafka, RabbitMQ), черги повідомлень, Pub/Sub.
- **`04-networking-and-microservices/`** — Протоколи (HTTP/1.1, HTTP/2, gRPC, WebSockets), API Gateway, Service Discovery, мікросервісні патерни.
- **`05-distributed-patterns/`** — Обмеження швидкості (Rate Limiting), Circuit Breaker, розподілені блокування, Consistent Hashing.
- **`06-case-studies/`** — Класичні системні дизайни:
  - Design URL Shortener (TinyURL)
  - Design Twitter / News Feed
  - Design Chat System (WhatsApp / Slack)
  - Design Uber / Ride Sharing
  - Design Video Streaming (YouTube / Netflix)

### 3. `03-practice-projects/` — Комплексні проєкти
- Мікросервіси на Go, реалізація кеш-серверів, черг повідомлень або веб-додатків з урахуванням вивчених паттернів System Design.

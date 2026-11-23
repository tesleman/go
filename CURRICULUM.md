# Go & System Design Curriculum & Roadmap

Цей документ містить детальну програму навчання та перелік найкращих матеріалів (курсів, книг, репозиторіїв) для заповнення створеної структури папок.

---

## 📌 Модуль 1: Поглиблене вивчення Go (`01-go-review/`)

### 1. `01-basics/`
- **Теми:** Типи даних, змінні, константи, управляючі конструкції (`if`, `switch`, `for`), функції, зрізи (Slices під капотом: len, cap, backing array), мапи (Maps), вказівники.
- **Рекомендовані джерела:** 
  - [Go by Example](https://gobyexample.com/)
  - [A Tour of Go](https://tour.golang.org/)

### 2. `02-concurrency/`
- **Теми:** Модель акторів/CSP, горутини (`go`), канали (буферизовані та небуферизовані), пакет `sync` (`Mutex`, `RWMutex`, `WaitGroup`, `Once`, `Cond`), паттерни конварентності (`Worker Pool`, `Pipeline`, `Fan-in/Fan-out`, `Context` пакет для скасування та таймаутів).
- **Рекомендовані джерела:**
  - Книга: *"Concurrency in Go"* by Katherine Cox-Buday.
  - [Go Concurrency Patterns (Google IO)](https://blog.golang.org/pipelines)

### 3. `03-interfaces-and-structs/`
- **Теми:** Структури та композиція (вбудовування), інтерфейси (статична та динамічна типізація), порожні інтерфейси (`any`/`interface{}`), `type assertions` та `type switches`, робота з `nil` інтерфейсами.
- **Рекомендовані джерела:**
  - Effective Go (розділ про інтерфейси)

### 4. `04-advanced-features/`
- **Teми:** Дженерики (Type parameters, constraints), рефлексія (`reflect` пакет), робота з пам'яттю (Escape Analysis, Garbage Collector у Go, вирівнювання пам'яті / struct padding).
- **Рекомендовані джерела:**
  - [The Go Blog: An Introduction to Generics](https://go.dev/blog/generics-tutorial)

### 5. `05-testing-and-profiling/`
- **Теми:** Модульне тестування (`testing`), табличні тести (Table-driven tests), моки (mocking), бенчмарки (`go test -bench`), профілювання (`pprof`, `trace`), розширена обробка помилок (`errors.Is`, `errors.As`, `fmt.Errorf("%w")`).
- **Рекомендовані джерела:**
  - [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

### 6. `06-exercises/`
- **Теми:** Алгоритмічні задачі на LeetCode / HackerRank розв'язані на Go, реалізація власних структур даних (однозв'язний список, черга, стек, дерево).

---

## 🏗 Модуль 2: System Design (`02-system-design/`)

### 1. `01-fundamentals/`
- **Теми:** Масштабованість (Vertical vs Horizontal), висока доступність (High Availability 99.999%), CAP-теорема та PACELC, балансування навантаження (Round Robin, Least Connections, L4 vs L7), кешування (LRU, LFU, Write-Through / Write-Back).
- **Рекомендовані джерела:**
  - [The System Design Primer](https://github.com/donnemartin/system-design-primer)

### 2. `02-databases-and-storage/`
- **Теми:** Реляційні бази даних (PostgreSQL/MySQL: шардинг, реплікація Master-Slave, партиціонування), NoSQL (Key-Value, Document, Column-family, Graph), індекси (B-Tree, LSM-Tree), транзакції та рівня ізоляції (ACID, BASE).
- **Рекомендовані джерела:**
  - Книга: *"Designing Data-Intensive Applications"* (DDIA) by Martin Kleppmann.

### 3. `03-messaging-and-queues/`
- **Теми:** Синхронна та асинхронна комунікація, черги повідомлень (RabbitMQ, SQS), шини подій / стріми (Apache Kafka, NSQ, Redis Streams), паттерн Publish-Subscribe, гарантії доставки (At-least-once, At-most-once, Exactly-once).

### 4. `04-networking-and-microservices/`
- **Теми:** Мережеві протоколи (TCP/IP, HTTP/1.1, HTTP/2, gRPC, Protobuf, WebSockets), API Gateway, Service Discovery (Consul, Etcd), мікросервісні патерни (BFF, CQRS, Event Sourcing).

### 5. `05-distributed-patterns/`
- **Теми:** Обмеження швидкості (Rate Limiting: Token Bucket, Leaky Bucket), Circuit Breaker, розподілені блокування (Redlock / etcd), Consistent Hashing, Service Mesh (Istio / Linkerd).

### 6. `06-case-studies/`
- **Теми:** Розбір кейсів з реального світу:
  - URL Shortener (TinyURL)
  - News Feed / Twitter Timeline
  - Chat System (WhatsApp / Messenger)
  - Ride Sharing (Uber)
  - Video Streaming (YouTube / Netflix)
- **Рекомендовані джерела:**
  - Книги: *"System Design Interview – An insider's guide"* (Vol. 1 & 2) by Alex Xu.
  - [ByteByteGo](https://bytebytego.com/)

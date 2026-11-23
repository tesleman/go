# Практичні проєкти (Go + System Design)

Проєкти для зв'язування теорії Go та архітектури System Design у коді. Рекомендований порядок — від простого до складного.

## Рівень 1: Початківець
1. **In-memory KV-сховище**
   - `map[string]string` + `sync.RWMutex`.
   - Протокол: простий TCP-сервер, що приймає команди `SET`, `GET`, `DEL`.
   - *Практика:* горутини, канали, mutex, робота з мережею (`net` пакет).

2. **LRU Cache**
   - Комбінація `container/list` (дredi список) та `map`.
   - Операції `Get`/`Put` за O(1).
   - *Практика:* структури даних, потокобезпека.

## Рівень 2: Середній
3. **Rate Limiter (Token Bucket)**
   - Реалізація middleware для HTTP-сервера.
   - *Практика:* System Design патерн у коді, `time.Ticker`.

4. **Key-Value Store з WAL (Write-Ahead Log)**
   - Додавання персистентності до KV-сховища шляхом запису операцій у лог-файл перед застосуванням.
   - *Практика:* реплікація-подібна механіка, робота з диском (`os`).

## Рівень 3: Поглибжений
5. **Простий URL Shortener**
   - HTTP API (`net/http` або фреймворк), Base62 генерація ключів, Redis кеш, SQL база (PostgreSQL).
   - *Практика:* повноцінний System Design кейс.

6. **Горутинний веб-скрапер / Worker Pool**
   - Конварентна обробка URL-ів з фіксованим пулом горутин, обмеженням частоти та Graceful Shutdown через `context`.

7. **Pub/Sub черга повідомлень (міні-Kafka)**
   - Брокер з топиками, підписниками та зберіганням логів на диску з партиціями.
   - *Практика:* кульмінація Go-конварентності та messaging-систем.

---

## Порівняння: якби ти робив те саме на Node.js

| Проєкт | Node.js стек | Go стек | Головна різниця |
|--------|-------------|---------|-----------------|
| KV-сховище | `net` (TCP) + `Map` | `net` (TCP) + `map` + `sync.RWMutex` | Node.js single-thread → без mutex; Go — потрібна синхронізація |
| LRU Cache | `Map` (JS Map зберігає порядок) + DoublyLinkedList | `container/list` + `map` | JS Map вже має insertion-order; Go — потрібно `container/list` |
| Rate Limiter | Express middleware + `setTimeout` | `http.Handler` middleware + `time.Ticker` | Node.js: однопотоковий таймер; Go: горутина + channel |
| KV + WAL | `fs.createWriteStream()` + Map | `os.File` + `bufio.Writer` + `sync` | Node.js: небажання блокувати event loop; Go: без проблем |
| URL Shortener | Express + Redis + pg | `net/http` / chi/gin + Redis + pgx | Go швидше (~2-5x), Node.js швидше розробляти |
| Worker Pool | `worker_threads` / `p-limit` | Goroutines + `chan` | Node.js: потрібні worker threads; Go: нативна підтримка |
| Pub/Sub (mini-Kafka) | EventEmitter + fs + Bacon | Goroutines + channels + `os.File` | Node.js EventEmitter ≈ Go channel, але Go краще для паралелізму |

### Ключове ментальне перемикання (Node.js → Go)
1. **Single-thread → Multi-thread:** Node.js все в одному потоці. У Go — треба думати про race conditions.
2. **Async/await → Goroutines + Channels:** Замість `await promise` — `go func() {}` + `ch <- data`.
3. **npm packages → Standard library first:** У Go багато чого є в стандарти (HTTP-сервер, тести, JSON, залежності через `go.mod`).
4. **try/catch → `if err != nil`:** Помилки перевіряються явно після кожного виклику.
5. **Express middleware → HTTP Handler wrappers:** `func(next http.Handler) http.Handler` замість `app.use()`.
6. **callback hell → goto hell:** Замість callback/Promise nesting — ланцюжок `if err != nil { return err }`.
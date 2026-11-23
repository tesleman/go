# Урок 2: Конкурентність у Go (Горутини та Канали)

Модель конварентності Go побудована на концепції CSP (Communicating Sequential Processes): *"Не спілкуйтеся через спільну пам'ять; натомість діліться пам'яттю, спілкуючись"*.

## 1. Горутини (Goroutines) — аналог async-функцій, але паралельні
Горутина — це легковажний потік виконання, керований рантаймом Go. Початковий стек ~2 КБ.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        fmt.Println(s)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go say("світ") // Запуск нової горутини (аналог async () => {})
    say("привіт")  // Виконання в основній горутині
}
```

> **Від JS:** `go say()` ≈ `say()` як async-функція, **але** вона виконується на іншому ядрі одночасно (parallel), а не в одному потоці (concurrent) як Event Loop Node.js.
> - **Node.js:** `await say()` — блокує тільки поточну корутину, інші Promise продовжують виконуватись. Один потік.
> - **Go:** `go say()` — реально виконується паралельно з основним кодом на різних CPU cores.
>
> **Коли краще Go:** коли потрібна реальна паралельна обробка (CPU-bound задачі, тисячі одночасних з'єднань). **Коли краще Node.js:** коли це IO-bound (API-запити, файли) і немає потреби у декількох потоках.

### Реальні приклади використання горутин

#### Приклад 1: Обробка веб-запитів (HTTP Server)
Найпоширеніше використання — кожен вхідний запит обробляється в окремій горутині. Коли клієнт підключається до сервера, Go автоматично створює горутину для з'єднання.

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Кожен запит виконується в окремій горутині автоматично!
    // Шукаємо в базі даних (симуляція 500ms затримки)
    time.Sleep(500 * time.Millisecond)
    fmt.Fprintf(w, "Відповідь на %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Сервер на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
    // 100 одночасних клієнтів = 100 горутин, кожна чекає окремо
}
```

> **Від Node.js:** У Node.js `http.createServer` обробляє запити через Event Loop. Кожен `await` не блокує інші запити, але це **кооперативна багатозадачність** — один довгий обчислювальний запит блокує всі інші. У Go кожен запит — **справжня паралельна горутина** на окремому ядрі.
>
> **Коли краще Go:** веб-сервери з високим навантаженням (тисячі одночасних з'єднань), де кожен запит може містити CPU-обчислення. **Коли краще Node.js:** прості API з IO-bound логікою (CRUD, проксі), де немає важких обчислень.

#### Приклад 2: Паралельна обробка файлів / зображень
```go
package main

import (
    "fmt"
    "sync"
)

func processImage(filename string, wg *sync.WaitGroup) {
    defer wg.Done()
    // Симуляція обробки зображення (2 секунди на файл)
    fmt.Printf("Обробка %s...\n", filename)
    // Тут може бути: resize, filter, compress — CPU-bound
    fmt.Printf("%s готово!\n", filename)
}

func main() {
    files := []string{"photo1.jpg", "photo2.jpg", "photo3.jpg", "photo4.jpg"}
    var wg sync.WaitGroup

    for _, f := range files {
        wg.Add(1)
        go processImage(f, &wg) // Запускаємо горутину для кожного файлу
    }

    wg.Wait() // Чекаємо завершення всіх горутин
    fmt.Println("Усі файли оброблено!")
}
```

> **Від Node.js:** У Node.js для обробки зображень використовують `worker_threads` або сторонні бібліотеки (sharp). Кожен worker — це важкий потік з власним V8. У Go горутина коштує ~2 КБ, тому можна запустити **мільйони** горутин без проблем.
>
> **Коли краще Go:** масова паралельна обробка даних (файли, зображення, аналітика). **Коли краще Node.js:** обробка 1-2 файлів з `sharp` — простіше API, не потрібна ручна синхронізація.

#### Приклад 3: Fan-in / Fan-out патерн (обробка даних)
Класичний патерн: багато горутин-виробників → канали → багато горутин-споживачів.

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()        // коли worker завершиться — лічильник WaitGroup зменшиться на 1
    for job := range jobs { // читаємо задачі з каналу jobs, поки він не закриється
        fmt.Printf("Worker %d обробляє job %d\n", id, job)
        results <- job * 2 // відправляємо результат у канал results
    }
}

func main() {
    jobs := make(chan int, 100)    // буферизований канал для задач (до 100 задач у черзі)
    results := make(chan int, 100) // буферизований канал для результатів
    var wg sync.WaitGroup          // лічильник активних горутин

    // Запускаємо 3 воркери
    for i := 1; i <= 3; i++ {
        wg.Add(1)                        // +1 до лічильника (очікуємо на 1 горутину)
        go worker(i, jobs, results, &wg) // запускаємо горутину-воркер
    }

    // Відправляємо 10 задач
    for i := 1; i <= 10; i++ {
        jobs <- i // відправляємо задачу в канал (воркери заберуть)
    }
    close(jobs) // закриваємо канал jobs — воркери вийдуть з `for range` після обробки всіх задач

    // Чекаємо завершення воркерів в окремій горутині
    go func() {
        wg.Wait()      // блокує, поки лічильник не стане 0 (усі 3 воркери завершились)
        close(results) // закриваємо канал results — main вийде з `for range`
    }()

    // Збираємо результати
    for r := range results { // читаємо результати, поки канал не закриється
        fmt.Println("Результат:", r)
    }
}
```

> **Від Node.js:** У Node.js це можна зробити через `async`/`await` з `Promise.all()`, але обмеження паралелізму доведеться робити вручну (бібліотеки типу `p-limit`). У Go канали + горутини дають це **з коробки**.
>
> **Коли краще Go:** pipelines обробки даних, черги задач, обробка подій. **Коли краще Node.js:** прості Promise.all для 5-10 паралельних HTTP-запитів.

#### Приклад 4: Таймери та періодичні задачі (Ticker)
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Створюємо тікер — він буде "цокати" кожну 1 секунду
    ticker := time.NewTicker(1 * time.Second)

    // Канал для сигналу зупинки (done — закриється коли треба зупинити горутину)
    done := make(chan bool)

    // Запускаємо горутину, яка слухає тікер і done
    go func() {
        count := 0
        for {
            select { // select — очікує на канали (як Promise.race)
            case <-done: // якщо прийшов сигнал з done — виходимо
                fmt.Println("Зупинка тікера")
                return
            case t := <-ticker.C: // якщо прийшов тік — обробляємо
                count++
                fmt.Printf("Тік %d: %s\n", count, t.Format("15:04:05"))
            }
        }
    }()

    // Основна горутина спить 5 секунд (чекаємо 5 тіків)
    time.Sleep(5 * time.Second)

    done <- true // Сигнал горутині: зупинись
    ticker.Stop() // Зупиняємо тікер (звільняє ресурси)
}
```

> **Від Node.js:** `setInterval` в Node.js — це callback в Event Loop. У Go — окрема горутина з `select`, яка не блокує основний потік. Можна мати **десятки тиків одночасно** без блокування.
>
> **Коли краще Go:** моніторинг метрик, health-checks, періодичні фонові задачі. **Коли краще Node.js:** прості `setInterval` для UI-оновлень або логування.

#### Приклад 5: Graceful Shutdown (плавне завершення)
```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    srv := &http.Server{Addr: ":8080"}

    go func() {
        fmt.Println("Сервер запущено")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Помилка сервера: %v\n", err)
        }
    }()

    // Чекаємо сигнал Ctrl+C
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit

    // Graceful shutdown з таймаутом 5 секунд
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    fmt.Println("Завершення сервера...")
    if err := srv.Shutdown(ctx); err != nil {
        fmt.Printf("Помилка завершення: %v\n", err)
    }
    fmt.Println("Сервер завершено")
}
```

> **Від Node.js:** У Node.js graceful shutdown роблять через `server.close()` + `process.on('SIGTERM')`. У Go це **ідіоматично** через `context.Context` + горутини. Контекст передається через усі шари — це вбудована система скасування.
>
> **Коли краще Go:** мікросервіси, де потрібен коректний shutdown з drain з'єднань. **Коли краще Node.js:** прості скрипти, де достатньо `process.exit()`.

#### Приклад 6: Rate Limiter (обмеження швидкості)
```go
package main

import (
    "fmt"
    "time"
)

func rateLimiter(requests chan string, delay time.Duration) {
    limiter := time.NewTicker(delay)
    defer limiter.Stop()
    for req := range requests {
        <-limiter.C // Блокуємо до наступного "тику"
        fmt.Println("Оброблено:", req)
    }
}

func main() {
    requests := make(chan string)

    // Rate limiter: 1 запит на 500ms
    go rateLimiter(requests, 500*time.Millisecond)

    // Відправляємо 5 запитів
    for i := 1; i <= 5; i++ {
        requests <- fmt.Sprintf("request-%d", i)
    }
    close(requests)
    time.Sleep(3 * time.Second) // Чекаємо обробку
}
```

> **Від Node.js:** У Node.js rate limiting — це бібліотеки (`express-rate-limit`, `bottleneck`). У Go — горутина з `time.Ticker`, **без залежностей**. Можна легко комбінувати з каналами для розподілу навантаження.
>
> **Коли краще Go:** API gateways, мікросервіси з обмеженням швидкості. **Коли краще Node.js:** Express-застосунки, де вже є готові middleware.

#### Приклад 7: Scout / Leader Election (координація)
```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    leaderCh := make(chan string, 1)

    // 3 кандидати "змагаються" за лідерство
    candidates := []string{"Node-A", "Node-B", "Node-C"}
    for _, name := range candidates {
        wg.Add(1)
        go func(candidate string) {
            defer wg.Done()
            // Симуляція "перегонів"
            fmt.Printf("%s намагається стати лідером...\n", candidate)
            leaderCh <- candidate // Перший, хто запише — лідер
        }(name)
    }

    wg.Wait()
    close(leaderCh)
    fmt.Println("Лідер:", <-leaderCh)
}
```

> **Від Node.js:** У Node.js таке роблять через Redis/ZooKeeper + бібліотеки. У Go — канали + горутини **з коробки**, без залежностей.
>
> **Коли краще Go:** розподілені системи, мікросервіси, координація. **Коли краще Node.js:** прості锁 contention без розподіленості.

## 2. Канали (Channels) — аналог Async Iterator / EventEmitter
Канали слугують для синхронізації та передачі даних між горутинами.

```go
ch := make(chan string) // Небуферизований канал

go func() {
    ch <- "дані з горутини" // Надсилання (аналог emit('event', data))
}()

msg := <-ch // Отримання (аналог await iterator.next())
fmt.Println(msg)
```

> **Від JS:** `ch <- data` ≈ `emitter.emit('event', data)` або `asyncIterator.next()`. `<-ch` ≈ `await asyncIterator.next()` або `emitter.once('event', handler)`.
> - Node.js EventEmitter — publish/subscribe (один → багато).
> - Go channel — point-to-point (один → один) з блокуванням до отримання.
>
> **Коли краще Go:** коли потрібна гарантія доставки і синхронізація між горутинами. **Коли краще Node.js:** коли потрібні broadcast-події (pub/sub) без блокування.

## 3. Синхронізація (sync) — аналог Atomics / Lock у Node.js
У Go потрібна явна синхронізація для спільної пам'яті. У Node.js — ні (бо один потік).

```go
var mu sync.Mutex
var count int

mu.Lock()
count++
mu.Unlock()
```

> **Від JS:** У Node.js немає mutex, бо все в одному потоці. Але якщо використовувати `worker_threads` — потрібні `SharedArrayBuffer` + `Atomics.add()`. 
> - Go `sync.Mutex` ≈ Node.js `Atomics` (але простіше).
> - Go `sync.WaitGroup` ≈ `await Promise.all([...])` — очікування завершення групи горутин.
>
> **Коли краще Go:** коли потрібні паралельні обчислення з спільним станом. **Коли краще Node.js:** коли обчислення легкі і достатньо одного потоку.

## 4. Практичні сценарії: коли використовувати горутини

| Сценарій | Чому горутини? | Аналог у Node.js |
|----------|----------------|-------------------|
| **Веб-сервер з 10k+ з'єднань** | Кожне з'єднання = горутина (~2KB), 10k = ~20MB RAM. Node.js = 10k Event Listeners в одному потоці | `http.createServer` + EventEmitter |
| **Паралельна обробка файлів** | 4 файли по 2с = 2с загалом (parallel), не 8с (sequential). Node.js `worker_threads` = важкі потоки | `worker_threads`, `sharp` |
| **API Gateway / Proxy** | Кожен запит → горутина → паралельні виклики до мікросервісів | `Promise.all()` з p-limit |
| **Фонові задачі (cron)** | Окрема горутина з `time.Ticker` — не блокує основний код | `setInterval` / `node-cron` |
| **Graceful shutdown** | `context.Context` + горутини = коректне завершення з drain з'єднань | `server.close()` + SIGTERM |
| **Rate Limiting** | Горутina + `time.Ticker` = обмеження швидкості без залежностей | `express-rate-limit`, `bottleneck` |
| **Розподілена координація** | Канали + горутини = лідер-елекція, розподіл задач | Redis/ZooKeeper + бібліотеки |
| **CPU-bound обчислення** | Go використовує всі ядра CPU автоматично (GOMAXPROCS) | Обмежено одним ядром без `worker_threads` |

## 5. Коли НЕ використовувати горутини

| Ситуація | Чому не треба | Краще |
|----------|---------------|-------|
| Прости лінійний скрипт | Немає паралельності — горутини марні | Звичайний код |
| IO-bound з 1-2 запитами | Event Loop Node.js достатньо | `async`/`await` |
| Робота з спільним масивом без синхронізації | Гонки даних (data races) — непередбачувані баги | `sync.Mutex` або канали |
| Короткі обчислення (<1ms) | Створення горутини коштує часу | Звичайний виклик |

## 6. Вправи для практики

1. **Web Spider:** Напиши горутину, яка паралельно "завантажує" 10 URL (симуляція `time.Sleep`). Порівняй час з послідовним виконанням.
2. **Ping Service:** Створи горутину, яка кожні 2 секунди "пінгує" сервер (симуляція), і горутину, яка відправляє результати в канал. Основна горутина читає з каналу.
3. **Worker Pool:** Створи 5 воркерів, які читають задачі з каналу. Відправ 20 задач. Воркери мають виконати кожну за 100ms.
4. **Merge Sort (паралельний):** Реалізуй merge sort, який розділяє масив на дві частини, сортує кожну в окремій горутині, потім об'єднує.
5. **Pipeline:** Створи pipeline з 3 стадій: генерація чисел → фільтрація парних → піднесення до квадрату. Кожна стадія — окрема горутина з каналом.
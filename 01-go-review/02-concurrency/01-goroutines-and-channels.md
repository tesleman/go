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
> **Коли краще Go:** коли потрібні паралельні обчислення з спільним станом. **Коли краще Node.js:** коли обчисленнялегкі і достатньо одного потоку.
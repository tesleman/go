package main

// ---------------------------------------------------------------------------
// Вправа 5. Mutex vs Канал — два способи лічильника.
//
// У цій вправі порівнюються два підходи до створення потокобезпечного лічильника:
//
//   (a) counterMutex — звичайний int, захищений sync.Mutex.
//   (b) counterChan  — лічильник, що працює в окремій горутині, яка приймає
//                       команди через канали.
//
// Запускай з -race:
//
//	go run -race .\01-go-review\02-concurrency\02-mutexes-exercises.go
//
// Задача: запусти 1000 горутин, кожна викликає Inc() на обох лічильниках.
// Перевір, що обидва показують 1000.
// Порівняй, який код тобі зрозуміліший.
//
// ---------------------------------------------------------------------------
// РЕальні приклади використання каналів у Go-проєктах:
//
// 1. Worker Pool — класичний патерн у веб-серверах, обробці завдань.
//    Приклад: HTTP-запити обробляються обмеженою кількістю горутин.
//
//    func worker(jobs <-chan Job, results chan<- Result) {
//        for j := range jobs {
//            results <- process(j)
//        }
//    }
//
//    jobs := make(chan Job)
//    results := make(chan Result)
//    for range 10 { go worker(jobs, results) }
//    jobs <- job
//    result := <-results
//
// 2. Graceful Shutdown — коректне завершення сервера.
//    Канал signal.Notify приймає SIGINT/SIGTERM, потім Shutdown.
//
//    quit := make(chan os.Signal, 1)
//    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//    <-quit
//    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//    server.Shutdown(ctx)
//
// 3. Rate Limiter — обмеження кількості запитів.
//    Токен-бакет реалізований через канал і тікер.
//
//    type RateLimiter struct {
//        tokens chan struct{}
//        ticker *time.Ticker
//    }
//    func (rl *RateLimiter) Allow() bool {
//        select {
//        case <-rl.tokens:
//            return true
//        default:
//            return false
//        }
//    }
//
// 4. Pipeline — обробка даних через ланцюг каналів.
//    Кожен етап — окрема горутини, що читає з попереднього і пише в наступний.
//
//    func sq(in <-chan int) <-chan int {
//        out := make(chan int)
//        go func() {
//            for n := range in {
//                out <- n * n
//            }
//            close(out)
//        }()
//        return out
//    }
//
// 5. Pub/Sub — шина подій, до якої підписуються кілька споживачів.
//    Використовується в мікросервісах для сповіщень.
//
//    type Bus struct {
//        subs map[string]chan Event
//        mu   sync.RWMutex
//    }
//    func (b *Bus) Publish(topic string, e Event) {
//        b.mu.RLock()
//        defer b.mu.RUnlock()
//        if ch, ok := b.subs[topic]; ok { ch <- e }
//    }
//
// 6. Actor — окрема горутини з приватним станом і каналом повідомлень.
//    Інші горутини не мають прямого доступу, лише надсилають команди.
//    CounterChan у цій вправі — саме такий приклад.
//
//    type Actor struct {
//        mailbox chan Message
//        done    chan struct{}
//    }
//    func (a *Actor) Send(msg Message) { a.mailbox <- msg }
//    func (a *Actor) Stop() { close(a.done) }
//
// ---------------------------------------------------------------------------
// Коли використовувати mutex, коли — канал:
//   Mutex — для простого захисту даних (лічильники, мапи, структури).
//   Канал — коли потрібна координація між горутинами, обмін даними,
//            архітектура "актор" або pipeline.
// ---------------------------------------------------------------------------

import (
	"fmt"
	"sync"
)

// counterMutex — лічильник на основі Mutex.
// Даних зберігається в полі n, доступ захищений mu.
type CounterMutex struct {
	mu sync.Mutex
	n  int
}

// Inc збільшує лічильник на 1.
func (c *CounterMutex) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

// Value повертає поточне значення лічильника.
func (c *CounterMutex) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// counterChan — лічильник на основі каналу.
// Всі дані зберігаються в одній горутині-власнику.
// Інші горутини не мають прямого доступу до лічильника,
// вони лише надсилають команди через канали.
type CounterChan struct {
	inc  chan struct{} // команда "інкремент"
	read chan chan int // команда "прочитати значення"
}

// NewCounterChan створює новий лічильник і запускає горутину-власника.
func NewCounterChan() *CounterChan {
	c := &CounterChan{
		inc:  make(chan struct{}),
		read: make(chan chan int),
	}
	go func() {
		n := 0
		for {
			select {
			case <-c.inc:
				n++
			case resp := <-c.read:
				resp <- n
			}
		}
	}()
	return c
}

// Inc надсилає команду інкременту.
func (c *CounterChan) Inc() {
	c.inc <- struct{}{}
}

// Value запитує поточне значення лічильника.
func (c *CounterChan) Value() int {
	resp := make(chan int)
	c.read <- resp
	return <-resp
}

// Exercise5 запускає 1000 горутин, які інкрементують обидва лічильники.
func Exercise5() {
	var wg sync.WaitGroup

	cm := &CounterMutex{}
	cc := NewCounterChan()

	for range 1000 {
		wg.Go(func() {
			cm.Inc()
			cc.Inc()
		})
	}
	wg.Wait()

	fmt.Printf("Вправа 5: mutex = %d, chan = %d (очікується 1000/1000)\n", cm.Value(), cc.Value())
}

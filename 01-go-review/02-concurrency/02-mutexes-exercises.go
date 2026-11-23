package main

import (
	"fmt"
	"sync"
)

// Вправи до уроку 02-mutexes.md
//
// Як перевіряти гонки:
//   go run -race 02-mutexes-exercises.go
//   (або: go run -race . )
//
// У main() розкоментуй по одній вправі за раз.

func main() {
	exercise1()
	// exercise2()
	// exercise3() // УВАГА: навмисний дедлок — запускай з Ctrl+C
	// exercise4()
	// exercise5()
	fmt.Println("Розкоментуй одну вправа у main()")
}

// ---------------------------------------------------------------------------
// Вправа 1. Виправ гонку.
//
// counter++ виконується з 1000 горутин одночасно.
// 1) Запусти з -race — побачиш WARNING: DATA RACE.
// 2) Додай sync.Mutex навколо counter++ так, щоб результат завжди був 1000.
// 3) Перезапусти з -race — гонки бути не повинно.
// ---------------------------------------------------------------------------
func exercise1() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for range 1000 {
		wg.Go(func() {
			// TODO: захистити інкремент м'ютексом
			mutex.Lock()
			defer mutex.Unlock()
			counter++
		})
	}
	wg.Wait()

	fmt.Printf("Вправа 1: counter = %d (очікується 1000)\n", counter)
}

// ---------------------------------------------------------------------------
// Вправа 2. Cache з RWMutex.
//
// Створи структуру Cache з полями:
//   mu   sync.RWMutex
//   data map[string]string
// Методи:
//   Get(k) (string, bool)  -> під mu.RLock()
//   Set(k, v)              -> під mu.Lock()
// Напиши тест: 100 горутин-читачів + 10 горутин-писарів паралельно.
// Перевір під -race.
// ---------------------------------------------------------------------------

// TODO: реалізуй тип Cache тут
// type Cache struct {
//     mu   sync.RWMutex
//     data map[string]string
// }

func exercise2() {
	// TODO: створи Cache, запусти читачів і писарів, перевір під -race
	fmt.Println("Вправа 2: див. TODO")
}

// ---------------------------------------------------------------------------
// Вправа 3. Дедлок навмисно (навчальна).
//
// Функція бере Lock() двічі поспіль -> самоблокування (дедлок).
// Запусти, переконайся що програма зависає (зупожи Ctrl+C).
// Потім ВИПРАВ: додай defer Unlock() після першого Lock(),
// або прибери другий Lock().
// ---------------------------------------------------------------------------
func exercise3() {
	var mu sync.Mutex

	mu.Lock()
	fmt.Println("Вправа 3: перший Lock узято")
	mu.Lock() // ДЕДЛОК — цей рядок ніколи не пройде
	fmt.Println("Цього рядка ти не побачиш, поки не виправиш")
	mu.Unlock()
	mu.Unlock()

	fmt.Println("Вправа 3: виправлено")
}

// ---------------------------------------------------------------------------
// Вправа 4. sync.Map.
//
// Перепиши Cache з вправи 2 на sync.Map (Store/Load/Delete/LoadOrStore).
// Порівняй читабельність коду (не лише швидкість).
// ---------------------------------------------------------------------------
func exercise4() {
	var m sync.Map

	// TODO: Store, Load, LoadOrStore, Delete
	m.Store("key", "value")

	if v, ok := m.Load("key"); ok {
		fmt.Printf("Вправа 4: key = %v\n", v)
	}
	fmt.Println("Вправа 4: див. TODO для повного варіанта")
}

// ---------------------------------------------------------------------------
// Вправа 5. Mutex vs Канал — два способи лічильника.
//
// (a) counterMutex: Mutex + int.
// (b) counterChan: горутина-власник, яка приймає інкременти по каналу.
// Порівняй, який код тобі зрозуміліший.
// ---------------------------------------------------------------------------

type counterMutex struct {
	mu sync.Mutex
	n  int
}

func (c *counterMutex) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *counterMutex) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// counterChan: володіння даними в однієї горутини.
type counterChan struct {
	inc  chan struct{}
	read chan chan int
}

func newCounterChan() *counterChan {
	c := &counterChan{
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

func (c *counterChan) Inc()       { c.inc <- struct{}{} }
func (c *counterChan) Value() int { resp := make(chan int); c.read <- resp; return <-resp }

func exercise5() {
	var wg sync.WaitGroup

	cm := &counterMutex{}
	cc := newCounterChan()

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			cm.Inc()
			cc.Inc()
		})
	}
	wg.Wait()

	fmt.Printf("Вправа 5: mutex = %d, chan = %d (очікується 1000/1000)\n", cm.Value(), cc.Value())
}

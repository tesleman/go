package main

import (
	"fmt"
	"sync"
)

// ---------------------------------------------------------------------------
// Вправа 1. Виправ гонку.
// ---------------------------------------------------------------------------

func exercise1() {
	var counter int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for range 1000 {
		wg.Go(func() {
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
// ---------------------------------------------------------------------------

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.data[k]
	return v, ok
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func exercise2() {
	cache := &Cache{
		data: make(map[string]string),
	}

	var wg sync.WaitGroup

	for i := range 10 {
		i := i
		wg.Go(func() {
			cache.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		})
	}

	for i := range 100 {
		i := i
		wg.Go(func() {
			key := fmt.Sprintf("key%d", i%10)
			if v, ok := cache.Get(key); ok {
				fmt.Printf("Got %s: %s\n", key, v)
			}
		})
	}

	wg.Wait()
	fmt.Println("Вправа 2: див. TODO")
}

// ---------------------------------------------------------------------------
// Вправа 3. Дедлок навмисно (навчальна).
// ---------------------------------------------------------------------------

func exercise3() {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Вправа 3: перший Lock узято")
	fmt.Println("Цього рядка ти не побачиш, поки не виправиш")
	fmt.Println("Вправа 3: виправлено")
}

// ---------------------------------------------------------------------------
// Вправа 4. sync.Map.
// ---------------------------------------------------------------------------

type Cache2 struct {
	data sync.Map
}

func (c *Cache2) Get(k string) (string, bool, error) {
	v, ok := c.data.Load(k)
	if !ok {
		return "", false, fmt.Errorf("key not found")
	}
	return v.(string), true, nil
}

func (c *Cache2) Store(k, v string) {
	c.data.Store(k, v)
}

func (c *Cache2) Delete(k string) {
	c.data.Delete(k)
}

func (c *Cache2) LoadOrStore(k, v string) (string, bool) {
	actual, loaded := c.data.LoadOrStore(k, v)
	return actual.(string), loaded
}

func exercise4() {
	var m Cache2

	m.Store("key", "value")

	if v, ok, _ := m.Get("key"); ok {
		fmt.Printf("Вправа 4: key = %v\n", v)
	}
	fmt.Println("Вправа 4: див. TODO для повного варіанта")
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	Exercise5()
}
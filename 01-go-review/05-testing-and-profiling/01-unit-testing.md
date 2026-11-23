# Урок 5: Тестування та Бенчмарки

## 1. Юніт-тести та Table-driven tests — аналог Jest / Vitest
Стандартний пакет `testing` — вбудований у Go. Table-driven tests — ідіоматичний підхід.

```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

> **Від Jest / Vitest:**
> - `func TestAdd(t *testing.T)` ≈ `describe('Add', () => { it('...', () => {}) })`.
> - Table-driven tests ≈ `test.each([...])('name', (a, b, want) => {})` — вбудовано, не треба Jest.
> - `t.Errorf()` ≈ `expect(x).toBe(y)` — тільки без fluent API. Перевіряєш вручну через `if`.
> - **Мокінг:** Go — створюєш фейковий тип з тими самими методами. Jest — `jest.fn()` / `vi.fn()` «з коробки».
>
> **Коли краще Go:** коли потрібні швидкі тести без залежностей, вбудовані у стандартну бібліотеку. **Коли краще Jest/Vitest:** коли потрібні мокі «з коробки», snapshot-тести, rich matchers.

## 2. Бенчмарки — аналог performance.now() у Node.js, але вбудовано
Для вимірювання продуктивності — функції з префіксом `Benchmark`.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(10, 20)
    }
}
```
Запуск: `go test -bench=.`

> **Від Node.js:**
> - Go: `func BenchmarkXxx(b *testing.B)` — вбудовано, автоматично прогріває та вимірює середнє.
> - Node.js: немає вбудованого бенчмарку в Jest/Vitest. Потрібно `console.time()` / `performance.now()` вручну, або `benchmark.js`.
> - Go: видає ns/op (наносекунди на операцію). Node.js: потрібно рахувати вручну.
>
> **Коли краще Go:** коли потрібна точна мікробенчмарк-статистика без сторонніх бібліотек. **Коли краще Node.js:** коли достатньо грубої оцінки `performance.now()`.
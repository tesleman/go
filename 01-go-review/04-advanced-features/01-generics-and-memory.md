# Урок 4: Дженерики та Робота з Пам'яттю

## 1. Дженерики (Generics) — аналог TS Generics
Починаючи з Go 1.18, мова підтримує дженерики через параметри типів.

```go
package main

import "fmt"

// Функція з обмеженням типу (Type Constraint)
func Sum[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Sum(5, 10))       // 15 (int)
    fmt.Println(Sum(3.14, 2.71))  // 5.85 (float64)
}
```

> **Від TS:**
> - `func Sum[T int | float64]()` ≈ `function sum<T extends number>(a: T, b: T): T`.
> - `int | float64` — union type constraint ≈ TS `extends number`.
> - **Головна різниця:** Go генерує конкретний машинний код для кожного типу (monomorphization). TS — стирає типи (type erasure). Тому Go дженерики швидші за TS.
>
> **Коли краще Go:** коли потрібна висока продуктивність з універсальними функціями. **Коли краще TS:** коли потрібна гнучкість та швидка розробка.

## 2. Робота з Пам'яттю — аналог GC у Node.js, але з контролем
- **Escape Analysis (Аналіз втеч):** Компілятор Go визначає, де виділити змінну — на стеку чи на купі (heap). Якщо змінна повертається за вказівником — вона "втікає" на купу.
- **Garbage Collector (GC):** Конкурентний триколірний mark-and-sweep GC з низькими паузами (~µs).

> **Від Node.js (V8):**
> - GC у Go і V8 — обидва generational + concurrent, обидва працюють у фоні.
> - **Головна різниця:** у Go можна **побачити** Escape Analysis (`go build -gcflags="-m"`). У Node.js — ні (V8 приховує).
> - Go: можна виділити на стеку (`var x int`) або на купі (`new(int)` / вказівник). Node.js: все на купі V8.
> - Go: GC паузи ~sub-microsecond. V8: теж низькі, але Go краще для low-latency.
>
> **Коли краще Go:** коли потрібен контроль над пам'яттю + мінімальні паузи GC. **Коли краще Node.js:** коли немає потреби у тонкому контролі пам'яті.
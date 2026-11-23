# Урок 3: Структури та Інтерфейси

## 1. Структури та Композиція — аналог класів, але без наслідування
Go не підтримує класи та класичне наслідування. Замість цього — композиція та вбудовування (embedding).

```go
package main

import "fmt"

type User struct {
    Name string
    Email string
}

func (u *User) Notify() {
    fmt.Printf("Відправка email на %s\n", u.Email)
}

// Композиція (Embedding) — аналог composition в TS
type Admin struct {
    User // Вбудована структура (аналог extends, але composition)
    Level int
}

func main() {
    ad := Admin{
        User: User{Name: "Іван", Email: "ivan@example.com"},
        Level: 1,
    }
    ad.Notify() // Метод User доступний напряму через Admin
}
```

> **Від JS/TS:**
> - `struct` ≈ `interface` (TS) / `class` (TS), **але** без методів всередині.
> - Методи додаються окремо: `func (u *User) Notify()` ≈ `class User { notify() {} }`.
> - Вбудовування (embedding) ≈ композиція в TS: `class Admin extends User`, **але** без поліморфізму. Це як `{ ...user, level: 1 }` — копіює поведінку, а не наслідує.
>
> **Коли краще Go:** коли потрібна композиція без глибокого успадкування (уникаємо "diamond problem"). **Коли краще TS:** коли потрібне класичне ООП з extends, поліморфізмом та модифікаторами доступу (`private`, `protected`).

## 2. Інтерфейси — аналог TS structural typing, але implicit
Інтерфейси у Go є **неявними (implicit)**. Тип реалізує інтерфейс, якщо має всі методи цього інтерфейсу.

```go
type Greeter interface {
    Greet() string
}

type EnglishSpeaker struct{}

// Метод автоматично задовольняє Greeter — не треба писати "implements"
func (e EnglishSpeaker) Greet() string {
    return "Hello!"
}

func PrintGreeting(g Greeter) {
    fmt.Println(g.Greet())
}
```

> **Від TS:**
> - TS: `class EnglishSpeaker implements Greeter { greet(): string { ... } }` — explicit.
> - Go: **не треба** `implements` — компілятор сам побачить, що `EnglishSpeaker` має `Greet() string` і автоматично задовольняє `Greeter`.
> - Це як TS **structural typing** без `implements` keyword: будь-який тип з правильними методами автоматично підходить.
>
> **Коли краще Go:** коли потрібна гнучка сумісність без явних декларацій (open-closed principle). **Коли краще TS:** коли потрібна явна документація контрактів через `implements`.
>
> **Підводний камінь:** Go-інтерфейс з `nil` всередині — це **не** `nil` (складається з двох слів: тип + дані). Аналог TS `null as any` — виглядає як null, але не є.
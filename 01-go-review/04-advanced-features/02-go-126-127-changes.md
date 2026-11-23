# Відмінності Go: 1.25.6 → 1.26 → 1.27

Ти вчив Go 1.25.6. Актуальна версія — **Go 1.27** (серпень 2026). Go 1.25 вже застаріла (підтримка припинена). Ось що змінилося:

---

## Go 1.26 (лютий 2026)

### 1. `new(expr)` — тепер приймає вираз, не лише тип
```go
//Було (Go 1.25):
x := 36
p := &x

// Стало (Go 1.26):
p := new(36) // *int, одразу з значенням
```
> **Аналогія з JS:** було `let x = 36; let p = x` (через зайву змінну). Стало як `const p = 36` — одразу. Найкорисніше для optional-полів у JSON (`Age *int`).

### 2. `errors.AsType[T]()` — типобезпечний аналог `errors.As`
```go
// Було (Go 1.25):
var pe *fs.PathError
if errors.As(err, &pe) { ... }

// Стало (Go 1.26):
if pe, ok := errors.AsType[*fs.PathError](err); ok {
    fmt.Println(pe.Path)
}
```
> **Аналогія з TS:** було як `if (err instanceof PathError)` з ручним приведенням. Стало — як дженерик `asType<T>()`.

### 3. Green Tea GC — новий збирач сміття (за замовчуванням)
- Покращена локальність пам'яті для багатоядерних CPU.
- Зменшення GC overhead на 10–40% для allocation-heavy programs.
- На AMD Zen 4+ / Intel Ice Lake+ — ще ~10% завдяки SIMD.
- Вимкнути: `GOEXPERIMENT=nogreenteagc` (приберуть у 1.27).

### 4. Рекурсивні дженерік-констрейнти
```go
// Go 1.26 дозволяє типам посилатися на себе в type parameters
type Adder[T interface{ Add(T) T }] interface {
    Add(T) T
}
```

### 5. Інше
- `io.ReadAll` — до 2x швидше, на 50% менше алокацій.
- `fmt.Errorf` без формату — стільки ж алокацій, скільки `errors.New`.
- `bytes.Buffer.Peek()` — підглянути в буфер без зсуву.
- `log/slog.NewMultiHandler` — логування в кілька таргетів.
- pprof Web UI —火焰图 (flame graph) за замовчуванням.

---

## Go 1.27 (серпень 2026) — найважливіші зміни

### 1. Generic Methods — дженерики в методах (!)
```go
// Було (Go 1.25): допоміжні функції на рівні пакету
func Map[T, U any](s []T, f func(T) U) []U { ... }

// Стало (Go 1.27): метод зі своїм type parameter
func (s *Stack[T]) Map[U any](f func(T) U) []U {
    out := make([]U, 0, len(s.items))
    for _, v := range s.items {
        out = append(out, f(v))
    }
    return out
}
```
> **Аналогія з TS:** було — тільки зовнішня функція `map<T, U>(arr, fn)`. Стало — як `arr.map<U>(fn)` — метод має власний дженерік-параметр. **Обмеження:** інтерфейси все ще не можуть мати generic-методів.

### 2. `encoding/json/v2` — новий стандартний пакет
```go
// Було (Go 1.25): encoding/json
import "encoding/json"

// Стало (Go 1.27): строгіші дефолти
import "encoding/json/v2"

// Зміни поведінки:
// 1. nil slice → [] замість null
// 2. Case-sensitive matching полів (було case-insensitive)
```
> **Аналогія з JS:** було як `JSON.parse()` / `JSON.stringify()` — терпимі до дрібниць. Стало — строгіший, як `zod` schema validation. Старий `encoding/json` працює, але під капотом вже v2.

### 3. `uuid` в стандартній бібліотеці
```go
import "uuid"

uuid.New()     // v4 (random)
uuid.NewV7()   // v7 (time-ordered, краще для БД-індексів)
```
> **Аналогія з Node.js:** не потрібен `npm install uuid`. Тепер вбудовано, як `crypto.randomUUID()` але з підтримкою v7.

### 4. Struct literal keys — promoted fields напряму
```go
type Habitat struct { Burrow string }
type Gopher struct {
    Name    string
    Habitat // Embedded
}

// Було (Go 1.25):
g := Gopher{Name: "Gopher", Habitat: Habitat{Burrow: "#42"}}

// Стало (Go 1.27):
g := Gopher{Name: "Gopher", Burrow: "#42"} // Напряму!
```

### 5. Function type inference — розширена
```go
// Go 1.27 виводить T у composite literals, conversions, channel sends:
formatters := []IntFormatter{GenericFormatter} // T виводиться автоматично
ch <- GenericFormatter
```

### 6. Goroutine Leak Profile (стабілізовано)
- Автоматичне виявлення заблокованих горутин через `runtime/pprof`.
- Раніше — експериментально в 1.26, тепер — повноцінний профіль.

### 7. SIMD (експериментально)
- Портативний SIMD API для векторних операцій.
- `simd` та `simd/archsimd` пакети.
- Потрібен `GOEXPERIMENT=simd`.

### 8. Post-Quantum Crypto
- `crypto/mldsa` — ML-DSA підпис (FIPS 204), стійкий до квантових комп'ютерів.
- Інтегровано в `crypto/x509` та `crypto/tls`.

### 9. Менш помітне, але важливе
- **HTTP/2** — винесено з `h2_bundle.go` у `net/http/internal/http2`.
- **HTTP/3** — каркас для QUIC в `net/http` (поки неекспортований).
- **Memory allocation** — small objects (<80B) швидші на 30%.
- **`go fix`** — тепер modernizers: автоматично оновлює код до нових API.
- **`go mod tidy`** — консолідує `require` блоки в `go.mod`.
- **`strings.CutLast`** — аналог `Cut` з кінця рядка.
- **pprof** — labels у tracebacks: `{request: 42}`.
- **macOS 12** — більше не підтримується (потрібен macOS 13+).

---

## Зведена таблиця

| Фіча | Було (1.25) | Стало (1.27) | Аналогія з JS/TS |
|------|------------|-------------|-------------------|
| `new()` | тільки тип | тип + значення | `let p = 36` vs `let x = 36; p = x` |
| Generic methods | нема | так | `arr.map<U>(fn)` |
| JSON | `encoding/json` | `encoding/json/v2` | `JSON.parse` → zod schema |
| UUID | сторонній `github.com/google/uuid` | `uuid` в stdlib | `npm install uuid` → вбудовано |
| GC | старий mark-sweep | Green Tea GC | V8 generational GC |
| errors.As | `var pe *T; errors.As(err, &pe)` | `errors.AsType[*T](err)` | `instanceof` з дженеріком |
| Promoted fields | `User{Habitat: Habitat{...}}` | `User{Burrow: "..."}` | spread `{...base, field}` |
| Goroutine leak | нема | профіль в pprof | --inspect --detect-leaks |

> **Головне для тебе:** дженерік-методи (`Map[U any]`) та `new(expr)` — це дві зміни, які ти відчуєш щодня. JSON v2 зламає код з case-insensitive полями — звертай увагу при міграції.
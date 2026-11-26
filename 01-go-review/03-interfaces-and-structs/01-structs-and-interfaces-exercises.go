package main

import "fmt"

// ---------------------------------------------------------------------------
// Урок 3: Структури та Інтерфейси — вправи для саморозбору.
//
// Інструкція: кожна функція містить TODO. Допиши код самостійно,
// запусти `go run -race .` і перевір, що вивід збігається з очікуваним.
// ---------------------------------------------------------------------------

// ===========================================================================
// Вправа 1. Структура та методи.
// ===========================================================================

type Product struct {
	Name  string
	Price float64
	Qty   int
}

// Додай метод TotalPrice на Product, що повертає Price * Qty.
func (p Product) TotalPrice() float64 {
	// TODO: реалізуй
	return 0
}

// Додаj метод Discount(percent float64), що зменшує Price на відсоток.
func (p *Product) Discount(percent float64) {
	// TODO: реалізуй
}

func exercise1() {
	p := Product{Name: "Ноутбук", Price: 1000, Qty: 2}
	fmt.Printf("Загальна вартість: %.2f (очікується 2000)\n", p.TotalPrice())

	p.Discount(10)
	fmt.Printf("Після знижки 10%%: Price = %.2f (очікується 900)\n", p.Price)
}

// ===========================================================================
// Вправа 2. Вбудовування (embedding) — Manager містить Employee.
// ===========================================================================

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Info() string {
	return fmt.Sprintf("%s, зарплата %.2f", e.Name, e.Salary)
}

// Створи структуру Manager з вбудованим Employee та полем Department string.
type Manager struct {
	Employee
	Department string
}

// Метод Manage() на Manager повертає рядок "<Name> керує <Department>".
func (m Manager) Manage() string {
	// TODO: реалізуй
	return ""
}

func exercise2() {
	m := Manager{
		Employee: Employee{Name: "Олена", Salary: 5000},
		Department: "ІТ",
	}
	fmt.Println(m.Info())        // метод Employee через embedding
	fmt.Println(m.Manage())      // TODO: має вивести "Олена керує ІТ"
}

// ===========================================================================
// Вправа 3. Інтерфейс задовольняється неявно.
// ===========================================================================

type Greeter interface {
	Greet() string
}

// Створи типи EnglishSpeaker та UkrainianSpeaker з методом Greet().
type EnglishSpeaker struct{}
type UkrainianSpeaker struct{}

func (EnglishSpeaker) Greet() string {
	// TODO: реалізуй
	return ""
}

func (UkrainianSpeaker) Greet() string {
	// TODO: реалізуй
	return ""
}

func PrintGreeting(g Greeter) {
	fmt.Println(g.Greet())
}

func exercise3() {
	// TODO: виклич PrintGreeting з обох типів
	// Очікується: Hello! / Привіт!
}

// ===========================================================================
// Вправа 4. Інтерфейс із nil всередині — це не nil.
// ===========================================================================

type Printer interface {
	Print()
}

type ConsolePrinter struct{}

func (ConsolePrinter) Print() { fmt.Println("print") }

func exercise4() {
	var p Printer = (*ConsolePrinter)(nil)
	fmt.Printf("p == nil: %v (очікується false)\n", p == nil)

	// TODO: безпечно виклич p.Print() лише якщо п safe (type assertion)
	// Підказка: cp, ok := p.(*ConsolePrinter)
}

// ===========================================================================
// Вправа 5. Інтерфейс Shape.
// ===========================================================================

type Shape interface {
	Area() float64
}

// Створи типи Rectangle (Width, Height float64) та Circle (Radius float64)
// з методами Area().
type Rectangle struct {
	Width, Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	// TODO: реалізуй
	return 0
}

func (c Circle) Area() float64 {
	// TODO: реалізуй
	return 0
}

func PrintArea(s Shape) {
	fmt.Printf("Площа: %.2f\n", s.Area())
}

func exercise5() {
	// TODO: виведи площу Rectangle{10, 5} і Circle{7}
	// Очікується: 50.00 / 153.86
}

// ===========================================================================
// Вправа 6. Type switch — визначення конкретного типу.
// ===========================================================================

type Describer interface {
	Describe() string
}

// Створи тип Book (Title string) з методом Describe().
type Book struct {
	Title string
}

func (b Book) Describe() string {
	// TODO: реалізуй
	return ""
}

func WhatIs(d Describer) {
	// TODO: type switch — якщо Book, виведи "книга: <Title>",
	// інакше виведи результат Describe()
}

func exercise6() {
	// TODO: виклич WhatIs з Book{Title: "Go в дії"}
	// Очікується: книга: Go в дії
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}

	// Перевірка існування ключа (comma-ok idiom)
	val, exists := m["banana"]
	if exists {
		fmt.Println("Значення:", val, "Ключ існує", exists)
	}
}

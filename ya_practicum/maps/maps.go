package main

import "fmt"

func main() {
	m := map[int]string{1: "first"}
	v, ok := m[1]
	fmt.Println(v, ok)
	delete(m, 2)
	fmt.Println("deleted [2] item: ", m)

	delete(m, 1)
	fmt.Println("deleted [1] item: ", m)

	v, ok = m[1]

	fmt.Println(v, ok)

	fmt.Printf("\n\nloops\n\n")

	m1 := map[string]string{"foo": "bar", "baz": "jazz", "john": "doe", "key": "value"}
	for k, v := range m1 {
		if k == "key" {
			m1[k] = "upd"
		}
		fmt.Printf("key: %v, value: %v \n", k, v)
	}
	fmt.Println(m1)

	fmt.Printf("\n\nFrequency characters\n\n")

	// предложение
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}
}

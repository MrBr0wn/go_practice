package slices_mut

import (
	"bytes"
	"fmt"
)

func Main1() {
	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
	// 1. Создаём слайс s с базовым массивом на 7 элементов. 
	// Четыре первых элемента будут доступны в слайсе.

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
	// 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
	// Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7), 
	// то базовый массив остаётся прежним.
	// Слайс s тоже изменился, но его длина осталась прежней

	slice2 := append(s[1:2], 7)
	fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
	// 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

	slice3 := append(s, slice1[1:]...)
	fmt.Println(len(slice3), cap(slice3)) // 8 14
	// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,  
	// что больше ёмкости базового массива.
	// Будет создан новый базовый массив ёмкостью 14,
	// ёмкость нового базового массива подбирается автоматически 
	// и зависит от текущего размера и количества добавленных элементов

	// 5. Легко проверить, что slice3 ссылается на новый базовый массив
	s[1] = 99
	fmt.Println(s, slice1, slice2, slice3)
	// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]

	result1, result2 := mutation()
	fmt.Printf("%s", result1)
	fmt.Printf("%s", result2)
}

func mutation() (string, string) {
	bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
	// return base64.StdEncoding.EncodeToString(bytes.TrimSpace(bSlice)),
	//  	base64.StdEncoding.EncodeToString(bSlice)

	return string(bytes.TrimSpace(bSlice)), string(bSlice)
}

package main

import "fmt"

func main() {
	fmt.Println("Case 1:")
	a := []int{1} // init slice a, len == 1, cap == 1 (default == len)
	b := a[:1]    // init slice b from a. a & b refer to the same base array, because:
	// ... Slicing does not copy the slice’s data. Slicing does not copy the slice’s data.
	// This makes slice operations as efficient as manipulating array indices.
	// Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.
	b[0] = 0 // also change base array from slice a
	fmt.Println(a)
	a[0] = 1 // change base array, set [0] = 1 for a and b
	fmt.Println(b)

	fmt.Println()
	fmt.Println("Case 2:")

	a1 := []int{1}
	b1 := a1[0:1]
	a1 = append(a1, 2) // Дело в механике append: если ёмкость слайса позволяет разместить 
	// добавляемые элементы (то есть разница между длиной слайса и его 
	// ёмкостью больше или равна количеству размещаемых элементов),
	// то append возвращает новый слайс, который получен из существующего 
	// слайса добавлением новых элементов внутри базового массива.
	// Если же ёмкость слайса не позволяет разместить эти элементы, 
	// то создаётся новый базовый массив подходящего размера, в него 
	// копируются все элементы переданного слайса и добавляются новые. 
	a1[0] = 10
	fmt.Println(a1, b1)

	fmt.Println()
	fmt.Println("Case 3:")

	a2 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a2), cap(a2))
	b2 := a2[2:]
	fmt.Println(len(b2), cap(b2))
	a2[0] = 100
	b2[0] = 200
	fmt.Println(a2, b2)
	// [1, 2, 3, 4, 5] - base array
	// a2 refer to the base array
	// b2 refer to the base array from 2 to end [3, 4, 5]
	// a2 change 0th elemt of base array [100, 2, 3, 4, 5]
	// b2 change himself 0th element and 2th element of the base array [100, 2, 200, 4, 5]
}

package main

import (
	"fmt"
	"slices_test/slices_mut"
)

func main() {
	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
	workDaysSlice[0] = 100
	fmt.Println(workDaysSlice)
	fmt.Println(weekTempSlice)

	slices_mut.Main1()
}

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// countCall — функция-обёртка для подсчёта вызовов
func countCall(f func(string)) func(string) {
	count := 0
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		count++
		fmt.Printf("Function %s called %d times\n", funcname, count)
		f(s)
	}
}

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start))
	}
}

func myprint(s string) {
	fmt.Println(s)
}

func main() {
	countedPrint := countCall(myprint)
	countedPrint("Hello world")
	countedPrint("Hi")

	countAndMetricPrint := metricTimeCall(countedPrint)
	countAndMetricPrint("Hello")
	countAndMetricPrint("World")
}

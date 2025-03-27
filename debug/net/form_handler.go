package main
// некоторые импорты нужны для проверки
import (
	"fmt"
	// "io"
	//"log"
	"net/http"
	//"net/url"
	//"os"
	//"time"
    "strconv" // вдруг понадобиться вам ;)
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	// проверяем что метод POST
	if r.Method == "POST" {
		
        err := r.ParseForm()
        if err != nil {
            fmt.Println("Parsing error:", err)
            return
        }

        numberString := r.PostFormValue("count")
        
        num, err := strconv.Atoi(numberString)
        if err != nil {
        	fmt.Println("это не число:", err)
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("это не число"))
            return
        }
        count += num
        w.Write([]byte(strconv.Itoa(count)))
		return
	}
    if r.Method == "GET" {
        w.Write([]byte(strconv.Itoa(count)))
        return
    }
	w.Write([]byte("Разрешен только метод GET & POST!"))
}

var count int

func main() {
	// Регистрируем обработчик для пути "/count"
	http.HandleFunc("/count", handler)
	// Запускаем веб-сервер на порту 3333
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
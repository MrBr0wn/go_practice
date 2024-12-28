package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Для сериализации используется функция json.Marshal() пакета json. Дана структура:
// type Person struct {
//		Name        string
//		Email       string
//		DateOfBirth time.Time
// }
// Напишите код, который будет сериализовывать структуру в json-строку следующего вида:
// {
//		"Имя": "Aлекс",
//		"Почта": "alex@yandex.ru"
// }

type Person struct {
	Name        string    `json:"Name"`
	Email       string    `json:"Email"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	fmt.Println("Exercise 2: ")
	man := Person{
		Name:        "Alex",
		Email:       "alex@example.com",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(jsMan))

	fmt.Println("\n\nExercise 3:")

	// Есть пример API-вызова в формате JSON (переменная input)
	// Опишите данный объект в виде структуры на Go, в учебных целях 
	// отбросив «так делать нельзя» и «это не дело».
	// И на вход подаются сырые данные, требуется их десериализовать

	type Response struct {
		Header struct {
			Code    int    `json:"code"`
			Message string `json:"message,omitempty"`
		} `json:"header"`
		Data []struct {
			Type       string `json:"type"`
			Id         int    `json:"id"`
			Attributes struct {
				Email      string `json:"email"`
				ArticleIds []int  `json:"article_ids"`
			}
		} `json:"data,omitempty"`
	}

	input := `
		{
				"header": {
						"code": 0,
						"message": "aaa"
				},
				"data": [{
						"type": "user",
						"id": 100,
						"attributes": {
								"email": "bob@yandex.ru",
								"article_ids": [10, 11, 12]
						}
				}]
		}`
	var res Response // eq res := Response{}
	err1 := json.Unmarshal([]byte(input), &res)

	if err1 != nil {
		fmt.Println("Unmarshaling error", err1)
	}
	fmt.Printf("Response: %+v\n", res)
}

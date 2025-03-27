package main

import (
    "fmt"
    "os"
    "bufio"
    "encoding/json"
    "io"
    // "bytes"
)

type Result struct {
    Average float64
}

type Student struct {
    LastName string
    FirstName string
    MiddleName string
    Birthday string
    Address string
    Phone string
    Rating []int
}

type Students struct {
	ID int
	Number string
	Year int
	Students []Student
}

func(s Students) AvgRating() float64 {
    totalStudents := float64(len(s.Students))
    var totalRatings float64

    for _, student := range s.Students {
        totalRatings += float64(len(student.Rating))
    }

    return totalRatings / totalStudents
}

func main() {
    w := bufio.NewWriter(os.Stdout)
    file, err := os.Open("students.json")
    defer file.Close()

    raw_json, err := io.ReadAll(file)

    // var result bytes.Buffer
    // reader := bufio.NewReader(file)
    // buf := make([]byte, 1024) // Буфер размером 1 КБ

    // for {
    //     n, err := reader.Read(buf)
    //     result.Write(buf[:n])
    //     if err != nil {
    //         if err == io.EOF {
    //             break
    //         }
    //     }
    // }

    // raw_json := result.Bytes()
    
    var s Students

    if err := json.Unmarshal(raw_json, &s); err != nil {
	    fmt.Println(err)
	    return
    }
    
    res := Result{Average: s.AvgRating()}
    
    output, err := json.MarshalIndent(res, "", "    ")
    if err != nil {
        fmt.Println(err)
        return
    }
    w.WriteString(string(output))
    w.Flush()
}
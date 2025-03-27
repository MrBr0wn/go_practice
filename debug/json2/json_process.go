package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io"
)

type GlobalId struct {
    Id int `json:"global_id"`
}

func main() {
    file, _ := os.Open("file.json")
    defer file.Close()

    raw_json, _ := io.ReadAll(file)
    
    var s []GlobalId

    if err := json.Unmarshal(raw_json, &s); err != nil {
	    fmt.Println(err)
	    return
    }
    var total int
    for _, id := range s {
        total += id.Id
    }
    
    fmt.Println(total)
}